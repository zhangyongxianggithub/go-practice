package main

import (
	"errors"

	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
)

func NewClient(ak, sk, endpoint string) (*bce.BceClient, error) {
	if endpoint == "" {
		return nil, errors.New("endpoint should not be empty")
	}

	if ak == "" || sk == "" {
		return nil, errors.New("ak and sk should no be empty")
	}

	credentials, err := auth.NewBceCredentials(ak, sk)
	if err != nil {
		return nil, err
	}

	defaultSignOptions := &auth.SignOptions{
		HeadersToSign: auth.DEFAULT_HEADERS_TO_SIGN,
		ExpireSeconds: auth.DEFAULT_EXPIRE_SECONDS,
	}

	defaultConf := &bce.BceClientConfiguration{
		Endpoint:                  endpoint,
		Region:                    bce.DEFAULT_REGION,
		UserAgent:                 bce.DEFAULT_USER_AGENT,
		Credentials:               credentials,
		SignOption:                defaultSignOptions,
		Retry:                     bce.DEFAULT_RETRY_POLICY,
		ConnectionTimeoutInMillis: bce.DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS,
	}

	return NewClientWithConfig(defaultConf), nil
}

func NewClientWithConfig(config *bce.BceClientConfiguration) *bce.BceClient {
	signer := &auth.BceV1Signer{}
	return bce.NewBceClient(config, signer)
}

const HttpContentTypeHeader = "Content-Type"

const HttpApplicationJsonValue = "application/json"

func main() {
	// client, _ := NewClient("test-ak", "test-sk", "http://szth-ivip-qa-01.szth.baidu.com:8420")
	// builder := bce.NewRequestBuilder(client)
	// err := builder.
	// 	WithMethod(http.MethodPost).
	// 	WithHeader(HttpContentTypeHeader, HttpApplicationJsonValue).
	// 	WithURL(fmt.Sprintf(analyticsJobUriGroupFormat, req.WorkspaceID, req.ProjectName)).
	// 	WithBody(req).
	// 	WithResult(&resp).
	// 	Do()
}
