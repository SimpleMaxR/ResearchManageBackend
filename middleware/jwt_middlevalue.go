package middleware

import (
	"ResearchManage/utils"
	"ResearchManage/utils/errmsg"
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JWT struct {
	JwtKey []byte
}

// 创建 JWT 实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(utils.JwtKey),
	}
}

// JwtKey 加密秘钥 (string)
var JwtKey = []byte(utils.JwtKey)

// Claims token 结构体
type Claims struct {
	Userid int32 `json:"userid"`
	Role   int32 `json:"roleid"`
	jwt.StandardClaims
}

// 定义错误
var (
	TokenExpired     error = errors.New("token 已过期,请重新登录")
	TokenNotValueYet error = errors.New("此 token 无效,请重新登录")
	TokenMalFormed   error = errors.New("token 不正确,请重新登录")
	TokenInvalid     error = errors.New("token 格式错误,请重新登录")
)

// CreatToken 生成TOKEN(J*JWT)
func (j *JWT) CreatToken(claims Claims) (string, error) {
	//var J *JWT
	//expireTime := time.Now().Add(7 * 24 * time.Hour)
	//claims := &Claims{
	//	UserID: user.ID,
	//	StandardClaims: jwt.StandardClaims{
	//		ExpiresAt: expireTime.Unix(),
	//		IssuedAt:  time.Now().Unix(),
	//		Issuer:    "Yamada",
	//		Subject:   "user token",
	//	},
	//}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(j.JwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParserToken 验证TOKEN
func (j *JWT) ParserToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return JwtKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, nil, TokenMalFormed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, nil, TokenNotValueYet
			} else {
				return nil, nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claimsData, ok := token.Claims.(*Claims); ok && token.Valid {
			return nil, claimsData, nil
		}
		return nil, nil, TokenInvalid
	}

	return nil, nil, TokenInvalid
}

// JwtToken 鉴权
func JwtToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		// token := context.Request.Header.Get("Authorization")
		context.Next()
	}
}

// JwtTokenTeacher 鉴定
func JwtTokenTeacher() gin.HandlerFunc {
	return func(context *gin.Context) {
		var code int
		//获取Authorization
		tokenString := context.GetHeader("Authorization")

		//判断token 的情况
		if tokenString == "" {
			code = errmsg.ERROR_TOKEN_EXIST
			context.JSON(http.StatusUnauthorized, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			context.Abort()
			return
		}
		checkToken := strings.Split(tokenString, " ")
		if len(checkToken) == 0 {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			context.Abort()
			return
		}

		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			context.Abort()
			return
		}

		j := NewJWT()
		//解析Token
		_, claims, err := j.ParserToken(checkToken[1])
		if err != nil {
			if err == TokenExpired {
				context.JSON(http.StatusUnauthorized, gin.H{
					"status":  errmsg.ERROR,
					"message": TokenExpired,
					"data":    nil,
				})
				context.Abort()
				return
			}
			context.JSON(http.StatusUnauthorized, gin.H{
				"status":  errmsg.ERROR,
				"message": err.Error(),
				"data":    nil,
			})
			context.Abort()
			return
		}
		if claims.Role != 1 {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status":  errmsg.ERROR_USER_NOT_RIGHT,
				"message": errmsg.GetErrMsg(errmsg.ERROR_USER_NOT_RIGHT),
				"data":    nil,
			})
			context.Abort()
			return
		}

		context.Set("id", claims.Id)
		context.Set("userid", claims.Userid)
		context.Set("role", claims.Role)
		context.Next()
	}
}

// JwtTokenStudent 中间键
func JwtTokenStudent() gin.HandlerFunc {
	return func(context *gin.Context) {
		var code int
		//获取Authorization
		tokenString := context.GetHeader("Authorization")

		//判断token 的情况
		if tokenString == "" {
			code = errmsg.ERROR_TOKEN_EXIST
			context.JSON(http.StatusUnauthorized, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			context.Abort()
			return
		}
		checkToken := strings.Split(tokenString, " ")
		if len(checkToken) == 0 {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			context.Abort()
			return
		}

		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			context.Abort()
			return
		}

		j := NewJWT()
		//解析Token
		_, claims, err := j.ParserToken(checkToken[1])
		if err != nil {
			if err == TokenExpired {
				context.JSON(http.StatusUnauthorized, gin.H{
					"status":  errmsg.ERROR,
					"message": TokenExpired,
					"data":    nil,
				})
				context.Abort()
				return
			}
			context.JSON(http.StatusUnauthorized, gin.H{
				"status":  errmsg.ERROR,
				"message": err.Error(),
				"data":    nil,
			})
			context.Abort()
			return
		}

		context.Set("id", claims.Id)
		context.Set("userid", claims.Userid)
		context.Set("role", claims.Role)
		context.Next()
	}
}

// JwtTokenLocalManager 中间键
func JwtTokenLocalManager() gin.HandlerFunc {
	return func(context *gin.Context) {
		var code int
		//获取Authorization
		tokenString := context.GetHeader("token")
		//判断token 的情况
		if tokenString == "" {
			code = errmsg.ERROR_TOKEN_NOT_EXIST
			context.JSON(http.StatusUnauthorized, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			context.Abort()
			return
		}
		checkToken := strings.Split(tokenString, " ")
		if len(checkToken) == 0 {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			context.Abort()
			return
		}

		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			context.Abort()
			return
		}

		j := NewJWT()
		//解析Token
		_, claims, err := j.ParserToken(checkToken[1])
		if err != nil {
			if err == TokenExpired {
				context.JSON(http.StatusUnauthorized, gin.H{
					"status":  errmsg.ERROR,
					"message": TokenExpired,
					"data":    nil,
				})
				context.Abort()
				return
			}
			context.JSON(http.StatusUnauthorized, gin.H{
				"status":  errmsg.ERROR,
				"message": err.Error(),
				"data":    nil,
			})
			context.Abort()
			return
		}
		if claims.Role != 0 {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status":  errmsg.ERROR_USER_NOT_RIGHT,
				"message": errmsg.GetErrMsg(errmsg.ERROR_USER_NOT_RIGHT),
				"data":    nil,
			})
			context.Abort()
			return
		}

		context.Set("userid", claims.Userid)
		context.Set("role", claims.Role)
		context.Next()
	}
}
