package tools

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

const (
	AccessTokenDuration  = 2 * time.Hour
	RefreshTokenDuration = 30 * 24 * time.Hour
	TokenIssuer          = "12345678"
)

var Token VoteJwt

func NewToken(s string) {
	b := []byte("xxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	if s != "" {
		b = []byte(s)
	}

	Token = VoteJwt{Secret: b}
}

type VoteJwt struct {
	Secret []byte
}

// Claim 自定义的数据结构，这里使用了结构体的组合
type Claim struct {
	jwt.RegisteredClaims
	ID       int64  `json:"userId"`
	UserName string `json:"userName"'`
	Name     string `json:"name"`
	Role     string `json:"role""`
}

func (j *VoteJwt) getTime(t time.Duration) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Now().Add(t))
}

func (j *VoteJwt) keyFunc(token *jwt.Token) (interface{}, error) {
	return j.Secret, nil
}

// GetToken 颁发token access token 和 refresh token
func (j *VoteJwt) GetToken(id int64, userName string, name string, role string) (aToken string, err error) {
	rc := jwt.RegisteredClaims{
		ExpiresAt: j.getTime(AccessTokenDuration),
		Issuer:    TokenIssuer,
	}
	claim := Claim{
		ID:               id,
		UserName:         userName,
		Name:             name,
		Role:             role,
		RegisteredClaims: rc,
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	bytes := []byte(TokenIssuer)
	return claims.SignedString(bytes)

}

// VerifyToken 验证Token
func (j *VoteJwt) VerifyToken(tokenID string) (*Claim, error) {
	claim := &Claim{}
	//token, err := jwt.ParseWithClaims(tokenID, claim, j.keyFunc)
	token, err := jwt.ParseWithClaims(tokenID, claim, func(token *jwt.Token) (interface{}, error) {
		bytes := []byte(TokenIssuer)
		return bytes, nil
	})

	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("access token 验证失败")
	}

	return claim, nil
}

// ImJwtAuthMiddleware 用户单独聊天室的权限验证服务
func (j *VoteJwt) ImJwtAuthMiddleware(tokenStr string) (userid, username string, err error) {
	claim := &Claim{}
	if tokenStr == "" {
		fmt.Println("Im验证权限失败")
		return "", "", err
	}
	token, err := jwt.ParseWithClaims(tokenStr, claim, func(token *jwt.Token) (interface{}, error) {
		bytes := []byte(TokenIssuer)
		return bytes, nil
	})
	if err != nil {
		fmt.Println("Im验证权限失败")
		return "", "", err
	}
	if !token.Valid {
		fmt.Println("Im验证权限失败")
		return "", "", err
	}
	userid = strconv.Itoa(int(claim.ID))
	return userid, claim.UserName, nil
}

// RefreshToken 通过 refresh token 刷新 access token
//func (j *VoteJwt) RefreshToken(a, r string) (aToken, rToken string, err error) {
//	// r 无效直接返回
//	if _, err = jwt.Parse(r, j.keyFunc); err != nil {
//		return
//	}
//	// 从旧access token 中解析出claims数据
//	claim := &Claim{}
//	_, err = jwt.ParseWithClaims(a, claim, j.keyFunc)
//	// 判断错误是不是因为access token 正常过期导致的
//	if errors.Is(err, jwt.ErrTokenExpired) {
//		return j.GetToken(claim.ID, claim.UserName, claim.Name, claim.Role)
//	}
//	return
//}
