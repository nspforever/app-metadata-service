# API design

## Add/Modify App

Add new app metadata or modify an existing app metadata. App metadata is unique identified by title@version. Supports both JSON and YAML as payload.

### **Request**
```
PUT  http://<hostname>:[port]/api/v1/apps
```

**Payload:**

**YAML**
```
title: Valid App1
version: 0.0.1
maintainer:
- name: firstmaintainer app1
  email: firstmaintainer@hotmail.com
company: Random Inc
website: https://website.com
source: https://github.com/random/repo
license: Apache-2.0
description: |
  ### Interesting Title,
  Some application content, and description
```

**JSON**
```
{
  "title": "Valid App1",
  "version": "0.0.1",
  "maintainer": [
    {
      "name": "firstmaintainer app1",
      "email": "firstmaintainer@hotmail.com"
    }
  ],
  "company": "Random Inc",
  "website": "https://website.com",
  "source": "https://github.com/random/repo",
  "license": "Apache-2.0",
  "description": "### Interesting Title,\nSome application content, and description\n"
}
```

### Response

**200 OK {}**

**400 Bad request {error message ...}**

**500 Internal Server Error {error message ...}**


## Search apps

Search for apps metadata that previously created. Search criteria is specified by query parameters. The relationship between query parameters is 'AND'. If no query parameter is provied, all app metadata will be returned.

### Request

```
GET http://<hostname>:[port]/api/v1/apps?[query_params]
```

**Supported Parameters**

|Field|Usage|Description|
| ----- | ---- | ---- |
|title|title=app1|exact match the given string|
|version|version=1.0.0|exact match the given string|
|maintainer_has_name|maintainer_has_name=John%20Doe|one of the maintainer's name exact match the given string|
|maintainer_has_email|maintainer_has_name=john.doe%40outlook.com|one of the maintainer's email exact match the given string|
|company|company=contoso.com|exact match the given string|
|website|website=https%3A%2F%2Fcontoso.com|exact match the given string|
|source|source=https%3A%2F%2Fgithub.com/contoso/app1|exact match the given string|
|license|license=Apache-2.0|exact match the given string|
|description_has_text|description_has_text=content|description has the given text|

### Response

**400 Bad request {error message ...}**

**404 Not found {error message ...}**

**500 Internal Server Error {error message ...}**

**200 OK**
```
{
  "count": int,
  "data": [
    {
      "company": string,
      "description": string,
      "license": string,
      "maintainer": [
        {
          "email": string,
          "name": string
        }
      ],
      "source": string,
      "title": string,
      "version": string,
      "website": string
    }
  ],
  "limit": int,
  "offset": int
}
```

# Run the service

## Run the server

1. Run on local devbox with go environment

```
make run
```

2. alternatively run the service in docker

```
make docker-run
```

## Populate the test data
You can run below command to populate the test data
```
curl -X PUT --data-binary @test-data/app1.yaml -H "Content-type: application/x-yaml" http://localhost:9999/apps
curl -X PUT --data-binary @test-data/app2.yaml -H "Content-type: application/x-yaml" http://localhost:9999/apps
curl -X PUT --data-binary @test-data/app3.yaml -H "Content-type: application/x-yaml" http://localhost:9999/apps
```

Or run:
```
make smoke-test
```

## Search apps
```
curl localhost:9999/apps

curl http://localhost:9999/apps\?title\=Valid%20App1

curl http://localhost:9999/apps\?maintainer_has_email\=firstmaintainer%40hotmail.com\&company\=Random%20Inc\&license\=Apache-2.0

```

# Local Development

- Clone the code git@github.com:nspforever/app-metadata-service.git
- make build
- make test
- make pre-checkin

## Run tests
- **Run all tests**
  ```
  make test
  ```

- **Run test of a specific package**

  ```
  make test-package P=<path-of-package>

  e.g. make test-package P=github.com/nspforever/app-metadata-service/pkg/storage/memory
  ```

- **Run a specific test function**

  ```
  make test-func P=<path-of-package> T=<test-function-name>
  # example: make test-func P=github.com/nspforever/app-metadata-service/pkg/storage/memory T=TestUpsertApp
  ```

## Run Server
```
make run
```

## Run smoke test
1. run the server in one terminal window or in docker
  ```
  make run
  ```

2. Run the smoke tests
  ```
  make smoke-test
  ```

## Generate mocks
```
make mock
```

## Collect code coverage data
```
make coverage
```

## Run pre-checkin validation
Below command will run errcheck, goimports, go vet, go fmt, lint, build, test to the code quality
```
make pre-check
```


# Future TODOs
- Validation on version number
- Search by sematic version number
- Add dependency injection
- Support pagination and sort
- Support 'or' filters for listing API
- Benchmark test for the API
- Enable CI/CD

