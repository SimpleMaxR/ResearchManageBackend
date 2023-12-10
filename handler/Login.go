package handler

import (
	"ResearchManage/internal/database"
	"ResearchManage/middleware"
	"ResearchManage/utils/errmsg"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (apiCfg apiConfig) HandlerAdminLogin(c *gin.Context) {
	var (
		err  error
		user database.User
	)

	// 解析参数
	var adminInfo struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err = c.ShouldBindJSON(&adminInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 查询数据库, 得到用户信息
	user, err = apiCfg.DB.AdminLogin(c.Request.Context(), database.AdminLoginParams{
		Username: adminInfo.Username,
		Password: adminInfo.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	setToken(c, user)
}

func setToken(c *gin.Context, user database.User) {
	j := middleware.NewJWT()
	claims := middleware.Claims{
		Userid: user.Userid,
		Role:   user.Roleid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 7200,
			IssuedAt:  time.Now().Unix(),
			Issuer:    "Simple",
		},
	}

	// 生成token
	token, err := j.CreatToken(claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  errmsg.SUCCESS,
		"message": errmsg.GetErrMsg(errmsg.SUCCESS),
		"token":   token,
		"userid":  user.Userid,
		"role":    user.Roleid,
	})
}
