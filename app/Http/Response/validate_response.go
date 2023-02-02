package response

import models "github.com/apriliantocecep/ordent-restapi-artikel/app/Models"

type ValidateResponse struct {
	User models.User `json:"user"`
}
