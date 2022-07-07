MAIN_PKG:=app-metadata-service
BIN=$(strip $(MAIN_PKG))
MOCKGEN=$(GOPATH)/bin/mockgen
MOCK_DIR=$(CURDIR)/pkg/mocks

all: build test

clean:
	@rm -rf bin *.out

tools:
	go get golang.org/x/tools/cmd/goimports
	go get github.com/kisielk/errcheck
	go get github.com/axw/gocov/gocov
	go get github.com/matm/gocov-html
	go get github.com/tools/godep
	go get github.com/mitchellh/gox
	go get golang.org/x/lint/golint
	go get github.com/golang/mock/mockgen@latest

deps:
	go mod download
	go mod tidy

install:
	go mod tidy
	go install ./...

build: deps
	go build -o bin/$(BIN) github.com/nspforever/$(MAIN_PKG)/cmd/server

run: build
	bin/$(BIN)

test: deps
	go test -v -coverprofile=$(BIN).out ./...

# example: make test-package P=github.com/nspforever/app-metadata-service/pkg/storage/memory
test-package:
	go test -v $(P) -coverprofile=$(BIN).out

# example: make test-func P=github.com/nspforever/app-metadata-service/pkg/storage/memory T=TestUpsertApp
test-func:
	go test -v $(P) -run ^$(T)$$ -coverprofile=taskplanner_coverage.out

smoke-test: deps
	go test --tags=smoke -v github.com/nspforever/app-metadata-service/pkg/tests

coverage: deps
	gocov test ./... > $(CURDIR)/coverage.out 2>/dev/null
	gocov report $(CURDIR)/coverage.out
	if test -z "$$CI"; then \
	  gocov-html $(CURDIR)/coverage.out > $(CURDIR)/coverage.html; \
	  if which open &>/dev/null; then \
	    open $(CURDIR)/coverage.html; \
	  fi; \
	fi

vet:
	@go vet ./...

errors:
	errcheck -ignoretests -blank ./...

lint:
	golint ./...

imports:
	goimports -l -w .

fmt:
	@go fmt ./...

mock:
	$(MOCKGEN) -source=$(GOPATH)/src/github.com/nspforever/app-metadata-service/pkg/upserting/service.go \
						 -destination=$(GOPATH)/src/github.com/nspforever/app-metadata-service/pkg/mocks/upserting/mock_upserter.go \
						 -package=upserting
	$(MOCKGEN) -source=$(GOPATH)/src/github.com/nspforever/app-metadata-service/pkg/searching/service.go \
						 -destination=$(GOPATH)/src/github.com/nspforever/app-metadata-service/pkg/mocks/searching/mock_searcher.go \
						 -package=searching
	$(MOCKGEN) -source=$(GOPATH)/src/github.com/nspforever/app-metadata-service/pkg/storage/memory/repository.go \
						 -destination=$(GOPATH)/src/github.com/nspforever/app-metadata-service/pkg/mocks/storage/memory/mock_apps_filters_applier.go \
						 -package=memory

pre-checkin: fmt vet errors imports build test lint

docker:
	docker build -t $(MAIN_PKG) .

docker-run:
	docker run -d -p 9999:9999 --rm --name app-metadata-service app-metadata-service:latest

