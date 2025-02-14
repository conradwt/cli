package v1

import (
	"regexp"

	"github.com/confluentinc/cli/internal/pkg/secret"
)

const (
	authTokenRegex        = `^[\w-]*\.[\w-]*\.[\w-]*$`
	authRefreshTokenRegex = `^(v1\..*)$`
)

type ContextState struct {
	Auth             *AuthConfig `json:"auth"`
	AuthToken        string      `json:"auth_token"`
	AuthRefreshToken string      `json:"auth_refresh_token"`
	Salt             []byte      `json:"salt,omitempty"`
	Nonce            []byte      `json:"nonce,omitempty"`
}

func (c *ContextState) DecryptContextStateAuthToken(ctxName string) error {
	reg := regexp.MustCompile(authTokenRegex)
	if !reg.MatchString(c.AuthToken) && c.AuthToken != "" && c.Salt != nil {
		decryptedAuthToken, err := secret.Decrypt(ctxName, c.AuthToken, c.Salt, c.Nonce)
		if err != nil {
			return err
		}
		c.AuthToken = decryptedAuthToken
	}

	return nil
}

func (c *ContextState) DecryptContextStateAuthRefreshToken(ctxName string) error {
	reg := regexp.MustCompile(authRefreshTokenRegex)
	if !reg.MatchString(c.AuthRefreshToken) && c.AuthRefreshToken != "" && c.Salt != nil {
		decryptedAuthRefreshToken, err := secret.Decrypt(ctxName, c.AuthRefreshToken, c.Salt, c.Nonce)
		if err != nil {
			return err
		}
		c.AuthRefreshToken = decryptedAuthRefreshToken
	}

	return nil
}
