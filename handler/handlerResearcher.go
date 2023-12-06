package handler

import (
	"ResearchManage/internal/database"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (apiCfg *apiConfig) HandlerListResearcherAll(c *gin.Context) {
	var (
		err error
	)

	// 查询数据库
	researcherList, err := apiCfg.DB.ListResearcherAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"researcherList": researcherList,
	})
}

func (apiCfg *apiConfig) HandlerCreateResearcher(c *gin.Context) {
	var (
		err error
	)

	// 解析参数
	var researcherInfo struct {
		LabID             int32  `json:"LabId" binding:"required"`
		Name              string `json:"Name" binding:"required"`
		Gender            string `json:"Gender" binding:"required"`
		Title             string `json:"Title" binding:"required"`
		Age               int32  `json:"Age" binding:"required"`
		Researchdirection string `json:"Researchdirection"`
		Leader            bool   `json:"Leader" binding:"required"`
	}

	if err = c.ShouldBindJSON(&researcherInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 插入数据库
	id, err := apiCfg.DB.CreateResearcher(c.Request.Context(), database.CreateResearcherParams{
		Labid: sql.NullInt32{
			Int32: int32(researcherInfo.LabID),
			Valid: true,
		},
		Name:   researcherInfo.Name,
		Gender: researcherInfo.Gender,
		Title:  researcherInfo.Title,
		Age:    researcherInfo.Age,
		Researchdirection: sql.NullString{
			String: researcherInfo.Researchdirection,
			Valid:  true,
		},
		Leader: researcherInfo.Leader})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"researcherID": id,
	})
}

func (apiCfg *apiConfig) HandlerDeleteResearcher(c *gin.Context) {
	var (
		err error
	)

	// 获取参数
	var researcher struct {
		ResearcherID int `json:"ResearcherID"`
	}
	if err = c.ShouldBindJSON(&researcher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 删除数据库
	researcherInfo, err := apiCfg.DB.DeleteResearcher(c.Request.Context(), int32(researcher.ResearcherID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"researcherInfo": researcherInfo,
	})
}

func (apiCfg *apiConfig) HandlerUpdateResearcher(c *gin.Context) {
	var (
		err error
	)

	// 解析参数
	var researcherInfo struct {
		LabID             int32  `json:"LabId" binding:"required"`
		Name              string `json:"Name" binding:"required"`
		Gender            string `json:"Gender" binding:"required"`
		Title             string `json:"Title" binding:"required"`
		Age               int32  `json:"Age" binding:"required"`
		Researchdirection string `json:"Researchdirection"`
		Leader            bool   `json:"Leader" binding:"required"`
	}

	if err = c.ShouldBindJSON(&researcherInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 更新数据库
	researcher, err := apiCfg.DB.UpdateResearcher(c.Request.Context(), database.UpdateResearcherParams{
		Labid: sql.NullInt32{
			Int32: int32(researcherInfo.LabID),
			Valid: true,
		},
		Name:   researcherInfo.Name,
		Gender: researcherInfo.Gender,
		Title:  researcherInfo.Title,
		Age:    researcherInfo.Age,
		Researchdirection: sql.NullString{
			String: researcherInfo.Researchdirection,
			Valid:  true,
		},
		Leader: researcherInfo.Leader})

	// 返回数据
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"researcher_updated": researcher,
		})
		return
	}
}

func (apiCfg *apiConfig) HandlerListResearcherByLab(c *gin.Context) {
	var (
		err error
	)

	// 获取参数
	var lab struct {
		LabID int `json:"LabID"`
	}
	if err = c.ShouldBindJSON(&lab); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 查询数据库
	researcherList, err := apiCfg.DB.ListResearcherByLab(c.Request.Context(), sql.NullInt32{
		Int32: int32(lab.LabID),
		Valid: true,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"researcherList": researcherList,
	})
}
