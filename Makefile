MAIN_PKG:=app-metadata-service
BIN=$(strip $(MAIN_PKG))
MOCKGEN=$(GOPATH)/bin/mockgen
MOCK_DIR=$(CURDIR)/pkg/mocks

all: build test

clean:
	@rm -rf bin *.out

tools:
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/kisielk/errcheck@latest
	go install github.com/axw/gocov/gocov@latest
	go install github.com/matm/gocov-html@latest
	go install github.com/mitchellh/gox@latest
	go install golang.org/x/lint/golint@latest
	go install github.com/golang/mock/mockgen@v1.6.0

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

docker-run: docker
	docker run -d -p 9999:9999 --rm --name app-metadata-service app-metadata-service:latest

