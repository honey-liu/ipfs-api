package main

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"ipfs_api/api"
	"net/http"
)
//
var router *gin.Engine

type apiDoc struct {
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}

func InitApp() {
	router = gin.Default()
	router.Use(corsMiddleware)
	router.Use(gzip.Gzip(gzip.BestCompression))
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, map[string]string{"message": "Welcome to use ipfs api!"})
	})
	router.GET("/ipfs-api", welcome)
	v1:= router.Group("/ipfs-api")
	{
		v1.POST("/upload",api.UploadFile)
		//v1.GET("/catIPFS",api.)
	}
	router.NoRoute(NoRouteHandler)
	router.HandleMethodNotAllowed = true
	router.NoMethod(NoMethodHandler)
	router.Run(":80")
}

func welcome(c *gin.Context) {
	var apiList []apiDoc
	for _, r := range router.Routes() {
		apiList = append(apiList, apiDoc{Method: r.Method, Endpoint: r.Path})
	}
	c.JSON(200, apiList)
}

func NoRouteHandler(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotFound)
}
func NoMethodHandler(c *gin.Context) {
	c.AbortWithStatus(http.StatusMethodNotAllowed)
}

func corsMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", c.GetHeader("Origin"))
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Next()
}


