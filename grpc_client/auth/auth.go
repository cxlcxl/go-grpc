package auth

import "context"

type RpcAuth struct {
	Username string
	Pass     string
}

func (ra *RpcAuth) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{"user": ra.Username, "pass": ra.Pass}, nil
}

func (ra *RpcAuth) RequireTransportSecurity() bool {
	return false
}
