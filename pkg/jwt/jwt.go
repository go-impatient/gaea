package jwt

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"

	"moocss.com/gaea/pkg/base64"
	"moocss.com/gaea/pkg/conf"
	"moocss.com/gaea/pkg/errors"
)

var (
	SignKey     string = "Z3d0X3NpZ25fa2V5" // gwt_sign_key
	ExpiresTime int    = 259200

	_ JWT = (*jwt)(nil)
)

// Structured version of Claims Section, as referenced at
// https://tools.ietf.org/html/rfc7519#section-4.1
// See examples for how to use this with your own claim types
type CustomClaims struct {
	Data json.RawMessage
	jwtgo.StandardClaims
}

// JWT is a jwt interface.
type JWT interface {
	// 创建Token
	CreateToken(claims CustomClaims) (string, error)
	// 解析Token
	ParseToken(tokenString string) (*CustomClaims, error)
	// 更新Token
	RefreshToken(tokenString string) (string, error)

	// 验证Token
	TokenValid(tokenString string) error
	// 验证Token是否过期
	TokenExpires(tokenString string) (jwtgo.MapClaims, bool, error)

	GetSigningKey() string
}

type jwt struct {
	SigningKey  []byte
	ExpiresTime int64
}

// New new a jwt with options.
func New() JWT {
	signingKey := getSignKey()
	expiresTime := getExpiresTime()
	return &jwt{
		SigningKey:     signingKey,
		ExpiresTime: expiresTime,
	}
}

// 获取base64编码过的 signing key
func (j *jwt) GetSigningKey() string {
	return base64.Base64UrlEncode(j.SigningKey)
}

func (j *jwt) CreateToken(claims CustomClaims) (string, error) {
	return j.createTokenString(claims)
}

func (j *jwt) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := j.verifyToken(tokenString)
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		// 控制台打印信息
		pretty(claims)

		return claims, nil
	}
	return nil, errors.TokenInvalid
}

func (j *jwt) RefreshToken(tokenString string) (string, error) {
	jwtgo.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwtgo.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwtgo.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwtgo.TimeFunc = time.Now
		return j.CreateToken(*claims)
	}
	return "", errors.TokenInvalid
}

// TokenValid 验证token
func (j *jwt) TokenValid(tokenString string) error {
	token, err := j.verifyToken(tokenString)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(*CustomClaims); !ok && !token.Valid {
		return err
	}
	return nil
}

// TokenExpires 验证Token是否过期, 三种状态：未过期、过期、过期但可以Refresh
func (j *jwt) TokenExpires(tokenString string) (jwtgo.MapClaims, bool, error) {
	token, err := j.verifyToken(tokenString)
	if err != nil {
		return nil, false, err
	}
	claims, ok := token.Claims.(jwtgo.MapClaims)
	if ok && token.Valid {
		return claims, false, errors.TokenInvalid
	}

	// ExpiresAt
	expiresAt := int64(claims["exp"].(float64))
	if expiresAt > time.Now().Unix() {
		return claims, false, nil
	}

	// IssuedAt
	issuedAt := int64(claims["iat"].(float64))
	if issuedAt > time.Now().Unix() {
		return claims, true, nil
	}

	return nil, false, nil
}

func (j *jwt) createTokenString(claims CustomClaims) (string, error) {
	var method jwtgo.SigningMethod
	signingMethod := conf.Get("JWT_SIGNING_METHOD")

	switch signingMethod {
	case "HS256":
		method = jwtgo.SigningMethodHS256
	case "HS384":
		method = jwtgo.SigningMethodHS384
	case "HS512":
		method = jwtgo.SigningMethodHS512
	default:
		method = jwtgo.SigningMethodHS256
	}

	claims.StandardClaims.ExpiresAt = j.ExpiresTime
	token := jwtgo.NewWithClaims(method, claims)
	tokenString, err := token.SignedString(j.SigningKey)
	if err != nil {
		return "", errors.TokenFailure
	}
	return tokenString, nil
}

func (j *jwt) verifyToken(tokenString string) (*jwtgo.Token, error) {
	if tokenString == "" {
		return nil, errors.TokenMalformed
	}

	token, err := jwtgo.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwtgo.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwtgo.SigningMethodHMAC); !ok {
			return nil, errors.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwtgo.ValidationError); ok {
			if ve.Errors&jwtgo.ValidationErrorMalformed != 0 {
				return nil, errors.TokenMalformed
			} else if ve.Errors&jwtgo.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, errors.TokenExpired
			} else if ve.Errors&jwtgo.ValidationErrorNotValidYet != 0 {
				return nil, errors.TokenNotValidYet
			} else {
				return nil, errors.TokenInvalid
			}
		}
	}

	return token, nil
}

// pretty display the claims licely in the terminal
func pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Printf("无法格式化数据: %v", err)
		return
	}

	fmt.Println(string(b))
}

// getSignKey 获取signing key
func getSignKey() []byte {
	signingKey := conf.Get("JWT_SIGNING_KEY")
	if len(signingKey) == 0 {
		signingKey = SignKey // 默认signing key
	}

	// 恢复到原始的 signing key 值, 例如: gwt_sign_key
	data, err:= base64.Base64UrlDecode(signingKey)
	if err != nil {
		return nil
	}
	return data
}

// getExpiresTime 获取过期时间
func getExpiresTime() int64 {
	jTime := conf.GetInt("JWT_EFFECTIVE_DURATION")
	if jTime == 0 {
		jTime = ExpiresTime
	}

	return time.Now().Add(time.Duration(jTime) * time.Second).Unix()
}