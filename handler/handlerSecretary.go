package handler

import (
	"ResearchManage/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	err error
)

var secretary struct {
	Name   string `json:"name" binding:"required"`
	Gender string `json:"gender" binding:"required"`
	Age    int32  `json:"age" binding:"required"`
	Mobile string `json:"mobile" binding:"required"`
	Email  string `json:"email" binding:"required"`
}

func (apiCfg apiConfig) CreateSecretary(c *gin.Context) {

	// 解析参数
	if err := c.ShouldBindJSON(&secretary); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter error " + err.Error()})
		return
	}

	// 查询数据库
	secretary, err := apiCfg.DB.CreateSecretary(c.Request.Context(), database.CreateSecretaryParams{
		Name:         secretary.Name,
		Gender:       secretary.Gender,
		Age:          secretary.Age,
		Mobilephone:  secretary.Mobile,
		Emailaddress: secretary.Email,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"code":      "200",
		"msg":       "success",
		"secretary": secretary,
	})
}

func (apiCfg apiConfig) DeleteSecretary(c *gin.Context) {

	var secretaryid int32
	// 解析参数
	if err := c.ShouldBindJSON(&secretaryid); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter error " + err.Error()})
		return
	}

	// 查询数据库
	err := apiCfg.DB.DeleteSecretary(c.Request.Context(), secretaryid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// TODO 新增数据库 trigger

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "success",
	})
}

func (apiCfg apiConfig) SetSecretary(c *gin.Context) {

	var (
		err    error
		id     int32
		params struct {
			SecretaryId    int32  `json:"SecretaryId" binding:"required"`
			LabId          int32  `json:"LabId" binding:"required"`
			EmployDate     string `json:"EmployDate" binding:"required"`
			responsibility string `json:"responsibility" binding:"required"`
		}
	)

	// 解析参数
	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter error " + err.Error()})
		return
	}

	// 查询数据库
	secretary, err := apiCfg.DB.CreateSecretaryService(c.Request.Context(), database.CreateSecretaryServiceParams{
		Secretaryid:      params.SecretaryId,
		LabID:            params.LabId,
		Employmentdate:   params.EmployDate,
		Responsibilities: params.responsibility,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"code":      "200",
		"msg":       "success",
		"secretary": secretary,
	})
}
