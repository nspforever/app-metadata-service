MAIN_PKG:=app-metadata-service
BIN=$(strip $(MAIN_PKG))

GO_FILES=$(shell find . -type f -name '*.go' -not -path "./vendor/*")

# mac only
.PHONY: dep-install
dep-install:
ifeq ($(OS), Darwin)
	@which goimports > /dev/null || go get -u golang.org/x/tools/cmd/goimports
	@which gosumcheck > /dev/null || go get -u golang.org/x/mod/gosumcheck
	@go get github.com/golang/mock/mockgen@latest
	@echo 'all dependencies installed, you are good to go'
endif

# download mod package
init:
	go mod download
	go mod tidy

.PHONY: install
install:
	go mod tidy
	go install -v ./...


build: init
	go build -o bin/$(BIN) github.com/nspforever/$(MAIN_PKG)/cmd/server

run: build
	bin/$(BIN)

test: init
	go test -v -coverprofile=$(BIN).out ./...

# example: make test-package P=github.com/nspforever/app-metadata-service/pkg/storage/memory
.PHONY: test-package
test-package:
	go test -v $(P) -coverprofile=$(BIN).out

# example: make test-func P=github.com/nspforever/app-metadata-service/pkg/storage/memory T=TestUpsertApp
.PHONY: test-func
test-func:
	go test -v $(P) -run ^$(T)$$ -coverprofile=taskplanner_coverage.out

.PHONY: fmt
fmt:
	@go fmt ./...

.PHONY: vet
vet:
	@go vet ./...

.PHONY: clean
clean:
	@rm -rf .gen bin mocks

.PHONY: imports
imports:
	goimports -w $(GO_FILES)