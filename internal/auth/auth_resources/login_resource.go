package auth_resources

import "simple-blog-api-golang/pkg/resource"

type LoginResource struct {
	resource.GeneralResource
	Token string
}

func NewLoginResource(generalResource resource.GeneralResource, token string) *LoginResource {
	return &LoginResource{GeneralResource: generalResource, Token: token}
}
