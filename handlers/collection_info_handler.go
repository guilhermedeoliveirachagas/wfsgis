package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/boundlessgeo/wfs3/model"
	"github.com/boundlessgeo/wfs3/ogc"
	"github.com/gin-gonic/gin"
)

func (h *HTTPServer) makeCollectionHandlers(d *model.DB) {
	h.router.GET("/collections", getCollectionsInfo(d))
	h.router.GET("/collections/:collid/schema", getCollectionInfo(d))
	h.router.PUT("/collections/:collid/schema", updateCollectionInfo(d))
	h.router.POST("/collections", createCollectionInfo(d))
	h.router.DELETE("/collections/:collid", deleteCollection(d))
}

/**
Deletes a collection
*/
func deleteCollection(db *model.DB) func(*gin.Context) {
	log.Printf("Not implemented yet")
	return nil
}

/**
Updates a collection
*/
func updateCollectionInfo(db *model.DB) func(*gin.Context) {
	log.Printf("Not implemented yet")
	return nil
}

/**
Gets all collections
*/
func getCollectionsInfo(db *model.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		cidbs := db.AllCollectionInfos()
		cis := make([]*ogc.CollectionInfo, 0)
		for _, v := range cidbs {
			v.CollectionInfo.Name = strings.ReplaceAll(v.CollectionInfo.Name, "$", ":")
			cis = append(cis, v.CollectionInfo)
		}
		c.JSON(http.StatusOK, gin.H{"collections": cis})
	}
}

func getCollectionInfo(db *model.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		collid := c.Param("collid")
		collid = strings.ReplaceAll(collid, "$", ":")
		collInfo, err := db.FindCollection(collid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("err=%s", err)})
		}
		collInfo.CollectionInfo.Name = strings.ReplaceAll(collInfo.CollectionInfo.Name, ":", "$")
		c.JSON(http.StatusOK, gin.H{"result": collInfo})
	}
}

/**
Creates a collection
*/
func createCollectionInfo(db *model.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var coll ogc.CollectionInfo
		if inputErr := c.BindJSON(&coll); inputErr != nil {
			c.JSON(http.StatusBadRequest, inputErr.Error)
			return
		}
		coll.Name = strings.ReplaceAll(coll.Name, ":", "$")
		collDB := model.CollectionInfoDB{CollectionInfo: &coll}
		err := db.AddCollection(&collDB)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		err = db.CreateCollectionTable(coll.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusCreated, "")
	}
}
