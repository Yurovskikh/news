build:
	@cd api/ && GOOS=linux GOARCH=amd64 go build -o api
	@cd storage/ && GOOS=linux GOARCH=amd64 go build -o storage

start:build
	@docker-compose up -d --force-recreate --build

pb:
	@protoc ./pb/*.proto --go_out=plugins=grpc:.