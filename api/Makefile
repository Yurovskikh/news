build:
	@go fmt ./...
	@GOOS=linux GOARCH=amd64 go build ./cmd/api.go

pb:
	@protoc ./pb/*.proto --go_out=plugins=grpc:.