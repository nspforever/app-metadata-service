package main

import (
	"github.com/nspforever/app-metadata-service/pkg/http/rest"
	"github.com/nspforever/app-metadata-service/pkg/searching"
	"github.com/nspforever/app-metadata-service/pkg/storage/memory"
	"github.com/nspforever/app-metadata-service/pkg/upserting"
)

func main() {
	var upserter upserting.Service
	var searcher searching.Service

	s := memory.New()
	upserter = upserting.NewService(s)
	searcher = searching.NewService(s)

	// set up the HTTP server
	hander := rest.NewHandler("localhost:9999", upserter, searcher)
	hander.Run()
}
