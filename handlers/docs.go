// Package classification Burgers API
//
// Documentation for Burgers API
//
//	Schemes: http
// 	BasePath: /
// 	Version: 0.0.1
//
// 	Consumes:
//	- application/json
//
// 	Produces:
// 	- application/json
//
// swagger:meta
package handlers

import "github.com/haroundjudzman/golang-microservice/data"

// List of all burgers
// swagger:response burgersResponse
type burgersResponseWrapper struct {
	// All burgers
	// in: body
	Body []data.Burger
}

// No content is returned
// swagger:response noContentResponse
type noContentResponseWrapper struct{}

// Bad request path
// swagger:response badRequestResponse
type badRequestResponseWrapper struct{}

// No matching burger is found
// swagger:response notFoundResponse
type notFoundResponseWrapper struct{}

// Validation errors
// swagger:response validationErrorResponse
type validationErrorWrapper struct {
	// Collection of validation errors
	// in: body
	Body ValidationError
}

// Generic error
// swagger:response genericErrorResponse
type genericErrorWrapper struct {
	// Description of error
	// in: body
	Body GenericError
}
