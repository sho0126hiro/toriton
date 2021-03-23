PROJECT_DIR=${PWD}/..
VERSION=v1
OUT_DIR=${PROJECT_DIR}/app/internal/openapi/${VERSION}

# spec.ymlからソースコードを生成
oapi-gen:
	oapi-codegen -generate "spec" -package openapi ./${VERSION}/openapi.yml > ${OUT_DIR}/spec.gen.go
	oapi-codegen -generate "types" -package openapi ./${VERSION}/openapi.yml > ${OUT_DIR}/type.gen.go
	oapi-codegen -generate "server" -package openapi ./${VERSION}/openapi.yml > ${OUT_DIR}/server.gen.go