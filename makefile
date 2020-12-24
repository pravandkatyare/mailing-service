all: clean vendor build 
	bin/mailing-service

run:
	go run main.go

build:
	env GOOS=linux GOARCH=amd64 go build -o bin/mailing-service main.go
	chmod +x bin/mailing-service

build-mac: 
	env GOOS=freebsd GOARCH=amd64 go build -o bin/mailing-service main.go
	chmod +x bin/mailing-service

build-win:
	env GOOS=windows GOARCH=amd64 go build -o bin/mailing-service main.go
	chmod +x bin/mailing-service

clean:
	rm -rf ./gen ./bin ./vendor

vendor:
	go mod vendor

