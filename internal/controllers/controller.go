package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"example.com/m/internal/domain"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	hubService domain.HubService
}

var ctrl *Controller
var ctrlOnce sync.Once

func NewController(svc domain.HubService) *Controller {
	ctrlOnce.Do(func() {
		ctrl = &Controller{
			hubService: svc,
		}
	})
	return ctrl
}

func (hubController *Controller) GetHub(c *gin.Context) {
	hubID := c.Param("hub_id")
	hubIDUint, err := strconv.ParseUint(hubID, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch hub id"})
		return
	}

	hubDetails, cusErr := hubController.hubService.GetHubDetails(c, hubIDUint)
	if cusErr.Exists() {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hub not found"})
		return
	}
	fmt.Println("hubDetails", hubDetails)

	c.JSON(http.StatusOK, gin.H{"hub": hubDetails})

}
