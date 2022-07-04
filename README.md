### Run Server
```
cd cmd/server
go run .
```

### Add or modify app metadata
```
curl -X PUT --data-binary @test-data/invalid_app1.yaml -H "Content-type: application/x-yaml" http://localhost:9999/apps
```

### List apps
```
curl localhost:9999/apps
```
