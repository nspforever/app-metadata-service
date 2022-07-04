


## API design

### Add/Modify App
- Add and modify app via the same endpoint(upsert)
- Use app title and app version to unique identify an App
- Support both YAML and JSON payload

#### API signature

### List apps

#### supported filter

#### API signature

## Local Development

- Go version 1.15+
- Clone the code git@github.com:nspforever/app-metadata-service.git
- make build
- make test

### Run unit tests
- Run all tests
```
make test
```

- Run test of a specific package

```
make test-package P=<path-of-package>

e.g. make test-package P=github.com/nspforever/app-metadata-service/pkg/storage/memory
```

- Run a specific test function

```
make test-func P=<path-of-package> T=<test-function-name>
# example: make test-func P=github.com/nspforever/app-metadata-service/pkg/storage/memory T=TestUpsertApp
```

## Local E2E test


### Run Server
```
cd cmd/server
go run .
```

### Add or modify app metadata
- Return all apps if not filter is provided(unsafe)
- Return apps by that meet the filter criteria
- Relationship between filters are always 'AND'
- Supported filters

title eq
version eq(gte, lte, gt, lt)
maintainers has
company eq
website eq
source eq
license eq
description has, eq


```
curl -X PUT --data-binary @test-data/invalid_app1.yaml -H "Content-type: application/x-yaml" http://localhost:9999/apps
```

### List apps
```
curl localhost:9999/apps
```

### TODO
- options pattern
- query filter design

### Future TODO
- Validation on version number
- Search by sematic version number
- Add dependency injection
- Support pagination
- Support 'or' filters for listing API
- Benchmark test for the API
- Enable CI/CD

