package an

import (
	"context"
	"crypto/x509"
	"nbltest/util"
	"sync"
	"time"

	"bitbucket.org/Axxonsoft/bl/auth"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"

	"google.golang.org/grpc"
)

const anonKey = `
-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAzLhprNSyvjcuiw166J15fBRkQU1AAHS2PHp80I90z62GMPL+
cS3K5QxudC2OaWnxZb2gbcpBNFzh7+tJUEzq/wI7lRKaKtiGtB1FaAKumPghkhd2
/Tr8VTCrCY6FgJlTKqkcAv78VKeBBNonGpccN8nUbGvXJMXSzAk7bImwDRQXNBCv
uRG0FQtmE2oAoEWLXG0gH8EW1QztrFy9fHN/eR+Rcwzic4T4dvBYHZZNBX3VjPOi
u6OIK11DqrGJi3e9TN1QUcN4C9dpqldxPgvx9wgFncog9jD7HwJQ/DE/IMAoVHwZ
410KsN+o4tLJjkTVsEcW7rkbbJaBHfhD+yqmqQIDAQABAoIBAGwY0Avnf746ywO3
iXe8dwJSjiGBFdNnzTYgAznpFef3G01LcZj3chQBvWzxBSqaO1HakBSI4GcyyEZz
+ZFCHC+s0SNE4EzRakc+0YA7MlApbSiD5VuPQuAEw7tXGx0tWKHFC7p1Q227yL90
JYO/2LGIi/b39nE/V7BPC6ajCWfNizFNcpgcPWgwonf73j0mG/kbD8R9Cazi/7J9
MZ60t2US/Ge3vJMlN3SBrXWNIUYVtCrVmeArEbcWohcEXKpQu3QlCro89r2pOyh+
rXCiVu4hRmDj/Z6CzOymMDbvQP9imNuBwMATeJaP8eTsW2R1tGte4msISKtW9Ua1
JPKiGRECgYEA5vZO5e9tner8ugJ9etPgYcOuvfCpCyDFHBuR5tnvPUY3JyNL2AHW
99LUDB96TZdowK2Vdbf4wtGoNDK5tPtrI5LheuAU1UDs75kYFjWBCC90QHcZsjYj
iZhA/4Zv8g618MB7DgWjzedgX4m8hK6fCwYMBqyC7sBhC0HZmZY3XfUCgYEA4unW
25y3K0RXX2IzZipYJh48fKgSOPMXDjlHCA1+TXmAhHM3y4yOrut9nWVipddjzKTh
/dVA+8wAFHb3TrPMutUbkjfqTJMIiQrhFC69LUnIsKaSYQu6iIuA8Zspg5ZcCOHz
LldMBs4nY+KBlhKtP9/vslCeBtQV6WaYetueIWUCgYEA4lhzHA6YS9I2WYkFNjGc
zdL7VnQbBqfX5GL0itv8BP3iIT4LHyc2aGs6mqLitlXzOBklx1dDuJHFmVo4+zAG
YLeauFQQtqnJSkqr+1/2E3KGKINQMIG0NC19Ta6P2RYnokjQj/5g+PKPVTHkCdgk
T6ZedM/uBVOOB31CZq17l10CgYEAkLok2B20llwYfjXcIqWPY4CVg8yPEtO5ONE/
hNtWW4PLfK8gPyt/NgHtNJ6dtLkUZkoj+goyUGdlBah7RC9ZvTB/TVtsjTqajw3p
UT4eWuxcnI8hfdRFPSH9NByK9erI+MFkoYH8c7q5VPP7QwTWi48BCvJwWFqdPyi4
yGObECECgYEAwQcfjhMPvDCMYyMvAyCTTH4hYc87Lk6r+xfwLkOJFSp5wNE94105
EuObIzr8rTZ8sjwHfMZkdaH2Pk/b34TxMBVXFuQBte5Nn5+vHaV8mrlClkf+8Us6
tCikE6sSudHeAi2nr2B4ngwcMieAqZyDeqt6sEzCFEAZSW2F8quEGXg=
-----END RSA PRIVATE KEY-----
`

const anonCrt = `
-----BEGIN CERTIFICATE-----
MIIDaTCCAlGgAwIBAgIJAJencyQhQ8qbMA0GCSqGSIb3DQEBCwUAMEoxCzAJBgNV
BAYTAkdCMQ8wDQYDVQQHDAZMb25kb24xFjAUBgNVBAoMDUF4eG9uU29mdCBMdGQx
EjAQBgNVBAMMCUFub255bW91czAgFw0xNzExMDQxNzQ2MjRaGA8yMTE3MTAxMTE3
NDYyNFowSjELMAkGA1UEBhMCR0IxDzANBgNVBAcMBkxvbmRvbjEWMBQGA1UECgwN
QXh4b25Tb2Z0IEx0ZDESMBAGA1UEAwwJQW5vbnltb3VzMIIBIjANBgkqhkiG9w0B
AQEFAAOCAQ8AMIIBCgKCAQEAzLhprNSyvjcuiw166J15fBRkQU1AAHS2PHp80I90
z62GMPL+cS3K5QxudC2OaWnxZb2gbcpBNFzh7+tJUEzq/wI7lRKaKtiGtB1FaAKu
mPghkhd2/Tr8VTCrCY6FgJlTKqkcAv78VKeBBNonGpccN8nUbGvXJMXSzAk7bImw
DRQXNBCvuRG0FQtmE2oAoEWLXG0gH8EW1QztrFy9fHN/eR+Rcwzic4T4dvBYHZZN
BX3VjPOiu6OIK11DqrGJi3e9TN1QUcN4C9dpqldxPgvx9wgFncog9jD7HwJQ/DE/
IMAoVHwZ410KsN+o4tLJjkTVsEcW7rkbbJaBHfhD+yqmqQIDAQABo1AwTjAdBgNV
HQ4EFgQUdRGQ83k1T57/uaFMswDyHGfjxHQwHwYDVR0jBBgwFoAUdRGQ83k1T57/
uaFMswDyHGfjxHQwDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAQEAvKNb
lR2hvMRRTZ/1hkORxKAHbglaxtOrTzMyEGLUyWUsliMt4vtXpeagivvQxbI1dzOm
y/1/hi6kGQ/aMSV8sw23H2mbDY7jgPTAvn6++ly2mStfpPd8gGhg2zdWryGVLXih
oNX40yASO2yDIDsbrsoNfu21vEhfK2WPG4xb1bk5Y+NCHXf19pBGm0qUrxg75IBH
C8SwFvPYezxCW68QuLCtPuG8nm7R22j1Gexz4cHrGvj45MRe/Ez73Zfu/uReIyj2
bnSNNu86vEk9gIXfAjuWOWivfBmfYV6pa718cWTClW2z2xlfTrW9Ft8hCqfJ3zUZ
l2hHTkJd4NqIUHblMQ==
-----END CERTIFICATE-----
`

type tokenCredential struct {
	name  string
	value string
}

// GetRequestMetadata returns request metadata
func (tc tokenCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		tc.name: tc.value,
	}, nil
}

// RequireTransportSecurity returns transport security
func (tc tokenCredential) RequireTransportSecurity() bool {
	return true
}

// MakeConnection makes grpc connection to server
func MakeConnection(ctx context.Context, serviceAddress, user, password string, connTimeout time.Duration) (conn *grpc.ClientConn, err error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, connTimeout)
	defer cancel()

	tokenName, tokenValue, err := authenticate(timeoutCtx, serviceAddress, user, password)
	if err != nil {
		return nil, err
	}
	tokenCreds := grpc.WithPerRPCCredentials(tokenCredential{tokenName, tokenValue})
	keepaliveOpt := grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:                20 * time.Second,
		PermitWithoutStream: true,
	})
	return grpc.DialContext(timeoutCtx, serviceAddress, certOpt(), tokenCreds, keepaliveOpt, grpc.WithBlock())
}

var certInit sync.Once
var transportCredentials credentials.TransportCredentials

func certOpt() grpc.DialOption {
	certInit.Do(func() {
		cp := x509.NewCertPool()
		util.CheckOk(cp.AppendCertsFromPEM([]byte(anonCrt)), "credentials: failed to append certificates")
		transportCredentials = credentials.NewClientTLSFromCert(cp, "Anonymous")
	})

	return grpc.WithTransportCredentials(transportCredentials)
}

func authenticate(ctx context.Context, address, user, password string) (string, string, error) {
	conn, err := grpc.DialContext(ctx, address, certOpt())
	if err != nil {
		return "", "", err
	}
	defer conn.Close()

	client := auth.NewAuthenticationServiceClient(conn)
	req := auth.AuthenticateRequest{UserName: user, Password: password}
	resp, err := client.Authenticate(ctx, &req)
	if err != nil {
		return "", "", err
	}
	return resp.GetTokenName(), resp.GetTokenValue(), nil
}
