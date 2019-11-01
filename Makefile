.PHONY: gen-go package

IMAGE_NAME=192.168.0.103:8080/axe/protos
VERSION=v0.0.3
GIT_COMMIT=$(shell git rev-parse HEAD)

SWAGGER_FILES=$(wildcard swagger/*.json) $(wildcard swagger/*/*.json)

default: gen-proto swagger-mixin gen-doc

package:
	docker build --build-arg VERSION=${VERSION} --build-arg GIT_COMMIT=${GIT_COMMIT} -t ${IMAGE_NAME}:${VERSION} .

push: package
	docker push ${IMAGE_NAME}:${VERSION}

gen-proto:
	@rm -rf goout
	docker run -it -v `pwd`:/opt/protos ${IMAGE_NAME}:${VERSION} sh gen.sh proto

gen-doc:
	@rm -rf doc 
	docker run -it -v `pwd`:/opt/protos ${IMAGE_NAME}:${VERSION} sh gen.sh doc

gen-swagger:
	@rm -rf swagger
	docker run -it -v `pwd`:/opt/protos ${IMAGE_NAME}:${VERSION} sh gen.sh swagger

gen-mock:
	@rm -rf gooutmock
	docker run -it -v `pwd`:/opt/protos ${IMAGE_NAME}:${VERSION} sh gen.sh mock

swagger-mixin: gen-swagger
	@docker run --rm -it -v `pwd`:/tmp/protos -w /tmp/protos 192.168.0.103:8080/axe/goswagger:latest -q mixin ${SWAGGER_FILES} -o swagger/swagger.json || true

swagger-ui:
	@docker run --rm -it -v `pwd`:/tmp/protos -p 8080:8080 -w /tmp/protos 192.168.0.103:8080/axe/goswagger:latest serve swagger/swagger.json -Fswagger --port 8080 --no-open

clean:
	@rm -rf goout javaout swagger