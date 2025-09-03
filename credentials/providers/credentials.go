package providers

// The credentials struct
type Credentials struct {
	AccessKeyId     string
	AccessKeySecret string
	SecurityToken   string
	ProviderName    string
}

// The credentials provider interface, return credentials and provider name
type CredentialsProvider interface {
	// Get credentials
	GetCredentials() (*Credentials, error)
	// Get credentials provider name
	GetProviderName() string
}
