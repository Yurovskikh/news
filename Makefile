start:
	@docker-compose up -d --force-recreate --build

pb:
	@protoc ./pb/*.proto --go_out=plugins=grpc:.