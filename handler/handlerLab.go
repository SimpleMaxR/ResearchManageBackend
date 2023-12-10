package handler

import (
	"ResearchManage/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (apiCfg *apiConfig) HandlerListLabAll(c *gin.Context) {
	var (
		err error
	)

	// 查询数据库
	labList, err := apiCfg.DB.ListLabAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"labList": labList,
	})
}

func (apiCfg *apiConfig) HandlerListLabByLabID(c *gin.Context) {
	var (
		err error
	)

	// 获取参数
	var lab struct {
		LabID int32 `json:"LabID"`
	}
	if err = c.ShouldBindJSON(&lab); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 查询数据库
	labInfo, err := apiCfg.DB.ListLabById(c.Request.Context(), lab.LabID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"labInfo": labInfo,
	})
}

func (apiCfg *apiConfig) HandlerCreateLab(c *gin.Context) {
	var (
		err error
	)

	// 获取参数
	var lab struct {
		Name              string  `json:"name"`
		OfficeArea        float64 `json:"office_area"`
		Address           string  `json:"address"`
		ResearchDirection string  `json:"research_direction"`
	}
	if err = c.ShouldBindJSON(&lab); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 插入数据库
	labID, err := apiCfg.DB.CreateLab(c.Request.Context(), database.CreateLabParams{
		Name:              lab.Name,
		OfficeArea:        lab.OfficeArea,
		Address:           lab.Address,
		ResearchDirection: lab.ResearchDirection,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"labID": labID,
	})
}

func (apiCfg *apiConfig) HandlerDeleteLab(c *gin.Context) {
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

	// 删除数据库
	labInfo, err := apiCfg.DB.DeleteLab(c.Request.Context(), int32(lab.LabID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"labInfo": labInfo,
	})
}

func (apiCfg *apiConfig) HandlerUpdateLab(c *gin.Context) {
	var (
		err error
	)

	// 获取参数
	var lab struct {
		Name              string  `json:"Name"`
		OfficeArea        float64 `json:"OfficeArea"`
		Address           string  `json:"Address"`
		ResearchDirection string  `json:"ResearchDirection"`
		LabID             int32   `json:"LabID"`
	}
	if err = c.ShouldBindJSON(&lab); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 更新数据库
	labInfo, err := apiCfg.DB.UpdateLab(c.Request.Context(), database.UpdateLabParams{
		Name:              lab.Name,
		OfficeArea:        lab.OfficeArea,
		Address:           lab.Address,
		ResearchDirection: lab.ResearchDirection,
		LabID:             lab.LabID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"labInfo": labInfo,
	})
}

func (apiCfg *apiConfig) HandlerHealthzDatabase(c *gin.Context) {
	var (
		err error
	)

	// 查询数据库
	_, err = apiCfg.DB.HealthzDatabase(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// schema of table Directors(研究室主任表)
// -- 创建研究室主任表
// CREATE TABLE Directors (
//     DirectorID INT PRIMARY KEY,
//     LabID INT,
//     StartDate DATE,
//     Term INT,
//     FOREIGN KEY (DirectorID) REFERENCES Researchers(ResearcherID),
//     FOREIGN KEY (LabID) REFERENCES Laboratories(LabID)
// );
