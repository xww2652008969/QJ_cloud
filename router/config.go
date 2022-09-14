package router

type Router struct {
	User UserRouter
}

var Routerapp = new(Router)
