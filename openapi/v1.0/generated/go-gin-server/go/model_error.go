/*
 * toriton Rest Api Specification
 *
 * LookBackDiaryのBackend API (Core)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Error - エラー
type Error struct {

	// エラー内容
	Type string `json:"type,omitempty"`

	Message string `json:"message,omitempty"`
}
