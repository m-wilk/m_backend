WGEN_SERVICE_NAME=w_gen


wgen-build-api:
	docker compose exec ${WGEN_SERVICE_NAME} env GOOS=linux CGO_ENABLED=0 go build -o api cmd/api/main.go 

wgen-build-scripts:
	docker compose exec ${WGEN_SERVICE_NAME} env GOOS=linux CGO_ENABLED=0 go build -o scripts cmd/scripts/main.go 

wgen-build: wgen-build-api wgen-build-scripts
	@echo "build all wgen services"

# ALL
build-api: wgen-build-api
	@echo "build all services api"

# ALL
build-script: wgen-build-scripts wgen-build-api
	@echo "build all services scripts"
