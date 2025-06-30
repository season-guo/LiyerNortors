package handler

import (
	"LiyerNortorsAIpart/internal/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func Save(c *gin.Context){
	var req models.SaveReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Err" : err})
		return
	}
	
	img, err  := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Err" : err})
		return
	}

	uid, err := GetUid(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Err" : err})
	}

	if err := models.Save(uid, img, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Err" : err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Status" : "Ok"})
}

func Analyze(c *gin.Context){
	var req models.AnalyzeReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Err" : err})
		return
	}

	uid, err := GetUid(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Err" : err})
	}

	ResultCid, err := models.Analyze(uid, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Err" : err})
		return
	}	

	c.JSON(http.StatusOK, gin.H{"result" : ResultCid}) 
}