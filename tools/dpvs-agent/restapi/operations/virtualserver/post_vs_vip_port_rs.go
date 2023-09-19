// Code generated by go-swagger; DO NOT EDIT.

package virtualserver

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostVsVipPortRsHandlerFunc turns a function with the right signature into a post vs vip port rs handler
type PostVsVipPortRsHandlerFunc func(PostVsVipPortRsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostVsVipPortRsHandlerFunc) Handle(params PostVsVipPortRsParams) middleware.Responder {
	return fn(params)
}

// PostVsVipPortRsHandler interface for that can handle valid post vs vip port rs params
type PostVsVipPortRsHandler interface {
	Handle(PostVsVipPortRsParams) middleware.Responder
}

// NewPostVsVipPortRs creates a new http.Handler for the post vs vip port rs operation
func NewPostVsVipPortRs(ctx *middleware.Context, handler PostVsVipPortRsHandler) *PostVsVipPortRs {
	return &PostVsVipPortRs{Context: ctx, Handler: handler}
}

/*
	PostVsVipPortRs swagger:route POST /vs/{VipPort}/rs virtualserver postVsVipPortRs

Update fully real server list to vip:port:proto
*/
type PostVsVipPortRs struct {
	Context *middleware.Context
	Handler PostVsVipPortRsHandler
}

func (o *PostVsVipPortRs) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostVsVipPortRsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}