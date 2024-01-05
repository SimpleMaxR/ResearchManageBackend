package handler

import "github.com/gin-gonic/gin"

func (apiCfg *apiConfig) HandlerOverview(c *gin.Context) {
	projectCount, err := apiCfg.DB.CountProject(c.Request.Context())
	researcherCount, err := apiCfg.DB.CountResearcher(c.Request.Context())
	labCount, err := apiCfg.DB.CountLab(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"projectCount":    projectCount,
		"researcherCount": researcherCount,
		"labCount":        labCount,
	})
}
