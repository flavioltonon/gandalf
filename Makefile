build:
	@env GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64 \
	go build -ldflags '-s -w' -o ./bin/gandalf ./cmd

check:
	@go vet ./...	

image:
	@docker build -t flavioltonon/gandalf:latest .

push:
	@docker push flavioltonon/gandalf:latest

release: image push

start: stop
	@docker-compose up --build

stop:
	@docker-compose down --remove-orphans

tidy:
	@go fmt ./...