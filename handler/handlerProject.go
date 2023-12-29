package handler

import (
	"ResearchManage/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// ListProjectAll handles listing all projects
func (apiCfg *apiConfig) ListProjectAll(c *gin.Context) {
	var (
		err error
	)

	// 查询数据库
	projectList, err := apiCfg.DB.ListProjectAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": projectList,
	})
}

// CreateProject handles the creation of a new project
func (apiCfg *apiConfig) CreateProject(c *gin.Context) {
	var (
		LeaderId int32
		params   struct {
			Name            string  `json:"projectName" binding:"required"`
			ResearchContent string  `json:"researchContent" binding:"required"`
			Fund            float64 `json:"fund" binding:"required"`
			StartDate       string  `json:"startDate" binding:"required"`
			EndDate         string  `json:"endDate" binding:"required"`
			QualityMonitor  int32   `json:"qualityMonitor" binding:"required"`
			ClientId        int32   `json:"clientId" binding:"required"`
			Leader          struct {
				Name        string `json:"name" binding:"required"`
				MobilePhone string `json:"mobilePhone" binding:"required"`
				Email       string `json:"email" binding:"required"`
			}
		}
	)

	// 解析参数
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 新建 leader
	// 查询是否已经存在
	isLeaderExist, err := apiCfg.DB.IsLeaderExists(c.Request.Context(), database.IsLeaderExistsParams{
		Name:         params.Leader.Name,
		Mobilephone:  params.Leader.MobilePhone,
		Emailaddress: params.Leader.Email,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 如果存在，获取 id, 否则新建
	if isLeaderExist {
		leaderId, err := apiCfg.DB.GetLeaderIdByInfo(c.Request.Context(), database.GetLeaderIdByInfoParams{
			Name:         params.Leader.Name,
			Mobilephone:  params.Leader.MobilePhone,
			Emailaddress: params.Leader.Email,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		LeaderId = leaderId
	} else {
		leaderId, err := apiCfg.DB.CreateLeader(c.Request.Context(), database.CreateLeaderParams{
			Name:         params.Leader.Name,
			Mobilephone:  params.Leader.MobilePhone,
			Emailaddress: params.Leader.Email,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		LeaderId = leaderId
	}

	projectID, err := apiCfg.DB.CreateProject(c.Request.Context(), database.CreateProjectParams{
		Projectleader:     LeaderId,
		Name:              params.Name,
		Researchcontent:   params.ResearchContent,
		Totalfunds:        params.Fund,
		Startdate:         params.StartDate,
		Enddate:           params.EndDate,
		Qualitymonitorsid: params.QualityMonitor,
		Clientid:          params.ClientId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": projectID,
	})
}

// UpdateProject updates an existing project
func (apiCfg *apiConfig) UpdateProject(c *gin.Context) {
	var project struct {
		Projectid         int32   `json:"projectId" binding:"required"`
		LeaderId          int32   `json:"leader" binding:"required"`
		Name              string  `json:"projectName" binding:"required"`
		Researchcontent   string  `json:"researchContent" binding:"required"`
		Fund              float64 `json:"fund" binding:"required"`
		Startdate         string  `json:"startDate" binding:"required"`
		Enddate           string  `json:"endDate" binding:"required"`
		Qualitymonitorsid int32   `json:"qualityMonitor" binding:"required"`
		Clientid          int32   `json:"clientId" binding:"required"`
	}
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedProject, err := apiCfg.DB.UpdateProject(c.Request.Context(), database.UpdateProjectParams{
		Projectid:         project.Projectid,
		Projectleader:     project.LeaderId,
		Name:              project.Name,
		Researchcontent:   project.Researchcontent,
		Totalfunds:        project.Fund,
		Startdate:         project.Startdate,
		Enddate:           project.Enddate,
		Qualitymonitorsid: project.Qualitymonitorsid,
		Clientid:          project.Clientid,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": updatedProject,
	})
}

// DeleteProject deletes an existing project
func (apiCfg *apiConfig) DeleteProject(c *gin.Context) {
	var projectIDparam struct {
		ProjectId int `json:"projectId" binding:"required"`
	}
	if err := c.ShouldBindJSON(&projectIDparam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	projectInfo := apiCfg.DB.DeleteProject(c.Request.Context(), int32(projectIDparam.ProjectId))

	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": projectInfo,
	})
}

func (apiCfg *apiConfig) ListProjectPartner(c *gin.Context) {
	var projectId int64
	projectId, err := strconv.ParseInt(c.Query("projectId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	projectPartnerList, err := apiCfg.DB.GetParterByProject(c.Request.Context(), int32(projectId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": projectPartnerList,
	})
}
