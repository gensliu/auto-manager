package response

import "auto-manager/model"

type Login struct {
	Manager   model.Manager `json:"manager"`
	Token     string        `json:"token"`
	ExpiresAt int64         `json:"expiresAt"`
}
