package main

import (
	"github.com/nspforever/app-metadata-service/pkg/http/rest"
	"github.com/nspforever/app-metadata-service/pkg/listing"
	"github.com/nspforever/app-metadata-service/pkg/storage/memory"
	"github.com/nspforever/app-metadata-service/pkg/upserting"
)

func main() {
	var upserter upserting.Service
	var lister listing.Service

	s := memory.New()
	upserter = upserting.NewService(s)
	lister = listing.NewService(s)

	// set up the HTTP server
	hander := rest.NewHandler("localhost:9999", upserter, lister)
	hander.Run()
}
