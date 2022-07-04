package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/nspforever/app-metadata-service/pkg/listing"
	"github.com/nspforever/app-metadata-service/pkg/models"
	"github.com/nspforever/app-metadata-service/pkg/upserting"
)

type Handler struct {
	router   *gin.Engine
	address  string
	upserter upserting.Service
	lister   listing.Service
}

func NewHandler(address string, upserter upserting.Service, lister listing.Service) *Handler {
	h := &Handler{
		address:  address,
		router:   gin.Default(),
		upserter: upserter,
		lister:   lister,
	}

	h.router.GET("/apps", h.getApps)
	h.router.PUT("/apps", h.upsertApps)

	return h
}

func (h *Handler) Run() {
	h.router.Run(h.address)
}

func (h *Handler) getApps(c *gin.Context) {
	apps, err := h.lister.GetApps()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, apps)
}

func (h *Handler) upsertApps(c *gin.Context) {

	var newApp models.AppMetadata

	if err := c.BindYAML(&newApp); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid app metadata"})
		return
	}
	h.upserter.UpsertApp(&newApp)
	c.Status(http.StatusOK)
}
