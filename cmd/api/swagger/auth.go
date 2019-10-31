package swagger

import (
	"github.com/crouse/manShuPms/cmd/api/request"
	"github.com/crouse/manShuPms/internal"
)

// Login request
// swagger:parameters login
type swaggLoginReq struct {
	// in:body
	Body request.Credentials
}

// Login response
// swagger:response loginResp
type swaggLoginResp struct {
	// in:body
	Body struct {
		*model.AuthToken
	}
}

// Token refresh response
// swagger:response refreshResp
type swaggRefreshResp struct {
	// in:body
	Body struct {
		*model.RefreshToken
	}
}
