package credentials

import (
	"errors"
	"os"

	"github.com/rancher/muchang/credentials/providers"
	"github.com/rancher/muchang/utils/tea"
)

// Credential is an interface for getting actual credential
type Credential interface {
	GetType() *string
	GetCredential() (*CredentialModel, error)
}

// Config is important when call NewCredential
type Config struct {
	// Credential type, including access_key, sts, bearer, ecs_ram_role, ram_role_arn, rsa_key_pair, oidc_role_arn, credentials_uri
	Type            *string `json:"type"`
	AccessKeyId     *string `json:"access_key_id"`
	AccessKeySecret *string `json:"access_key_secret"`
	SecurityToken   *string `json:"security_token"`
	BearerToken     *string `json:"bearer_token"`

	// Used when the type is ram_role_arn or oidc_role_arn
	OIDCProviderArn       *string `json:"oidc_provider_arn"`
	OIDCTokenFilePath     *string `json:"oidc_token"`
	RoleArn               *string `json:"role_arn"`
	RoleSessionName       *string `json:"role_session_name"`
	RoleSessionExpiration *int    `json:"role_session_expiration"`
	Policy                *string `json:"policy"`
	ExternalId            *string `json:"external_id"`
	STSEndpoint           *string `json:"sts_endpoint"`

	// Used when the type is ecs_ram_role
	RoleName      *string `json:"role_name"`
	DisableIMDSv1 *bool   `json:"disable_imds_v1"`

	// Used when the type is credentials_uri
	Url *string `json:"url"`

	PublicKeyId    *string `json:"public_key_id"`
	PrivateKeyFile *string `json:"private_key_file"`
	Host           *string `json:"host"`

	// Read timeout, in milliseconds.
	// The default value for ecs_ram_role is 1000ms, the default value for ram_role_arn is 5000ms, and the default value for oidc_role_arn is 5000ms.
	Timeout *int `json:"timeout"`
	// Connection timeout, in milliseconds.
	// The default value for ecs_ram_role is 1000ms, the default value for ram_role_arn is 10000ms, and the default value for oidc_role_arn is 10000ms.
	ConnectTimeout *int `json:"connect_timeout"`

	Proxy          *string  `json:"proxy"`
	InAdvanceScale *float64 `json:"inAdvanceScale"`
}

func (s Config) String() string {
	return tea.Prettify(s)
}

func (s Config) GoString() string {
	return s.String()
}

func (s *Config) SetAccessKeyId(v string) *Config {
	s.AccessKeyId = &v
	return s
}

func (s *Config) SetAccessKeySecret(v string) *Config {
	s.AccessKeySecret = &v
	return s
}

func (s *Config) SetSecurityToken(v string) *Config {
	s.SecurityToken = &v
	return s
}

func (s *Config) SetRoleArn(v string) *Config {
	s.RoleArn = &v
	return s
}

func (s *Config) SetRoleSessionName(v string) *Config {
	s.RoleSessionName = &v
	return s
}

func (s *Config) SetPublicKeyId(v string) *Config {
	s.PublicKeyId = &v
	return s
}

func (s *Config) SetRoleName(v string) *Config {
	s.RoleName = &v
	return s
}

func (s *Config) SetPrivateKeyFile(v string) *Config {
	s.PrivateKeyFile = &v
	return s
}

func (s *Config) SetBearerToken(v string) *Config {
	s.BearerToken = &v
	return s
}

func (s *Config) SetRoleSessionExpiration(v int) *Config {
	s.RoleSessionExpiration = &v
	return s
}

func (s *Config) SetPolicy(v string) *Config {
	s.Policy = &v
	return s
}

func (s *Config) SetHost(v string) *Config {
	s.Host = &v
	return s
}

func (s *Config) SetTimeout(v int) *Config {
	s.Timeout = &v
	return s
}

func (s *Config) SetConnectTimeout(v int) *Config {
	s.ConnectTimeout = &v
	return s
}

func (s *Config) SetProxy(v string) *Config {
	s.Proxy = &v
	return s
}

func (s *Config) SetType(v string) *Config {
	s.Type = &v
	return s
}

func (s *Config) SetOIDCTokenFilePath(v string) *Config {
	s.OIDCTokenFilePath = &v
	return s
}

func (s *Config) SetOIDCProviderArn(v string) *Config {
	s.OIDCProviderArn = &v
	return s
}

func (s *Config) SetURLCredential(v string) *Config {
	if v == "" {
		v = os.Getenv("ALIBABA_CLOUD_CREDENTIALS_URI")
	}
	s.Url = &v
	return s
}

func (s *Config) SetSTSEndpoint(v string) *Config {
	s.STSEndpoint = &v
	return s
}

func (s *Config) SetExternalId(v string) *Config {
	s.ExternalId = &v
	return s
}

// NewCredential return a credential according to the type in config.
// if config is nil, the function will use default provider chain to get credentials.
// please see README.md for detail.
func NewCredential(config *Config) (credential Credential, err error) {
	if config == nil {
		return nil, errors.New("invalid credentials")
	}
	switch tea.StringValue(config.Type) {
	case "access_key":
		provider, err := providers.NewStaticAKCredentialsProviderBuilder().
			WithAccessKeyId(tea.StringValue(config.AccessKeyId)).
			WithAccessKeySecret(tea.StringValue(config.AccessKeySecret)).
			Build()
		if err != nil {
			return nil, err
		}
		credential = FromCredentialsProvider("access_key", provider)
	case "sts":
		provider, err := providers.NewStaticSTSCredentialsProviderBuilder().
			WithAccessKeyId(tea.StringValue(config.AccessKeyId)).
			WithAccessKeySecret(tea.StringValue(config.AccessKeySecret)).
			WithSecurityToken(tea.StringValue(config.SecurityToken)).
			Build()
		if err != nil {
			return nil, err
		}

		credential = FromCredentialsProvider("sts", provider)
	case "bearer":
		if tea.StringValue(config.BearerToken) == "" {
			err = errors.New("BearerToken cannot be empty")
			return
		}
		credential = newBearerTokenCredential(tea.StringValue(config.BearerToken))
	default:
		err = errors.New("invalid type option, support: access_key, sts, bearer")
		return
	}
	return credential, nil
}

type credentialsProviderWrap struct {
	typeName string
	provider providers.CredentialsProvider
}

// Get credentials
func (cp *credentialsProviderWrap) GetCredential() (cm *CredentialModel, err error) {
	c, err := cp.provider.GetCredentials()
	if err != nil {
		return
	}

	cm = &CredentialModel{
		AccessKeyId:     &c.AccessKeyId,
		AccessKeySecret: &c.AccessKeySecret,
		SecurityToken:   &c.SecurityToken,
		Type:            &cp.typeName,
		ProviderName:    &c.ProviderName,
	}
	return
}

func (cp *credentialsProviderWrap) GetType() *string {
	return &cp.typeName
}

func FromCredentialsProvider(typeName string, cp providers.CredentialsProvider) Credential {
	return &credentialsProviderWrap{
		typeName: typeName,
		provider: cp,
	}
}
