package handler

import (
	"ResearchManage/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var qm struct {
	Name         string `json:"name" binding:"required"`
	Address      string `json:"address" binding:"required"`
	Leader       int32  `json:"leaderId" binding:"required"`
	ContactName  string `json:"contactName" binding:"required"`
	ContactPhone string `json:"contactPhone" binding:"required"`
}

var partners struct {
	Name         string `json:"name" binding:"required"`
	Address      string `json:"address" binding:"required"`
	Leader       int32  `json:"leaderId" binding:"required"`
	OfficePhone  string `json:"officePhone" binding:"required"`
	ContactName  string `json:"contactName" binding:"required"`
	ContactPhone string `json:"contactPhone" binding:"required"`
}

var leader struct {
	Name        string `json:"name" binding:"required"`
	OfficePhone string `json:"officePhone" binding:"required"`
	MobilePhone string `json:"mobilePhone" binding:"required"`
	Email       string `json:"email" binding:"required"`
}

var client struct {
	Name         string `json:"name" binding:"required"`
	Address      string `json:"address" binding:"required"`
	Leader       int32  `json:"leaderId" binding:"required"`
	Phone        string `json:"phone" binding:"required"`
	ContactName  string `json:"contactName" binding:"required"`
	ContactPhone string `json:"contactPhone" binding:"required"`
}

// Leader 相关接口

// CreateLeader handles the creation of a new leader
func (apiCfg *apiConfig) CreateLeader(c *gin.Context) {
	var (
		err error
	)

	// 解析参数
	if err := c.ShouldBindJSON(&leader); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 查询数据库
	leader, err := apiCfg.DB.CreateLeader(c.Request.Context(), database.CreateLeaderParams{
		Name:         leader.Name,
		Mobilephone:  leader.MobilePhone,
		Emailaddress: leader.Email,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"code":   "200",
		"msg":    "success",
		"leader": leader,
	})
}

// GetLeader get leader by id
func (apiCfg *apiConfig) GetLeader(c *gin.Context) {
	var (
		err error
		id  int32
	)

	// 解析参数
	err = c.ShouldBindJSON(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	// 查询数据库
	leader, err := apiCfg.DB.GetLeader(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"code":   "200",
		"msg":    "success",
		"leader": leader,
	})
}

// Quality Monitor 相关接口

// CreateQM handles the creation of a new qm
func (apiCfg *apiConfig) CreateQM(c *gin.Context) {
	var (
		err      error
		LeaderId int32
		params   struct {
			Name         string `json:"name" binding:"required"`
			Address      string `json:"address" binding:"required"`
			ContactName  string `json:"contactName" binding:"required"`
			ContactPhone string `json:"contactPhone" binding:"required"`
			Leader       struct {
				Name        string `json:"name" binding:"required"`
				MobilePhone string `json:"mobilePhone" binding:"required"`
				Email       string `json:"email" binding:"required"`
			}
			ProjectId int32 `json:"projectId" binding:"required"`
		}
	)

	// 解析参数
	if err = c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "params error" + err.Error()})
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

	// 新建 QM
	QMId, err := apiCfg.DB.CreateQM(c.Request.Context(), database.CreateQMParams{
		Name:         params.Name,
		Address:      params.Address,
		Leaderid:     LeaderId,
		Contactname:  params.ContactName,
		Contactphone: params.ContactPhone,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": QMId,
	})
}

// ListQM handles listing all qms
func (apiCfg *apiConfig) ListQMAll(c *gin.Context) {
	var (
		err error
	)

	// 查询数据库
	qmList, err := apiCfg.DB.ListQM(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": qmList,
	})
}

// GetQMByProjectID list qm by project id
func (apiCfg *apiConfig) GetQMByProjectID(c *gin.Context) {
	var (
		err       error
		projectId int64
	)

	// 解析参数
	projectId, err = strconv.ParseInt(c.Query("projectId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 查询数据库
	project, err := apiCfg.DB.GetProjectById(c.Request.Context(), int32(projectId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	qm, err := apiCfg.DB.GetQMById(c.Request.Context(), project.Qualitymonitorsid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": qm,
	})

}

// GetQMByID
func (apiCfg *apiConfig) GetQMByID(c *gin.Context) {
	var (
		err error
		id  int64
	)

	// 解析参数
	id, err = strconv.ParseInt(c.Query("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	// 查询数据库
	qm, err := apiCfg.DB.GetQMById(c.Request.Context(), int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": qm,
	})
}

// Partners 相关接口

// CreatePartner handles the creation of a new partners
func (apiCfg *apiConfig) CreatePartner(c *gin.Context) {
	var (
		err      error
		LeaderId int32
		params   struct {
			Name         string `json:"name" binding:"required"`
			Address      string `json:"address" binding:"required"`
			OfficePhone  string `json:"officePhone" binding:"required"`
			ContactPhone string `json:"contactPhone"`
			ContactName  string `json:"contactName"`
			Leader       struct {
				Name        string `json:"name" binding:"required"`
				MobilePhone string `json:"mobilePhone" binding:"required"`
				Email       string `json:"email" binding:"required"`
			}
			ProjectId int32 `json:"projectId" binding:"required"`
		}
	)

	// 解析参数
	if err = c.ShouldBindJSON(&params); err != nil {
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

	// 新建 partner
	partnerId, err := apiCfg.DB.CreatePartner(c.Request.Context(), database.CreatePartnerParams{
		Name:         params.Name,
		Address:      params.Address,
		Leaderid:     LeaderId,
		Officephone:  params.OfficePhone,
		Contactphone: params.ContactPhone,
		Contactname:  params.ContactName,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "success",
		"data": partnerId,
	})
}

// ListPartners handles listing all partners
func (apiCfg *apiConfig) ListPartners(c *gin.Context) {
	var (
		err error
	)

	// 查询数据库
	partnerList, err := apiCfg.DB.ListPartner(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Database error" + err.Error(),
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "success",
		"data": partnerList,
	})
}

// GetPartnerByProjectID list partner by project id
func (apiCfg *apiConfig) GetPartnerByProjectID(c *gin.Context) {
	var (
		err       error
		projectID int64
	)

	// 解析参数
	projectID, err = strconv.ParseInt(c.Query("projectId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "projectID is required"})
		return
	}

	// 查询数据库
	partners, err := apiCfg.DB.GetParterByProject(c.Request.Context(), int32(projectID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "success",
		"data": partners,
	})
}

// LinkProjectPartner links a project with a partner
func (apiCfg *apiConfig) LinkProjectPartner(c *gin.Context) {
	var projectPartner struct {
		ProjectID int32 `json:"projectId" binding:"required"`
		PartnerID int32 `json:"partnerId" binding:"required"`
	}

	if err := c.ShouldBindJSON(&projectPartner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	projectPartnerInfo := apiCfg.DB.LinkProjectPartner(c.Request.Context(), database.LinkProjectPartnerParams{
		Projectid: projectPartner.ProjectID,
		Partnerid: projectPartner.PartnerID,
	})

	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": projectPartnerInfo,
	})
}

// Client 相关接口

// CreateClient
func (apiCfg *apiConfig) CreateClient(c *gin.Context) {
	var (
		LeaderId int32
		err      error
		params   struct {
			Name         string `json:"name" binding:"required"`
			Address      string `json:"address" binding:"required"`
			OfficePhone  string `json:"officePhone" binding:"required"`
			ContactName  string `json:"contactName" binding:"required"`
			ContactPhone string `json:"contactPhone" binding:"required"`
			Leader       struct {
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

	// 查询数据库
	client, err := apiCfg.DB.CreateClient(c.Request.Context(), database.CreateClientParams{
		Name:         params.Name,
		Address:      params.Address,
		Leaderid:     LeaderId,
		Officephone:  params.OfficePhone,
		Contactname:  params.ContactName,
		Contactphone: params.ContactPhone,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"code":   "200",
		"msg":    "success",
		"client": client,
	})
}

// UpdateClient
func (apiCfg *apiConfig) UpdateClient(c *gin.Context) {
	var (
		err    error
		params struct {
			Clientid     int32  `json:"clientId" binding:"required"`
			Name         string `json:"name" binding:"required"`
			Address      string `json:"address" binding:"required"`
			Leader       int32  `json:"leaderId" binding:"required"`
			Phone        string `json:"phone" binding:"required"`
			ContactName  string `json:"contactName"`
			ContactPhone string `json:"contactPhone"`
		}
	)

	// 解析参数
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 查询数据库
	client, err := apiCfg.DB.UpdateClient(c.Request.Context(), database.UpdateClientParams{
		Clientid:     params.Clientid,
		Name:         params.Name,
		Address:      params.Address,
		Leaderid:     params.Leader,
		Officephone:  params.Phone,
		Contactname:  params.ContactName,
		Contactphone: params.ContactPhone,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"code":   "200",
		"msg":    "success",
		"client": client,
	})
}

// GetClientByProjectID list client by project id
func (apiCfg *apiConfig) GetClientByProjectID(c *gin.Context) {
	var (
		err       error
		projectID int32
	)

	// 解析参数
	err = c.ShouldBindJSON(projectID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "projectID is required"})
		return
	}

	// 查询数据库
	project, err := apiCfg.DB.GetProjectById(c.Request.Context(), projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	client, err := apiCfg.DB.GetClient(c.Request.Context(), project.Clientid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"code":   "200",
		"msg":    "success",
		"client": client,
	})
}
