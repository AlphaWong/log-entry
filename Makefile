cover: 
	go test -coverprofile=coverage.out -covermode=atomic ./...

coverage:
	go test -tags=coverage -coverpkg=./utilshttp/...,./jsonschema/... ./... -coverprofile coverage.out && go tool cover -func=coverage.out

coveragehtml: coverage
	go tool cover -html=coverage.out -o public/coverage.html

coverhtml: coverage
	go tool cover -html=coverage.out