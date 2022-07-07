package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/martian/log"

	"github.com/nspforever/app-metadata-service/pkg/filtering/app"
	"github.com/nspforever/app-metadata-service/pkg/models"
	"github.com/nspforever/app-metadata-service/pkg/searching"
	"github.com/nspforever/app-metadata-service/pkg/upserting"
)

// Handler repesents a http handler
type Handler struct {
	router   *gin.Engine
	address  string
	upserter upserting.Service
	searcher searching.Service
}

// NewHandler initialize an http handler
func NewHandler(address string, upserter upserting.Service, searcher searching.Service) *Handler {
	h := &Handler{
		address:  address,
		router:   gin.Default(),
		upserter: upserter,
		searcher: searcher,
	}

	h.router.GET("api/v1/apps", h.searchApps)
	h.router.PUT("api/v1/apps", h.upsertApps)

	return h
}

// Run starts the http server
func (h *Handler) Run() error {
	return h.router.Run(h.address)
}

func (h *Handler) upsertApps(c *gin.Context) {
	var newApp models.AppMetadata

	if err := c.ShouldBind(&newApp); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid app metadata " + err.Error()})
		log.Errorf("invalid app metadata %+v \n", err)
		return
	}
	if err := h.upserter.UpsertApp(&newApp); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		log.Errorf("failed to persist app metadata %+v\n", err)
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) searchApps(c *gin.Context) {
	params := make(map[string]string)
	if err := c.BindQuery(&params); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		log.Errorf("malform search filter %+v\n", err)
		return
	}

	filters, err := getFilters(params)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		log.Errorf("error: %+v\n", err)
		return
	}

	apps, err := h.searcher.SearchApps(filters)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		log.Errorf("an error returned by app searcher %+v\n", err)
		return
	}

	c.IndentedJSON(http.StatusOK,
		&models.AppSearchResponse{
			Count: len(apps),
			Data:  apps,
		})
}

func getFilters(params map[string]string) (filters *app.Filters, err error) {
	var opts []app.FilterOption
	for name, value := range params {
		var opt app.FilterOption
		opt, err = app.FilterOptionFactory(name, value)
		if err != nil {
			break
		}
		opts = append(opts, opt)
	}
	filters = app.NewFilters(opts...)
	return
}
