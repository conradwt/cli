package errors

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"regexp"
	"strings"

	corev1 "github.com/confluentinc/cc-structs/kafka/core/v1"
	"github.com/confluentinc/ccloud-sdk-go-v1"
	mds "github.com/confluentinc/mds-sdk-go/mdsv1"
	"github.com/confluentinc/mds-sdk-go/mdsv2alpha1"
	srsdk "github.com/confluentinc/schema-registry-sdk-go"
	"github.com/pkg/errors"
)

/*
	HANDLECOMMON HELPERS
	see: https://github.com/confluentinc/cli/blob/master/errors.md
*/

const quotaExceededRegex = ".* is currently limited to .*"

type errorResponseBody struct {
	Error   []errorDetail `json:"errors"`
	Message string        `json:"message"`
}

type errorDetail struct {
	Detail string `json:"detail"`
}

func catchTypedErrors(err error) error {
	if typedErr, ok := err.(CLITypedError); ok {
		return typedErr.UserFacingError()
	}
	return err
}

func parseMDSOpenAPIErrorType1(err error) (*MDSV2Alpha1ErrorType1, error) {
	if openAPIError, ok := err.(mdsv2alpha1.GenericOpenAPIError); ok {
		var decodedError MDSV2Alpha1ErrorType1
		err = json.Unmarshal(openAPIError.Body(), &decodedError)
		if err != nil {
			return nil, err
		}
		if reflect.DeepEqual(decodedError, MDSV2Alpha1ErrorType1{}) {
			return nil, fmt.Errorf("failed to parse")
		}
		return &decodedError, nil
	}
	return nil, fmt.Errorf("unexpected type")
}

func parseMDSOpenAPIErrorType2(err error) (*MDSV2Alpha1ErrorType2Array, error) {
	if openAPIError, ok := err.(mdsv2alpha1.GenericOpenAPIError); ok {
		var decodedError MDSV2Alpha1ErrorType2Array
		err = json.Unmarshal(openAPIError.Body(), &decodedError)
		if err != nil {
			return nil, err
		}
		if reflect.DeepEqual(decodedError, MDSV2Alpha1ErrorType2Array{}) {
			return nil, fmt.Errorf("failed to parse")
		}
		return &decodedError, nil
	}
	return nil, fmt.Errorf("unexpected type")
}

func catchMDSErrors(err error) error {
	switch err2 := err.(type) {
	case mds.GenericOpenAPIError:
		return Errorf(GenericOpenAPIErrorMsg, err.Error(), string(err2.Body()))
	case mdsv2alpha1.GenericOpenAPIError:
		if strings.Contains(err.Error(), "Forbidden Access") {
			return NewErrorWithSuggestions(UnauthorizedErrorMsg, UnauthorizedSuggestions)
		}
		openAPIError, parseErr := parseMDSOpenAPIErrorType1(err)
		if parseErr == nil {
			return openAPIError.UserFacingError()
		} else {
			openAPIErrorType2, parseErr2 := parseMDSOpenAPIErrorType2(err)
			if parseErr2 == nil {
				return openAPIErrorType2.UserFacingError()
			} else {
				return Errorf(GenericOpenAPIErrorMsg, err.Error(), string(err2.Body()))
			}
		}
	}
	return err
}

// All errors from CCloud backend services will be of corev1.Error type
// This catcher function should then be used last to not accidentally convert errors that
// are supposed to be caught by more specific catchers.
func catchCoreV1Errors(err error) error {
	if err, ok := err.(*corev1.Error); ok {
		return Wrap(err, CCloudBackendErrorPrefix)
	}
	return err
}

func catchCCloudTokenErrors(err error) error {
	switch err.(type) {
	case *ccloud.InvalidLoginError:
		return NewErrorWithSuggestions(InvalidLoginErrorMsg, AvoidTimeoutSuggestions)
	case *ccloud.InvalidTokenError:
		return NewErrorWithSuggestions(CorruptedTokenErrorMsg, CorruptedTokenSuggestions)
	case *ccloud.ExpiredTokenError:
		return NewErrorWithSuggestions(ExpiredTokenErrorMsg, ExpiredTokenSuggestions)
	}
	return err
}

func catchOpenAPIError(err error) error {
	if openAPIError, ok := err.(srsdk.GenericOpenAPIError); ok {
		return New(string(openAPIError.Body()))
	}
	return err
}

/*
Error: 1 error occurred:
	* error creating ACLs: reply error: invalid character 'C' looking for beginning of value
Error: 1 error occurred:
	* error updating topic ENTERPRISE.LOANALT2-ALTERNATE-LOAN-MASTER-2.DLQ: reply error: invalid character '<' looking for beginning of value
*/
func catchCCloudBackendUnmarshallingError(err error) error {
	backendUnmarshllingErrorRegex := regexp.MustCompile(`reply error: invalid character '.' looking for beginning of value`)
	if backendUnmarshllingErrorRegex.MatchString(err.Error()) {
		errorMsg := fmt.Sprintf(prefixFormat, UnexpectedBackendOutputPrefix, BackendUnmarshallingErrorMsg)
		return NewErrorWithSuggestions(errorMsg, UnexpectedBackendOutputSuggestions)
	}
	return err
}

/*
	CCLOUD-SDK-GO CLIENT ERROR CATCHING
*/

func CatchCCloudV2Error(err error, r *http.Response) error {
	if err == nil {
		return nil
	}

	if r == nil {
		return err
	}

	body, _ := io.ReadAll(r.Body)
	var resBody errorResponseBody
	_ = json.Unmarshal(body, &resBody)
	if len(resBody.Error) > 0 {
		detail := resBody.Error[0].Detail
		if ok, _ := regexp.MatchString(quotaExceededRegex, detail); ok {
			return NewWrapErrorWithSuggestions(err, detail, QuotaExceededSuggestions)
		} else if detail != "" {
			return Wrap(err, strings.TrimSuffix(detail, "\n"))
		}
	}
	if resBody.Message != "" {
		return Wrap(err, strings.TrimRight(resBody.Message, "\n"))
	}

	return err
}

func CatchResourceNotFoundError(err error, resourceId string) error {
	if err == nil {
		return nil
	}

	if _, ok := err.(*KafkaClusterNotFoundError); ok || isResourceNotFoundError(err) {
		errorMsg := fmt.Sprintf(ResourceNotFoundErrorMsg, resourceId)
		suggestionsMsg := fmt.Sprintf(ResourceNotFoundSuggestions, resourceId)
		return NewErrorWithSuggestions(errorMsg, suggestionsMsg)
	}

	return err
}

func CatchEnvironmentNotFoundError(err error, r *http.Response) error {
	if err == nil {
		return nil
	}

	if r != nil && r.StatusCode == http.StatusForbidden {
		return NewWrapErrorWithSuggestions(CatchCCloudV2Error(err, r), "environment not found or access forbidden", EnvNotFoundSuggestions)
	}

	return CatchCCloudV2Error(err, r)
}

func CatchKafkaNotFoundError(err error, clusterId string, r *http.Response) error {
	if err == nil {
		return nil
	}
	if isResourceNotFoundError(err) {
		return &KafkaClusterNotFoundError{ClusterID: clusterId}
	}

	if r != nil && r.StatusCode == http.StatusForbidden {
		suggestions := ChooseRightEnvironmentSuggestions
		if r.Request.Method == http.MethodDelete {
			suggestions = KafkaClusterDeletingSuggestions
		}
		return NewWrapErrorWithSuggestions(CatchCCloudV2Error(err, r), "Kafka cluster not found or access forbidden", suggestions)
	}

	return CatchCCloudV2Error(err, r)
}

func CatchClusterConfigurationNotValidError(err error, r *http.Response) error {
	if err == nil {
		return nil
	}

	if r == nil {
		return err
	}

	err = CatchCCloudV2Error(err, r)
	if strings.Contains(err.Error(), "CKU must be greater") {
		return New(InvalidCkuErrorMsg)
	}

	return err
}

func CatchApiKeyForbiddenAccessError(err error, operation string, r *http.Response) error {
	if r != nil && r.StatusCode == http.StatusForbidden || strings.Contains(err.Error(), "Unknown API key") {
		return NewWrapErrorWithSuggestions(CatchCCloudV2Error(err, r), fmt.Sprintf("error %s API key", operation), APIKeyNotFoundSuggestions)
	}
	return CatchCCloudV2Error(err, r)
}

func CatchKSQLNotFoundError(err error, clusterId string) error {
	if err == nil {
		return nil
	}
	if isResourceNotFoundError(err) {
		errorMsg := fmt.Sprintf(ResourceNotFoundErrorMsg, clusterId)
		return NewErrorWithSuggestions(errorMsg, KSQLNotFoundSuggestions)
	}
	return err
}

func CatchServiceNameInUseError(err error, r *http.Response, serviceName string) error {
	if err == nil {
		return nil
	}

	if r == nil {
		return err
	}

	err = CatchCCloudV2Error(err, r)
	if strings.Contains(err.Error(), "Service name is already in use") {
		errorMsg := fmt.Sprintf(ServiceNameInUseErrorMsg, serviceName)
		return NewErrorWithSuggestions(errorMsg, ServiceNameInUseSuggestions)
	}

	return err
}

func CatchServiceAccountNotFoundError(err error, r *http.Response, serviceAccountId string) error {
	if err == nil {
		return nil
	}

	if r != nil {
		switch r.StatusCode {
		case http.StatusNotFound:
			errorMsg := fmt.Sprintf(ServiceAccountNotFoundErrorMsg, serviceAccountId)
			return NewErrorWithSuggestions(errorMsg, ServiceAccountNotFoundSuggestions)
		case http.StatusForbidden:
			return NewWrapErrorWithSuggestions(CatchCCloudV2Error(err, r), "service account not found or access forbidden", ServiceAccountNotFoundSuggestions)
		}
	}

	return CatchCCloudV2Error(err, r)
}

func isResourceNotFoundError(err error) bool {
	return strings.Contains(err.Error(), "resource not found")
}

/*
Error: 1 error occurred:
	* error creating topic bob: Topic 'bob' already exists.
*/
func CatchTopicExistsError(err error, clusterId string, topicName string, ifNotExistsFlag bool) error {
	if err == nil {
		return nil
	}
	compiledRegex := regexp.MustCompile(`error creating topic .*: Topic '.*' already exists\.`)
	if compiledRegex.MatchString(err.Error()) {
		if ifNotExistsFlag {
			return nil
		}
		errorMsg := fmt.Sprintf(TopicExistsErrorMsg, topicName, clusterId)
		suggestions := fmt.Sprintf(TopicExistsSuggestions, clusterId, clusterId)
		return NewErrorWithSuggestions(errorMsg, suggestions)
	}
	return err
}

/*
failed to produce offset -1: Unknown error, how did this happen? Error code = 87
*/
func CatchProduceToCompactedTopicError(err error, topicName string) (bool, error) {
	if err == nil {
		return false, nil
	}
	compiledRegex := regexp.MustCompile(`Unknown error, how did this happen\? Error code = 87`)
	if compiledRegex.MatchString(err.Error()) {
		errorMsg := fmt.Sprintf(ProducingToCompactedTopicErrorMsg, topicName)
		return true, NewErrorWithSuggestions(errorMsg, ProducingToCompactedTopicSuggestions)
	}
	return false, err
}

/*
Error: 1 error occurred:
	* error listing topics: Authentication failed: 1 extensions are invalid! They are: logicalCluster: Authentication failed
Error: 1 error occurred:
	* error creating topic test-topic: Authentication failed: 1 extensions are invalid! They are: logicalCluster: Authentication failed
*/
func CatchClusterNotReadyError(err error, clusterId string) error {
	if err == nil {
		return nil
	}
	if strings.Contains(err.Error(), "Authentication failed: 1 extensions are invalid! They are: logicalCluster: Authentication failed") {
		errorMsg := fmt.Sprintf(KafkaNotReadyErrorMsg, clusterId)
		return NewErrorWithSuggestions(errorMsg, KafkaNotReadySuggestions)
	}
	return err
}

func CatchSchemaNotFoundError(err error, r *http.Response) error {
	if err == nil {
		return nil
	}

	if r == nil {
		return err
	}

	if strings.Contains(r.Status, "Not Found") {
		return NewErrorWithSuggestions(SchemaNotFoundErrorMsg, SchemaNotFoundSuggestions)
	}

	return err
}

func CatchNoSubjectLevelConfigError(err error, r *http.Response, subject string) error {
	if err == nil {
		return nil
	}

	if r == nil {
		return err
	}

	if strings.Contains(r.Status, "Not Found") {
		return errors.New(fmt.Sprintf(NoSubjectLevelConfigErrorMsg, subject))
	}

	return err
}
