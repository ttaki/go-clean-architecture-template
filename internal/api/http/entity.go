/*
 * Swagger MapSNS
 *
 * gobackend api
 *
 * API version: 1.0.0
 * Contact: apiteam@swagger.io
 * Generated by: OpenAPI Generator (https://http-generator.tech)
 */

package http

import (
	"github.com/gin-gonic/gin"
	"github.com/ttaki/go-clean-architecture-template/internal/controller"
	"github.com/ttaki/go-clean-architecture-template/pkg/logger"
)

var Entity = controller.NewEntity(logger.GetLogger())

// CreateEntity -
func CreateEntity(c *gin.Context) {
	Entity.Create(c)
}

// GetEntity - Your GET endpoint
func GetEntity(c *gin.Context) {
	Entity.FindByID(c)
}

// GetEntities - Your GET endpoint
func GetEntities(c *gin.Context) {
	Entity.FindAll(c)
}

// GetEntitiesByArea - Your GET endpoint

// UpdateEntity -
func UpdateEntity(c *gin.Context) {
	Entity.Update(c)
}
