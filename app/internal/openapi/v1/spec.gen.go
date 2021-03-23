// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7xVW2sTTxT/KmH+/weFJZt6eclbW30o+CCtPkkJk92TZNrdmenMpBjKgrtrSx8UFapS",
	"nxShVsWKgtJipR9mmpZ+C5nZ3JpsaEXt017OnDPnd5kzK8hjIWcUqJKovIIESM6oBPsxhf1ZWGqCVObL",
	"Y1QBta+Y84B4WBFG3QXJqPknvQaE2Lz9L6CGyug/t1/azaLSvSkEEyiKIgf5ID1BuCmCykinz3T6Xqdv",
	"dLqvk4PD3cdHn96iyEEzVIGgOJgDsQwiy//33STfTB/pU53ut9dWdbJtm9tHZmkn2xSfFoAV3JUgBoga",
	"BrZlIX0/3vxx8ujr8cZ2d6mDuGAchCIZ3VXC6gLzRst8qBYHVEZSCULrhocqEarh4/wghJgEuRGKQ8gN",
	"NKWhNTcYOUjAUpMI8FH5Xn9lp1p3u3mnm8iqC+ApU3WQkMxJZ2lg1prMnrLDSnSpH6YrBClxPR9c9mNs",
	"rfbaantnzwChzdBAXMYBCXxroUoNkwB85CDScV5FWutVwHbYRz1A1wgPFtV4L7RXv5w8eKWTDzp9oZNd",
	"nbzTyWedrl+MJ2pEQBVLqBA/Nz7m91grccFqJIAKCeuVpgguzG4mk9AaGyX6FmOLU9hbvEGwaOl4x7wD",
	"9QuTt2cKl6aZgMvIQYqowJRTTBDFaGEWpCpMclKY4+CRWmekIActg5BZ3YliqVgycBgHijlBZXS1WCpO",
	"GN2waljJXNWZA3WwDyOnrTTjozKa5OROdvhPjdorpdJvTbWxRwHu45BnsLJtzjTr6PQzmb3jaydeMwyx",
	"aFkHr+lkT6frR5uJjj8e/jw43tjW8Usdb+n4iY5f6/i5jh+afXFdGjkt2nlTxTXKup6dEBYDO8e8RM4Q",
	"gf0JgzLfgFRTzG/9tUthdKZHpy2qRBOiP9TvvA30RMi5ME/zFDnoWtZEXu1es+7ApR456Pp5UvKu4GFb",
	"DIvWld8KNR9FUfQrAAD//9a26T5rCAAA",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}