

run: 
	go run main.go

deps: 
	go get -u "cloud.google.com/go/logging/v2"

tidy:
	go mod tidy