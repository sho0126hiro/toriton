/*
 * toriton Rest Api Specification
 *
 * LookBackDiaryのBackend API (Core)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CreateUserRequest - ユーザ登録用Request
type CreateUserRequest struct {

	Username string `json:"username"`

	Name string `json:"name"`

	Biography string `json:"biography,omitempty"`

	Birthday string `json:"birthday,omitempty"`

	Email string `json:"email"`
}
