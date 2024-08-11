package routers

import (
	"github.com/phongnd2802/go-ecommerce/internal/routers/manage"
	"github.com/phongnd2802/go-ecommerce/internal/routers/user"
)

type RouterGroup struct {
	User   user.UserRouterGroup
	Manage manage.ManageRouterGroup
}


var RouterApp = new(RouterGroup)
