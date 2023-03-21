package jwt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	jwtGo "github.com/dgrijalva/jwt-go"
)

/**
Example:

rsa := rdmJwt.RSA{}
rsa.SetPrivateKey(string(privateKey)).SetPublicKey(string(publicKey))
signedToken, err := rsa.RsaSignClaims(claims)   // sign claim
goerror.Fatal(err)

mapClaims, err := rsa.RsaParseClaims(signedToken)	// parse claim
goerror.Fatal(err)
fmt.Printf("%#v", mapClaims)
**/

type RSA struct {
	// PrivatePem is private key at PEM format
	privatePem string
	//  PublicPem is public key at PEM format
	publicPem string
}

func (r *RSA) GetPrivateKey() string {
	return r.privatePem
}

func (r *RSA) GetPublicKey() string {
	return r.publicPem
}

func (r *RSA) SetPrivateKey(privatePem string) *RSA {
	r.privatePem = privatePem
	return r
}

func (r *RSA) SetPublicKey(publicPem string) *RSA {
	r.publicPem = publicPem
	return r
}

// RsaGenerateKey 產生並回傳 RS256 PEM 格式的 private key 到結構中. see https://pkg.go.dev/github.com/kataras/jwt#readme-generate-keys
func (r *RSA) RsaGeneratePemKey() error {
	// 初始化
	r.SetPrivateKey("")
	r.SetPublicKey("")

	// Generate HMAC
	sharedKey := make([]byte, 32)
	_, _ = rand.Read(sharedKey)

	bitSize := 2048
	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return err
	}

	privPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		},
	)

	publiceKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		panic(err)
	}

	pubPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: publiceKeyBytes,
		},
	)

	r.SetPrivateKey(string(privPem))
	r.SetPublicKey(string(pubPem))
	return nil
}

// RsaSignClaims 使用傳入的 PEM 格式 privatePemKey, 加密 claim 字串 (需要加上 json tag), 得到 signed token
func (r *RSA) RsaSignClaims(claims jwt.Claims) (string, error) {
	privateRSA, err := jwtGo.ParseRSAPrivateKeyFromPEM([]byte(r.GetPrivateKey()))
	if err != nil {
		return "", err
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedtoken, err := tokenClaims.SignedString(privateRSA)
	return signedtoken, nil
}

// RsaParseClaims 解析傳入的 jwt signed token, 利用 public key 解碼回傳 claims (可對回傳的 map claim 進行 json unmarshal 或是直接判斷)
func (r *RSA) RsaParseClaims(signedToken string) (*jwt.MapClaims, error) {

	token, err := jwt.ParseWithClaims(signedToken, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Expected token algorithm '%v' but got '%v'",
				jwt.SigningMethodRS256.Name,
				token.Method.Alg(),
			)
		}

		publicRSA, err := jwt.ParseRSAPublicKeyFromPEM([]byte(r.GetPublicKey()))
		if err != nil {
			return nil, err
		}

		return publicRSA, nil
	})

	if err != nil {
		return nil, err
	}

	return token.Claims.(*jwt.MapClaims), nil
}

// deprecated 加密: not a jwt format, todo remove
// RsaEncryptPayload 使用傳入的 PEM 格式 privatePemKey, 加密 payload 字串, 得到 signature
func (r *RSA) RsaSignPayload(payload string) (string, error) {

	rsaPrivKey, err := jwtGo.ParseRSAPrivateKeyFromPEM([]byte(r.GetPrivateKey()))
	if err != nil {
		return "", err
	}

	signature, err := jwtGo.SigningMethodRS256.Sign(payload, rsaPrivKey)
	if err != nil {
		return "", err
	}

	return signature, nil
}

// deprecated 加密: not a jwt format, todo remove
// RsaVerifyPayload 驗證加密後的 signature 是否與原始資料 payload 一致,
func RsaVerifyPayload(payload string, signature string, publicPemKey string) error {

	rsaPubKey, err := jwtGo.ParseRSAPublicKeyFromPEM([]byte(publicPemKey))
	if err != nil {
		return err
	}

	err = jwtGo.SigningMethodRS256.Verify(payload, signature, rsaPubKey)
	if err != nil {
		return err
	}

	return nil
}
