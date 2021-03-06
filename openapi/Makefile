GENERATOR=go-gin-server
VERSION=v1.0
# oepanpi.ymlからソースコードを生成
gen:
	docker run --rm -v ${PWD}/:/local openapitools/openapi-generator-cli generate -i /local/${VERSION}/spec.yaml -g ${GENERATOR} -o /local/${VERSION}/generated/${GENERATOR} 
	rm -rf ./${VERSION}/generated/${GENERATOR}/go/api_*
	rm -rf ./${VERSION}/generated/${GENERATOR}/go/routers.go
	rm -rf ./${VERSION}/generated/${GENERATOR}/go/README.md
	rm -rf ./${VERSION}/generated/${GENERATOR}/Dockerfile
	rm -rf ./${VERSION}/generated/${GENERATOR}/main.go
	rm -rf ./${VERSION}/generated/${GENERATOR}/.openapi-generator-ignore