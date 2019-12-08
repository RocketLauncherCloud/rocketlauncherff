package main

import (
	"github.com/RocketLauncherFF/rocketlauncherff/core"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()
	datastore, err := NewDataStore(os.Getenv("DATASTORE_URL"))

	router.POST("/v1/flags", func(c *gin.Context) {
		var json core.FeatureFlag
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		_, err = datastore.Save(&json)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, json)
	})

	router.GET("/v1/flags/:name", func(c *gin.Context) {
		ff, err := datastore.Find(c.Param("name"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, ff)
	})

	router.PUT("/v1/flags/", func(c *gin.Context) {
		var json core.FeatureFlag
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		_, err = datastore.Update(&json)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, json)
	})

	router.DELETE("/v1/flags/:id", func(c *gin.Context) {
		err := datastore.Delete(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusOK)
	})

	router.GET("/v1/flags", func(c *gin.Context) {
		flags, err := datastore.FindAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, flags)
	})

	router.Run(":8089")
}
