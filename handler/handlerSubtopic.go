package handler

import (
	"ResearchManage/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// ListSubtopicByProject listing all subtopics by project
func (apiCfg *apiConfig) ListSubtopicByProject(c *gin.Context) {
	var (
		err       error
		projectId int64
	)

	projectId, err = strconv.ParseInt(c.Query("projectId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "projectId is required"})
		return
	}

	// 查询数据库
	subtopics, err := apiCfg.DB.ListSubtopicByProject(c, int32(projectId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": subtopics})
}

// ListSubtopicByLeader listing all subtopics by leader
func (apiCfg *apiConfig) ListSubtopicByLeader(c *gin.Context) {
	var (
		err      error
		leaderId int64
	)

	leaderId, err = strconv.ParseInt(c.Query("leaderId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "leaderId is required"})
		return
	}

	// 查询数据库
	subtopics, err := apiCfg.DB.ListSubtopicByLeader(c, int32(leaderId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": subtopics})
}

// CreateSubtopic creates a new subtopic
func (apiCfg *apiConfig) CreateSubtopic(c *gin.Context) {
	var (
		err      error
		subtopic struct {
			ProjectId int64   `json:"projectId" binding:"required"`
			LeaderId  int64   `json:"leaderId" binding:"required"`
			Name      string  `json:"name" binding:"required"`
			Enddate   string  `json:"enddate" binding:"required"`
			Fund      float64 `json:"fund" binding:"required"`
			Tech      string  `json:"tech" binding:"required"`
		}
	)

	if err = c.ShouldBindJSON(&subtopic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 查询数据库
	subtopicInfo, err := apiCfg.DB.CreateSubtopic(c, database.CreateSubtopicParams{
		Projectid: int32(subtopic.ProjectId),
		Leaderid:  int32(subtopic.LeaderId),
		Name:      subtopic.Name,
		Enddate:   subtopic.Enddate,
		Fund:      subtopic.Fund,
		Tech:      subtopic.Tech,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": subtopicInfo,
	})
}

// UpdateSubtopic updates an existing subtopic
func (apiCfg *apiConfig) UpdateSubtopic(c *gin.Context) {
	var (
		err      error
		subtopic struct {
			SubtopicId int64   `json:"subtopicId" binding:"required"`
			ProjectId  int64   `json:"projectId" binding:"required"`
			LeaderId   int64   `json:"leaderId" binding:"required"`
			Name       string  `json:"name" binding:"required"`
			Enddate    string  `json:"enddate" binding:"required"`
			Fund       float64 `json:"fund" binding:"required"`
			Tech       string  `json:"tech" binding:"required"`
		}
	)

	if err = c.ShouldBindJSON(&subtopic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 查询数据库
	subtopicInfo, err := apiCfg.DB.UpdateSubtopic(c, database.UpdateSubtopicParams{
		Subtopicid: int32(subtopic.SubtopicId),
		Projectid:  int32(subtopic.ProjectId),
		Leaderid:   int32(subtopic.LeaderId),
		Name:       subtopic.Name,
		Enddate:    subtopic.Enddate,
		Fund:       subtopic.Fund,
		Tech:       subtopic.Tech,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": subtopicInfo,
	})
}

// DeleteSubtopic deletes an existing subtopic
func (apiCfg *apiConfig) DeleteSubtopic(c *gin.Context) {
	var (
		err        error
		subtopicId int64
	)

	subtopicId, err = strconv.ParseInt(c.Query("subtopicId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "subtopicId is required"})
		return
	}

	// 查询数据库
	subtopicInfo := apiCfg.DB.DeleteSubtopic(c, int32(subtopicId))

	c.JSON(http.StatusOK, gin.H{
		"data": subtopicInfo,
	})
}
