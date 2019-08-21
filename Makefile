.PHONY: gen-go package

IMAGE_NAME=192.168.0.103:8080/axe/protos
VERSION=v0.0.2
GIT_COMMIT=$(shell git rev-parse HEAD)

SWAGGER_FILES=$(wildcard swagger/*.json) $(wildcard swagger/*/*.json)

default: gen-proto swagger-mixin

package:
	docker build --build-arg VERSION=${VERSION} --build-arg GIT_COMMIT=${GIT_COMMIT} -t ${IMAGE_NAME}:${VERSION} .

push: package
	docker push ${IMAGE_NAME}:${VERSION}

gen-proto: clean
	docker run -it -v $(PWD):/opt/protos ${IMAGE_NAME}:${VERSION} sh gen.sh

swagger-mixin:
	@docker run --rm -it -v $(PWD):/tmp/protos -w /tmp/protos quay.io/goswagger/swagger -q mixin ${SWAGGER_FILES} -o swagger/swagger.json || true

swagger-ui:
	@docker run --rm -it -v $(PWD):/tmp/protos -p 8080:8080 -w /tmp/protos quay.io/goswagger/swagger serve swagger/swagger.json -Fswagger --port 8080 --no-open

clean:
	@rm -rf goout javaout swagger