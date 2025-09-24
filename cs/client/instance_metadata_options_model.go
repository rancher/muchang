package client

import (
	"github.com/rancher/muchang/utils/tea/dara"
)

type iInstanceMetadataOptions interface {
	dara.Model
	String() string
	GoString() string
	SetHttpTokens(v string) *InstanceMetadataOptions
	GetHttpTokens() *string
}

type InstanceMetadataOptions struct {
	// example:
	//
	// optional
	HttpTokens *string `json:"http_tokens,omitempty" xml:"http_tokens,omitempty"`
}

func (s InstanceMetadataOptions) String() string {
	return dara.Prettify(s)
}

func (s InstanceMetadataOptions) GoString() string {
	return s.String()
}

func (s *InstanceMetadataOptions) GetHttpTokens() *string {
	return s.HttpTokens
}

func (s *InstanceMetadataOptions) SetHttpTokens(v string) *InstanceMetadataOptions {
	s.HttpTokens = &v
	return s
}

func (s *InstanceMetadataOptions) Validate() error {
	return dara.Validate(s)
}
