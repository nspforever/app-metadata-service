package server

import (
	"github.com/gin-gonic/gin"
)

type StorageType int
const (
	Memory Type = iota
	JSON
)

func main() {
	var adder adding.Service
	var lister listing.Service

	router := rest.Handler(adder, lister, reviewer)
	fmt.Println("The beer server is on tap now: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}


type DBClient interface {
	Store(*AppMetadata) error
	Query() ([]*AppMetadata, error)
}

type HTTPServer struct {
	db 	DBClient
}

func (s *HTTPServer) Run() error {
	router := gin.Default()
	router.GET("/apps", getApps)

	return router.Run("localhost:9999")
}
router := gin.Default()
	router.GET("/apps", getApps)
	router.Run("localhost:9999")