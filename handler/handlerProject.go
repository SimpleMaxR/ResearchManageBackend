package handler

import (
	"ResearchManage/internal/database"
	"database/sql"
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

func (apiCfg *apiConfig) ListProjectByName(c *gin.Context) {
	var (
		err  error
		name string
	)

	// 获取参数
	name, exist := c.GetQuery("name")
	if exist == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 查询数据库
	project, err := apiCfg.DB.GetProjectByName(c.Request.Context(), sql.NullString{
		String: name,
		Valid:  exist,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"data": project,
	})
}

// CreateProject handles the creation of a new project
func (apiCfg *apiConfig) CreateProject(c *gin.Context) {
	var (
		params struct {
			Name            string  `json:"projectName" binding:"required"`
			ResearchContent string  `json:"researchContent" binding:"required"`
			Fund            float64 `json:"fund" binding:"required"`
			StartDate       string  `json:"startDate" binding:"required"`
			EndDate         string  `json:"endDate" binding:"required"`
			QualityMonitor  int32   `json:"qualityMonitor" binding:"required"`
			ClientId        int32   `json:"clientId" binding:"required"`
			Projectleader   int32   `json:"projectleader" binding:"required"`
		}
	)

	// 解析参数
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	projectID, err := apiCfg.DB.CreateProject(c.Request.Context(), database.CreateProjectParams{
		Projectleader:     params.Projectleader,
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

func (apiCfg *apiConfig) LinkProjectResearcher(c *gin.Context) {
	var params struct {
		ProjectId  int32  `json:"projectId" binding:"required"`
		Researcher int32  `json:"researcherid" binding:"required"`
		Join       string `json:"joindate" binding:"required"`
		Workload   string `json:"workload" binding:"required"`
	}

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := apiCfg.DB.LinkProjectResearcher(c.Request.Context(), database.LinkProjectResearcherParams{
		Projectid:    params.ProjectId,
		Researcherid: params.Researcher,
		Joindate:     params.Join,
		Workload:     params.Workload,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": nil,
	})
}

func (apiCfg *apiConfig) ListProjectResearcher(c *gin.Context) {
	var projectId int64
	projectId, err := strconv.ParseInt(c.Query("projectId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	projectResearcherList, err := apiCfg.DB.ListProjectResearcher(c.Request.Context(), int32(projectId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "ok",
		"data": projectResearcherList,
	})
}
