
GO_MODULE_ON=GO111MODULE=on
GO_ENV=${GO_MODULE_ON} GOOS=linux GOARCH=amd64 CGO_ENABLED=0

.PHONY: build-helper
build-helper:
	mkdir -p ./tmp/docker/
	rm -rf ./tmp/docker/helper
	${GO_ENV} go build -ldflags '-w -s' -v -tags netgo -o ./tmp/docker/helper github.com/stoneshi-yunify/local-path-pod-helper/cmd/helper/...

.PHONY: docker-build
docker-build: build-helper
	docker build -t local-path-helper -f ./build/docker/helper/Dockerfile ./tmp/docker/
	docker build -t local-path-provisioner -f ./build/docker/provisioner/Dockerfile ./tmp/docker/
