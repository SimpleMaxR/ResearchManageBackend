package handler

import (
	"ResearchManage/internal/database"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

var qm struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
	Leader  int32  `json:"leaderId" binding:"required"`
}

var partners struct {
	Name        string `json:"name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Leader      int32  `json:"leaderId" binding:"required"`
	OfficePhone string `json:"officePhone" binding:"required"`
}

var leader struct {
	Name        string `json:"name" binding:"required"`
	OfficePhone string `json:"officePhone" binding:"required"`
	MobilePhone string `json:"mobilePhone" binding:"required"`
	Email       string `json:"email" binding:"required"`
}

var client struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
	Leader  int32  `json:"leaderId" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
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
		Officephone:  leader.OfficePhone,
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
		err error
	)

	// 解析参数
	if err := c.ShouldBindJSON(&qm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 查询数据库
	qmList, err := apiCfg.DB.CreateQM(c.Request.Context(), database.CreateQMParams{
		Name:     qm.Name,
		Address:  qm.Address,
		Leaderid: qm.Leader,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"qmList": qmList,
	})
}

// ListQM handles listing all qms
func (apiCfg *apiConfig) ListQM(c *gin.Context) {
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
		"qmList": qmList,
	})
}

// GetQMByProjectID list qm by project id
func (apiCfg *apiConfig) GetQMByProjectID(c *gin.Context) {
	var (
		err       error
		projectID int32
	)

	// 解析参数
	err = c.ShouldBindJSON(projectID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	qm, err := apiCfg.DB.GetQMById(c.Request.Context(), project.Qualitymonitorsid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"qm": qm,
	})

}

// SetQMContact set qm contact
func (apiCfg *apiConfig) SetQMContact(c *gin.Context) {
	var (
		err                error
		setQMContactParams struct {
			QMID        int32  `json:"qmId" binding:"required"`
			Name        string `json:"name" binding:"required"`
			OfficePhone string `json:"officePhone" binding:"required"`
			MobilePhone string `json:"mobilePhone" binding:"required"`
			Email       string `json:"email" binding:"required"`
		}
	)

	// 解析参数
	if err := c.ShouldBindJSON(&setQMContactParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "params is wrong"})
		return
	}

	// 查询数据库
	contactId, err := apiCfg.DB.CreateContact(c.Request.Context(), database.CreateContactParams{
		Name:         qm.Name,
		Officephone:  setQMContactParams.OfficePhone,
		Mobilephone:  setQMContactParams.MobilePhone,
		Emailaddress: setQMContactParams.Email,
	})

	contactId, err = apiCfg.DB.SetContactQM(c.Request.Context(), database.SetContactQMParams{
		Baseqm: sql.NullInt32{
			Int32: setQMContactParams.QMID,
			Valid: true,
		},
		Contactid: contactId,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error " + err.Error()})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"contactId": contactId,
	})
}

// Partners 相关接口

// CreatePartners handles the creation of a new partners
func (apiCfg *apiConfig) CreatePartner(c *gin.Context) {
	var (
		err error
	)

	// 解析参数
	if err := c.ShouldBindJSON(&partners); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 查询数据库
	partner, err := apiCfg.DB.CreatePartner(c.Request.Context(), database.CreatePartnerParams{
		Name:        partners.Name,
		Address:     partners.Address,
		Leaderid:    partners.Leader,
		Officephone: partners.OfficePhone,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "success",
		"data": partner,
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
		projectID int32
	)

	// 解析参数
	err = c.ShouldBindJSON(projectID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "projectID is required"})
		return
	}

	// 查询数据库
	partners, err := apiCfg.DB.GetParterByProject(c.Request.Context(), projectID)
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
		ProjectID int32 `json:"ProjectID" binding:"required"`
		PartnerID int32 `json:"PartnerID" binding:"required"`
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
		"projectPartnerInfo": projectPartnerInfo,
	})
}

// Client 相关接口

// CreateClient
func (apiCfg *apiConfig) CreateClient(c *gin.Context) {
	var (
		err error
	)

	// 解析参数
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 查询数据库
	client, err := apiCfg.DB.CreateClient(c.Request.Context(), database.CreateClientParams{
		Name:        client.Name,
		Address:     client.Address,
		Leaderid:    client.Leader,
		Officephone: client.Phone,
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
			Clientid int32  `json:"clientId" binding:"required"`
			Name     string `json:"name" binding:"required"`
			Address  string `json:"address" binding:"required"`
			Leader   int32  `json:"leaderId" binding:"required"`
			Phone    string `json:"phone" binding:"required"`
		}
	)

	// 解析参数
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 查询数据库
	client, err := apiCfg.DB.UpdateClient(c.Request.Context(), database.UpdateClientParams{
		Clientid:    params.Clientid,
		Name:        params.Name,
		Address:     params.Address,
		Leaderid:    params.Leader,
		Officephone: params.Phone,
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

// SetClientContact
func (apiCfg *apiConfig) CreateClientContact(c *gin.Context) {
	var (
		err                    error
		setClientContactParams struct {
			ClientID    int32  `json:"clientId" binding:"required"`
			Name        string `json:"name" binding:"required"`
			OfficePhone string `json:"officePhone" binding:"required"`
			MobilePhone string `json:"mobilePhone" binding:"required"`
			Email       string `json:"email" binding:"required"`
		}
	)

	// 解析参数
	if err := c.ShouldBindJSON(&setClientContactParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "params is wrong"})
		return
	}

	// 查询数据库
	contactId, err := apiCfg.DB.CreateContact(c.Request.Context(), database.CreateContactParams{
		Name:         setClientContactParams.Name,
		Officephone:  setClientContactParams.OfficePhone,
		Mobilephone:  setClientContactParams.MobilePhone,
		Emailaddress: setClientContactParams.Email,
	})

	contactId, err = apiCfg.DB.SetContactClient(c.Request.Context(), database.SetContactClientParams{
		Baseclient: sql.NullInt32{
			Int32: setClientContactParams.ClientID,
			Valid: true,
		},
		Contactid: contactId,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error " + err.Error()})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"contactId": contactId,
	})
}
