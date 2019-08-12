.PHONY: gen-go package

IMAGE_NAME=192.168.0.103:8080/axe/protos
VERSION=v0.0.1
GIT_COMMIT=$(shell git rev-parse HEAD)

SWAGGER_FILES=$(wildcard swagger/*/*.json)

default: gen-proto

package:
	docker build --build-arg VERSION=${VERSION} --build-arg GIT_COMMIT=${GIT_COMMIT} -t ${IMAGE_NAME}:${VERSION} .

push: package
	docker push ${IMAGE_NAME}:${VERSION}

gen-proto: clean
	docker run -it -v $(PWD):/opt/protos ${IMAGE_NAME}:${VERSION} sh gen.sh

swagger-ui:
	swagger mixin ${SWAGGER_FILES} -o swagger.json -q || true
	swagger serve swagger.json -Fswagger

clean:
	@rm -rf goout javaout swagger