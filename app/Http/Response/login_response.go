package response

import models "github.com/apriliantocecep/ordent-restapi-artikel/app/Models"

type LoginResponse struct {
	User  models.User `json:"user"`
	Token string      `json:"token"`
}
