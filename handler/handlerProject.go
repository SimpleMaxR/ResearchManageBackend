package handler

import (
	"ResearchManage/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
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
		"projectList": projectList,
	})
}

// CreateProject handles the creation of a new project
func (apiCfg *apiConfig) CreateProject(c *gin.Context) {
	var project struct {
		Name            string  `json:"ProjectName" binding:"required"`
		Leader          int32   `json:"LeaderId" binding:"required"`
		ResearchContent string  `json:"ResearchContent" binding:"required"`
		Fund            float64 `json:"Fund" binding:"required"`
		StartDate       string  `json:"StartDate" binding:"required"`
		EndDate         string  `json:"EndDate" binding:"required"`
		QualityMonitor  int32   `json:"QualityMonitor" binding:"required"`
		ClientId        int32   `json:"ClientId" binding:"required"`
	}

	// 解析参数
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	projectID, err := apiCfg.DB.CreateProject(c.Request.Context(), database.CreateProjectParams{
		Peojectleader:     project.Leader,
		Name:              project.Name,
		Researchcontent:   project.ResearchContent,
		Totalfunds:        project.Fund,
		Startdate:         project.StartDate,
		Enddate:           project.EndDate,
		Qualitymonitorsid: project.QualityMonitor,
		Clientid:          project.ClientId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"projectID": projectID})
}

// UpdateProject updates an existing project
func (apiCfg *apiConfig) UpdateProject(c *gin.Context) {
	var project database.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedProject, err := apiCfg.DB.UpdateProject(c.Request.Context(), database.UpdateProjectParams{
		Projectid:         project.Projectid,
		Peojectleader:     project.Peojectleader,
		Name:              project.Name,
		Researchcontent:   project.Researchcontent,
		Totalfunds:        project.Totalfunds,
		Startdate:         project.Startdate,
		Enddate:           project.Enddate,
		Qualitymonitorsid: project.Qualitymonitorsid,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"project": updatedProject,
	})
}

// DeleteProject deletes an existing project
func (apiCfg *apiConfig) DeleteProject(c *gin.Context) {
	var projectIDparam int
	if err := c.ShouldBindJSON(&projectIDparam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	projectInfo := apiCfg.DB.DeleteProject(c.Request.Context(), int32(projectIDparam))

	c.JSON(http.StatusOK, gin.H{
		"projectInfo": projectInfo,
	})
}
