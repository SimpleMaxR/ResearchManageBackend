package handler

import (
	"ResearchManage/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	Achievement struct {
		Name      string `json:"name" binding:"required"`
		Obtain    string `json:"obtain" binding:"required"`
		Project   int32  `json:"baseproject" binding:"required"`
		Subtopoic int32  `json:"basesubtopic" binding:"required"`
		Type      int32  `json:"type" binding:"required"`
	}
)

func (apiCfg apiConfig) CreateAchievement(c *gin.Context) {

	// 解析参数
	if err := c.ShouldBindJSON(&Achievement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter error " + err.Error()})
		return
	}

	// 查询数据库
	achievement, err := apiCfg.DB.CreateAchievement(c.Request.Context(), database.CreateAchievementParams{
		Name:         Achievement.Name,
		Obtaine:      Achievement.Obtain,
		Baseproject:  Achievement.Project,
		Basesubtopic: Achievement.Subtopoic,
		Type:         Achievement.Type,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": achievement,
	})
}

func (apiCfg apiConfig) DeleteAchievement(c *gin.Context) {

	var achievementid int32
	// 解析参数
	if err := c.ShouldBindJSON(&achievementid); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter error " + err.Error()})
		return
	}

	// 查询数据库
	err := apiCfg.DB.DeleteAchievement(c.Request.Context(), achievementid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "success",
	})
}

func (apiCfg apiConfig) ListAchievementByProject(c *gin.Context) {

	var projectId int64

	// 解析参数
	projectId, err = strconv.ParseInt(c.Query("projectId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter error " + err.Error()})
		return
	}

	// 查询数据库
	achievements, err := apiCfg.DB.ListAchievementByProject(c.Request.Context(), int32(projectId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "success",
		"data": achievements,
	})
}

func (apiCfg apiConfig) ListAchievementBySubtopic(c *gin.Context) {

	var subtopicid int32
	// 解析参数
	if err := c.ShouldBindJSON(&subtopicid); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter error " + err.Error()})
		return
	}

	// 查询数据库
	achievements, err := apiCfg.DB.ListAchievementBySubtopic(c.Request.Context(), subtopicid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"code":         "200",
		"msg":          "success",
		"achievements": achievements,
	})
}
