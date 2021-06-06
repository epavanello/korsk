package transport

import (
	"github.com/epavanello/gorsk/pkg/models"
)

// User model response
// swagger:response userResp
type swaggUserResponse struct {
	// in:body
	Body struct {
		*models.User
	}
}

// Users model response
// swagger:response userListResp
type swaggUserListResponse struct {
	// in:body
	Body struct {
		Users []models.User `json:"users"`
		Page  int           `json:"page"`
	}
}
