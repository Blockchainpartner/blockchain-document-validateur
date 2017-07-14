package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetreceiptHandlerFunc turns a function with the right signature into a getreceipt handler
type GetreceiptHandlerFunc func(GetreceiptParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetreceiptHandlerFunc) Handle(params GetreceiptParams) middleware.Responder {
	return fn(params)
}

// GetreceiptHandler interface for that can handle valid getreceipt params
type GetreceiptHandler interface {
	Handle(GetreceiptParams) middleware.Responder
}

// NewGetreceipt creates a new http.Handler for the getreceipt operation
func NewGetreceipt(ctx *middleware.Context, handler GetreceiptHandler) *Getreceipt {
	return &Getreceipt{Context: ctx, Handler: handler}
}

/*Getreceipt swagger:route GET /recu getreceipt

Retourne le fichier avec le reçu

Retourne le fichier avec le reçu associé au hash fourni


*/
type Getreceipt struct {
	Context *middleware.Context
	Handler GetreceiptHandler
}

func (o *Getreceipt) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetreceiptParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
