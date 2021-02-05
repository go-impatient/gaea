package jwt

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestGetSignKey(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
		// TODO: Add test cases.
		{
			"default",
			[]byte("gwt_sign_key"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSignKey(); string(got) != string(tt.want) {
				t.Errorf("GetSignKey() = %v, want %v", got, tt.want)
			}
		})

	}
}

func TestJWT_CreateToken(t *testing.T) {
	type fields struct {
		SigningKey []byte
	}
	type args struct {
		claims CustomClaims
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"case 01",
			fields{SigningKey: []byte("gwt_sign_key")},
			args{CustomClaims{Data: []byte("123456")}},
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJEYXRhIjoxMjM0NTZ9.oQC5aJRtHlHkxBvOKNj6ne5FFUnznwO8hjdcoWClNjo",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &jwt{
				SigningKey: tt.fields.SigningKey,
			}
			got, err := j.CreateToken(tt.args.claims)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJWT_ParseToken(t *testing.T) {
	type fields struct {
		SigningKey []byte
	}
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CustomClaims
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "case 01",
			fields: fields{SigningKey: []byte("gwt_sign_key")},
			args:   args{tokenString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJEYXRhIjpudWxsfQ.AKy7KIJnXUwB20EmOoxWn6BGeAskGtnotlLPo10uGbk"},
			want: &CustomClaims{
				Data: json.RawMessage{0x6e, 0x75, 0x6c, 0x6c},
				StandardClaims: jwtgo.StandardClaims{
					Audience:  "",
					ExpiresAt: 0,
					Id:        "",
					IssuedAt:  0,
					Issuer:    "",
					NotBefore: 0,
					Subject:   "",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &jwt{
				SigningKey: tt.fields.SigningKey,
			}
			got, err := j.ParseToken(tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Logf("输出got: %v", got)
				t.Errorf("ParseToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJWT_RefreshToken(t *testing.T) {
	type fields struct {
		SigningKey []byte
	}
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"case 01",
			fields{SigningKey: []byte("gwt_sign_key")},
			args{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJuYW1lIjoi5byg5LiJIiwicGFzc3dvcmQiOiIxMjM0NTYiLCJydWxlcyI6bnVsbH0.agleKMaE-ncgJetG8jGU4eLMlNsCBZN4CyN2pOSht4o"},
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJEYXRhIjpudWxsfQ.AKy7KIJnXUwB20EmOoxWn6BGeAskGtnotlLPo10uGbk",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &jwt{
				SigningKey: tt.fields.SigningKey,
			}
			got, err := j.RefreshToken(tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("RefreshToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RefreshToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewJWT(t *testing.T) {
	tests := []struct {
		name string
		want *JWT
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJWT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSigningKey(t *testing.T) {
	secretKey := "Z3d0X3NpZ25fa2V5"
	j := New()

	assert.Equal(t, secretKey, j.GetSigningKey())
}

func TestCreateTokenAndParse(t *testing.T) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := CustomClaims{
		Data: []byte(`{"UserId":"test-user-id-0001","UserName":"Silly Hat","RoleId":"test-role-id-0001"}`),
		StandardClaims: jwtgo.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	secretKey := "Z3d0X3NpZ25fa2V5"
	j := New()

	assert.Equal(t, secretKey, j.GetSigningKey())

	token, err := j.CreateToken(claims)
	assert.Nil(t, err)
	assert.NotNil(t, token)
	t.Logf("输出Token: %v\n", token)


	c, err := j.ParseToken(token)
	t.Logf("输出自定义Claims: %#v\n", c)

	assert.Equal(t, claims.Data, c.Data)
	assert.Equal(t, claims.Issuer, c.Issuer)

	var dst interface{}
	jerr := json.Unmarshal(c.Data, &dst)
	if jerr != nil {
		t.Errorf("error: %v", jerr)
	}
	t.Logf("输出Data: %s\n", dst)

	refreshToken, err := j.RefreshToken(token)
	assert.Nil(t, err)
	assert.NotNil(t, refreshToken)
	assert.Equal(t, refreshToken, token)
}

func TestTokenExpires(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJEYXRhIjp7IlVzZXJJZCI6InRlc3QtdXNlci1pZC0wMDAxIiwiVXNlck5hbWUiOiJTaWxseSBIYXQiLCJSb2xlSWQiOiJ0ZXN0LXJvbGUtaWQtMDAwMSJ9LCJleHAiOjE2MTI2ODA5NTR9.JYmPu4Bec-6s8JM4ViA1-M5frtH4xtQVbXGUbPhR8yo"

	j := New()
	valid, err := j.TokenExpires(token)
	assert.Nil(t, err)
	assert.False(t, valid)
}