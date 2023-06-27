package controllers

import "github.com/gin-gonic/gin"

type Controller interface {
	List(c *gin.Context)
	Retrieve(c *gin.Context)
	Create(c *gin.Context)
	Delete(c *gin.Context)
}
