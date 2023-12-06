package handler

import (
	"ResearchManage/internal/database"
	"database/sql"
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

func (apiCfg *apiConfig) HandlerCreateLab(c *gin.Context) {
	var (
		err error
	)

	// 获取参数
	var lab struct {
		Name              string  `json:"Name"`
		OfficeArea        float64 `json:"OfficeArea"`
		Address           string  `json:"Address"`
		ResearchDirection string  `json:"ResearchDirection"`
	}
	if err = c.ShouldBindJSON(&lab); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 插入数据库
	labID, err := apiCfg.DB.CreateLab(c.Request.Context(), database.CreateLabParams{
		Name: lab.Name,
		Officearea: sql.NullFloat64{
			Float64: lab.OfficeArea,
			Valid:   true,
		},
		Address: sql.NullString{
			String: lab.Address,
			Valid:  true,
		},
		Researchdirection: sql.NullString{
			String: lab.ResearchDirection,
			Valid:  true,
		},
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
		LabID             int     `json:"LabID"`
	}
	if err = c.ShouldBindJSON(&lab); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 更新数据库
	labInfo, err := apiCfg.DB.UpdateLab(c.Request.Context(), database.UpdateLabParams{
		Name: lab.Name,
		Officearea: sql.NullFloat64{
			Float64: lab.OfficeArea,
			Valid:   true,
		},
		Address: sql.NullString{
			String: lab.Address,
			Valid:  true,
		},
		Researchdirection: sql.NullString{
			String: lab.ResearchDirection,
			Valid:  true,
		},
		Labid: int32(lab.LabID),
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

func (apiCfg *apiConfig) HandlerListDirectorAll(c *gin.Context) {
	var (
		err error
	)

	// 查询数据库
	directorList, err := apiCfg.DB.ListDirectorAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"directorList": directorList,
	})
}

func (apiCfg *apiConfig) HandlerCreateDirector(c *gin.Context) {
	var (
		err error
	)

	// 获取参数
	var director struct {
		DirectorID int `json:"DirectorID"`
		LabID      int `json:"LabID"`
		Term       int `json:"Term"`
	}
	if err = c.ShouldBindJSON(&director); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 插入数据库
	directorID, err := apiCfg.DB.CreateDirector(c.Request.Context(), database.CreateDirectorParams{
		Directorid: int32(director.DirectorID),
		Labid:      int32(director.LabID),
		Term:       int32(director.Term),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"directorID": directorID,
	})
}

func (apiCfg *apiConfig) HandlerDeleteDirector(c *gin.Context) {
	var (
		err error
	)

	// 获取参数
	var director struct {
		DirectorID int `json:"DirectorID"`
	}
	if err = c.ShouldBindJSON(&director); err != nil {
		c.J/*  */SON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 删除数据库
	directorInfo, err := apiCfg.DB.DeleteDirector(c.Request.Context(), int32(director.DirectorID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"directorInfo": directorInfo,
	})
}

func (apiCfg *apiConfig) HandlerUpdateDirector(c *gin.Context) {
	var (
		err error
	)

	// 获取参数
	var director struct {
		DirectorID int `json:"DirectorID"`
		LabID      int `json:"LabID"`
		Term       int `json:"Term"`
	}
	if err = c.ShouldBindJSON(&director); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 更新数据库
	directorInfo, err := apiCfg.DB.UpdateDirector(c.Request.Context(), database.UpdateDirectorParams{
		Directorid: int32(director.DirectorID),
		Labid:      int32(director.LabID),
		Term:       int32(director.Term),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"directorInfo": directorInfo,
	})
}
