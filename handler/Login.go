package handler

import (
	"ResearchManage/internal/database"
	"ResearchManage/middleware"
	"ResearchManage/utils/errmsg"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (apiCfg apiConfig) HandlerAdminLogin(c *gin.Context) {
	var (
		err error
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

	// 查询数据库
	data, err = apiCfg.DB.AdminLogin(c.Request.Context(), database.AdminLoginParams{
		Username: adminInfo.Username,
		Password: adminInfo.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if data.Code == errmsg.SUCCESS {
		setToken(c, data)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func setToken(c *gin.Context, data interface{}) {
	jwt := middleware.NewJWT()
	claims := middleware.Claims{
		Userid: data.ID,
		Role:   data.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 7200,
			Issuer:    "Simple",
		},
	}

	token, err := jwt.CreatToken(claims)
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
		"userid":  data.ID,
		"role":    data.Role,
	})
}
