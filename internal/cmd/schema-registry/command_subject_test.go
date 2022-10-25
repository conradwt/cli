package schemaregistry

import (
	"context"
	"net/http"
	"testing"

	"github.com/confluentinc/ccloud-sdk-go-v1"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	schedv1 "github.com/confluentinc/cc-structs/kafka/scheduler/v1"
	"github.com/confluentinc/ccloud-sdk-go-v1/mock"
	srsdk "github.com/confluentinc/schema-registry-sdk-go"
	srMock "github.com/confluentinc/schema-registry-sdk-go/mock"

	v1 "github.com/confluentinc/cli/internal/pkg/config/v1"
	cliMock "github.com/confluentinc/cli/mock"
)

const (
	subjectName = "Subject"
)

type SubjectTestSuite struct {
	suite.Suite
	conf             *v1.Config
	kafkaCluster     *schedv1.KafkaCluster
	srCluster        *schedv1.SchemaRegistryCluster
	srMothershipMock *mock.SchemaRegistry
	srClientMock     *srsdk.APIClient
}

func (suite *SubjectTestSuite) SetupSuite() {
	suite.conf = v1.AuthenticatedCloudConfigMock()
	ctx := suite.conf.Context()
	srCluster := ctx.SchemaRegistryClusters[ctx.GetEnvironment().GetId()]
	srCluster.SrCredentials = &v1.APIKeyPair{Key: "key", Secret: "secret"}
	cluster := ctx.KafkaClusterContext.GetActiveKafkaClusterConfig()
	suite.kafkaCluster = &schedv1.KafkaCluster{
		Id:         cluster.ID,
		Name:       cluster.Name,
		Endpoint:   cluster.APIEndpoint,
		Enterprise: true,
	}
	suite.srCluster = &schedv1.SchemaRegistryCluster{
		Id: srClusterID,
	}
}

func (suite *SubjectTestSuite) SetupTest() {
	suite.srMothershipMock = &mock.SchemaRegistry{
		CreateSchemaRegistryClusterFunc: func(ctx context.Context, clusterConfig *schedv1.SchemaRegistryClusterConfig) (*schedv1.SchemaRegistryCluster, error) {
			return suite.srCluster, nil
		},
		GetSchemaRegistryClusterFunc: func(ctx context.Context, clusterConfig *schedv1.SchemaRegistryCluster) (*schedv1.SchemaRegistryCluster, error) {
			return nil, nil
		},
	}

	suite.srClientMock = &srsdk.APIClient{
		DefaultApi: &srMock.DefaultApi{
			ListFunc: func(ctx context.Context, opts *srsdk.ListOpts) ([]string, *http.Response, error) {
				return []string{"subject 1", "subject 2"}, nil, nil
			},
			ListVersionsFunc: func(ctx context.Context, subject string, opts *srsdk.ListVersionsOpts) (int32s []int32, response *http.Response, e error) {
				return []int32{1234, 4567}, nil, nil
			},
			UpdateSubjectLevelConfigFunc: func(ctx context.Context, subject string, body srsdk.ConfigUpdateRequest) (request srsdk.ConfigUpdateRequest, response *http.Response, e error) {
				return srsdk.ConfigUpdateRequest{Compatibility: body.Compatibility}, nil, nil
			},
			UpdateModeFunc: func(ctx context.Context, subject string, body srsdk.ModeUpdateRequest) (request srsdk.ModeUpdateRequest, response *http.Response, e error) {
				return srsdk.ModeUpdateRequest{Mode: body.Mode}, nil, nil
			},
		},
	}
}

func (suite *SubjectTestSuite) newCMD() *cobra.Command {
	client := &ccloud.Client{
		SchemaRegistry: suite.srMothershipMock,
	}
	cmd := New(suite.conf, cliMock.NewPreRunnerMock(client, nil, nil, nil, suite.conf), suite.srClientMock)
	return cmd
}

func (suite *SubjectTestSuite) TestSubjectList() {
	cmd := suite.newCMD()
	cmd.SetArgs([]string{"subject", "list"})
	err := cmd.Execute()
	req := require.New(suite.T())
	req.Nil(err)
	apiMock, _ := suite.srClientMock.DefaultApi.(*srMock.DefaultApi)
	req.True(apiMock.ListCalled())
}

func (suite *SubjectTestSuite) TestSubjectListDeleted() {
	cmd := suite.newCMD()
	cmd.SetArgs([]string{"subject", "list", "--deleted"})
	err := cmd.Execute()
	req := require.New(suite.T())
	req.Nil(err)
	apiMock, _ := suite.srClientMock.DefaultApi.(*srMock.DefaultApi)
	req.True(apiMock.ListCalled())
	retVal := apiMock.ListCalls()[0]
	req.Equal(retVal.LocalVarOptionals.Deleted.Value(), true)
}

func (suite *SubjectTestSuite) TestSubjectUpdateMode() {
	cmd := suite.newCMD()
	cmd.SetArgs([]string{"subject", "update", subjectName, "--mode", "READWRITE"})
	err := cmd.Execute()
	req := require.New(suite.T())
	req.Nil(err)
	apiMock, _ := suite.srClientMock.DefaultApi.(*srMock.DefaultApi)
	req.False(apiMock.UpdateTopLevelModeCalled())
	req.True(apiMock.UpdateModeCalled())
	retVal := apiMock.UpdateModeCalls()[0]
	req.Equal(retVal.Subject, subjectName)
}

func (suite *SubjectTestSuite) TestSubjectUpdateCompatibility() {
	cmd := suite.newCMD()
	cmd.SetArgs([]string{"subject", "update", subjectName, "--compatibility", "BACKWARD"})
	err := cmd.Execute()
	req := require.New(suite.T())
	req.Nil(err)
	apiMock, _ := suite.srClientMock.DefaultApi.(*srMock.DefaultApi)
	req.True(apiMock.UpdateSubjectLevelConfigCalled())
	retVal := apiMock.UpdateSubjectLevelConfigCalls()[0]
	req.Equal(retVal.Subject, subjectName)
}

func (suite *SubjectTestSuite) TestSubjectUpdateNoArgs() {
	cmd := suite.newCMD()
	cmd.SetArgs([]string{"subject", "update", subjectName})
	err := cmd.Execute()
	req := require.New(suite.T())
	req.Error(err, "Error: flag string not set")
}

func (suite *SubjectTestSuite) TestSubjectDescribe() {
	cmd := suite.newCMD()
	cmd.SetArgs([]string{"subject", "describe", subjectName})
	err := cmd.Execute()
	req := require.New(suite.T())
	req.Nil(err)
	apiMock, _ := suite.srClientMock.DefaultApi.(*srMock.DefaultApi)
	req.True(apiMock.ListVersionsCalled())
	retVal := apiMock.ListVersionsCalls()[0]
	req.Equal(retVal.Subject, subjectName)
}

func (suite *SubjectTestSuite) TestSubjectDescribeDeleted() {
	cmd := suite.newCMD()
	cmd.SetArgs([]string{"subject", "describe", subjectName, "--deleted"})
	err := cmd.Execute()
	req := require.New(suite.T())
	req.Nil(err)
	apiMock, _ := suite.srClientMock.DefaultApi.(*srMock.DefaultApi)
	req.True(apiMock.ListVersionsCalled())
	retVal := apiMock.ListVersionsCalls()[0]
	req.Equal(retVal.Subject, subjectName)
	req.Equal(retVal.LocalVarOptionals.Deleted.Value(), true)
}

func TestSubjectSuite(t *testing.T) {
	suite.Run(t, new(SubjectTestSuite))
}
