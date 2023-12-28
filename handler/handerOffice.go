package handler

import (
	"ResearchManage/internal/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var office struct {
	Lab_id     int32   `json:"LabId" binding:"required"`
	Area       float64 `json:"Area" binding:"required"`
	Address    string  `json:"Address" binding:"required"`
	Manager_id int32   `json:"ManagerId" binding:"required"`
}

// CreateOffice handles the creation of a new office
func (apiCfg *apiConfig) CreateOffice(c *gin.Context) {

	// 解析参数
	if err := c.ShouldBindJSON(&office); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	officeID, err := apiCfg.DB.CreateOffice(c.Request.Context(), database.CreateOfficeParams{
		LabID:     office.Lab_id,
		Area:      office.Area,
		Address:   office.Address,
		Managerid: office.Manager_id,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"officeID": officeID})
}

// UpdateOffice updates an existing office
func (apiCfg *apiConfig) UpdateOffice(c *gin.Context) {
	var office database.Office
	if err := c.ShouldBindJSON(&office); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedOffice, err := apiCfg.DB.UpdateOffice(c.Request.Context(), database.UpdateOfficeParams{
		Officeid:  office.Officeid,
		LabID:     office.LabID,
		Area:      office.Area,
		Address:   office.Address,
		Managerid: office.Managerid,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"office": updatedOffice,
	})
}

// DeleteOffice deletes an existing office
func (apiCfg *apiConfig) DeleteOffice(c *gin.Context) {
	var officeIDparam int
	var officeID int32

	officeIDparam, err := strconv.Atoi(c.Param("officeID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Param officeID must be an integer"})
		return
	}

	officeID = int32(officeIDparam)

	err = apiCfg.DB.DeleteOffice(c, officeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Office deleted successfully"})
}

// ListOfficeAll lists all offices
func (apiCfg *apiConfig) ListOfficeAll(c *gin.Context) {
	offices, err := apiCfg.DB.ListOfficeAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, offices)
}
