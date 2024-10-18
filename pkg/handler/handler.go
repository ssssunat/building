package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	service "github.com/ssssunat/pkg/service"
)

type Handler struct {
	service *service.BuildingService
}

func NewHandler(service *service.BuildingService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/buildings", h.GetAllBuilding)
	router.POST("/buildings", h.CreateBuilding)

	return router
}

func (h *Handler) GetAllBuilding(c *gin.Context) {
	buildings, err := h.service.GetAllBuilding()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, buildings)
	// var buildings []Building

	// // Optional filters
	// city := c.Query("city")
	// year := c.Query("year")
	// floors := c.Query("floors_count")

	// query := `SELECT id, name, city, year, floors_count FROM buildings WHERE 1=1`
	// var args []interface{}

	// // Add filters to query dynamically
	// if city != "" {
	// 	query += " AND city = ?"
	// 	args = append(args, city)
	// }
	// if year != "" {
	// 	query += " AND year = ?"
	// 	args = append(args, year)
	// }
	// if floors != "" {
	// 	query += " AND floors_count = ?"
	// 	args = append(args, floors)
	// }

	// // Execute the query with filtering
	// err := db.Select(&buildings, query, args...)
}

func (h *Handler) CreateBuilding(c *gin.Context) {
	var building service.Building
	if err := c.BindJSON(&building); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	building, err := h.service.CreateBuilding(building)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, building)
}
