package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/flaviostutz/wfsgis/model"
	"github.com/flaviostutz/wfsgis/ogc"
	"github.com/gin-gonic/gin"
)

type FeatureHandler struct {
	Store *model.DB
}

func (h *HTTPServer) makeFeatureHandlers(d *model.DB) {
	h.router.GET("/collections/:collid/items", getFeatures(d))
	h.router.GET("/collections/:collid/items/:itemid", getFeatureById(d))
	h.router.POST("/collections/:collid/items", createFeature(d))
	h.router.PUT("/collections/:collid/items/:itemid", updateFeature(d))
	h.router.DELETE("/collections/:collid/items/:itemid", deleteFeature(d))
}

/**
Updates a feature
*/
func updateFeature(db *model.DB) func(*gin.Context) {
	return func(c *gin.Context) {

	}
}

/**
Creates a feature
*/
func createFeature(db *model.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		collid := c.Param("collid")
		collid = strings.ReplaceAll(collid, ":", "$")

		var fc ogc.FeatureCollection
		data, _ := ioutil.ReadAll(c.Request.Body)
		err := json.Unmarshal(data, &fc)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		log.Printf("createFeature %v %v", collid, fc.Features)
		if len(fc.Features) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "No features found in json body"})
			return
		}

		nids, err0 := db.InsertFeature(collid, fc.Features)
		if err0 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err0.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Features inserted successfully", "ids": nids})
	}
}

/**
Deletes a feature
*/
func deleteFeature(db *model.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		collid := c.Param("collid")
		collid = strings.ReplaceAll(collid, ":", "$")
		itemid := c.Param("itemid")
		if err := db.DeleteItem(collid, itemid); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"success": "true"})
		}
	}
}

/**
Gets a feature by id
*/
func getFeatureById(db *model.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		collid := c.Param("collid")
		collid = strings.ReplaceAll(collid, ":", "$")
		itemid := c.Param("itemid")
		if item, err := db.GetItem(collid, itemid); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		} else {
			c.Header("Cache-Control", "public, max-age=7776000")
			c.JSON(http.StatusOK, item)
		}
	}
}

/**
Gets features
*/
func getFeatures(db *model.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		collid := c.Param("collid")
		collid = strings.ReplaceAll(collid, ":", "$")
		limitStr := c.DefaultQuery("limit", "100")
		timeStr := c.Query("time")
		bboxStr := c.Query("bbox")
		params := c.Request.URL.Query()

		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid limit parameter"})
			return
		}

		filterAttrs := make(map[string]string)
		for k, v := range params {
			if k != "time" && k != "bbox" && k != "limit" {
				filterAttrs[k] = v[0]
			}
		}

		var dateStart *time.Time
		var dateEnd *time.Time
		if timeStr != "" {
			ts := strings.Split(timeStr, "/")
			if len(ts) == 1 {
				d, err := time.Parse(time.RFC3339, completeDate(ts[0], true))
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Invalid time parameter. date0=%s", ts[0])})
					return
				}
				dateStart = &d
				dateEnd = &d
			} else {
				if ts[0] != "" {
					d, err := time.Parse(time.RFC3339, completeDate(ts[0], true))
					if err != nil {
						c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Invalid time parameter. date1=%s", ts[0])})
						return
					}
					dateStart = &d
				}
				if ts[1] != "" {
					d, err := time.Parse(time.RFC3339, completeDate(ts[1], false))
					if err != nil {
						c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Invalid time parameter. date2=%s", ts[1])})
						return
					}
					dateEnd = &d
				}
			}
		}

		var bbox *ogc.Bbox
		if bboxStr != "" {
			bstr := strings.Split(bboxStr, ",")
			if len(bstr) != 4 {
				c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid bbox parameter. Should be something like bbox=-13.45,23,-13.895,23.45"})
				return
			}
			b0, err0 := strconv.ParseFloat(bstr[0], 64)
			b1, err1 := strconv.ParseFloat(bstr[1], 64)
			b2, err2 := strconv.ParseFloat(bstr[2], 64)
			b3, err3 := strconv.ParseFloat(bstr[3], 64)
			if err0 != nil || err1 != nil || err2 != nil || err3 != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid bbox parameter. Should be something like bbox=-13.45,23,-13.895,23.45"})
				return
			}

			bbox = ogc.NewBbox(b0, b1, b2, b3)
		}

		fc, err := db.GetFeatures(collid, bbox, filterAttrs, limit, dateStart, dateEnd)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.Header("Cache-Control", "public, max-age=86400")
		c.JSON(http.StatusOK, gin.H{"type": "FeatureCollection", "features": fc})
	}

}

func completeDate(dateStr string, midnight bool) string {
	if len(dateStr) == 10 {
		if midnight {
			return dateStr + "T00:00:00Z"
		} else {
			return dateStr + "T23:59:59Z"
		}
	}
	return dateStr
}
