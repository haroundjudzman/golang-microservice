.DEFAULT_GOAL := swagger

install_swagger:
	which swagger || go install github.com/go-swagger/go-swagger/cmd/swagger

swagger: install_swagger
	swagger generate spec -o ./swagger.yaml --scan-models