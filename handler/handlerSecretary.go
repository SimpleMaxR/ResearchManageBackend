package handler

import (
	"ResearchManage/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
		params struct {
			SecretaryId    int32  `json:"SecretaryId" binding:"required"`
			LabId          int32  `json:"LabId" binding:"required"`
			EmployDate     string `json:"EmployDate" binding:"required"`
			Responsibility string `json:"responsibility" binding:"required"`
		}
	)

	// 解析参数
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter error " + err.Error()})
		return
	}

	// 查询数据库
	secretary, err := apiCfg.DB.CreateSecretaryService(c.Request.Context(), database.CreateSecretaryServiceParams{
		Secretaryid:      params.SecretaryId,
		LabID:            params.LabId,
		Employmentdate:   params.EmployDate,
		Responsibilities: params.Responsibility,
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

func (apiCfg apiConfig) ListSecretaryByLab(c *gin.Context) {

	var labid int64
	var secretaryForLab struct {
		Secretaryservice []database.Secretaryservice
		Secretary        []database.Secretary
	}
	// 解析参数
	labid, err = strconv.ParseInt(c.Query("labID"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter error " + err.Error()})
		return
	}

	// 查询数据库
	// 查询秘书服务
	secretaryServices, err := apiCfg.DB.ListSecretaryServiceByLab(c.Request.Context(), int32(labid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 查询秘书
	for _, secretaryService := range secretaryServices {
		secretary, err := apiCfg.DB.ListSecretaryByID(c.Request.Context(), secretaryService.Secretaryid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		secretaryForLab.Secretaryservice = append(secretaryForLab.Secretaryservice, secretaryService)
		secretaryForLab.Secretary = append(secretaryForLab.Secretary, secretary)
	}
	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "success",
		"data": secretaryForLab.Secretary,
	})
}

func (apiCfg apiConfig) ListSecretaryServiceByLab(c *gin.Context) {

	var labid int64
	// 解析参数
	labid, err = strconv.ParseInt(c.Query("labID"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter error " + err.Error()})
		return
	}

	// 查询数据库
	// 查询 lab 的秘书
	secretary, err := apiCfg.DB.ListSecretaryServiceByLab(c.Request.Context(), int32(labid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "success",
		"data": secretary,
	})
}

func (apiCfg apiConfig) ListSecretaryAll(c *gin.Context) {

	// 查询数据库
	// 查询 lab 的秘书
	secretary, err := apiCfg.DB.ListSecretaryAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": secretary,
	})
}
