package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const fileName = "controller.go"

// Controller ...
type Controller struct {
	rw  http.ResponseWriter
	req *http.Request
	Ctx context.Context
}

// New ...
func New(rw http.ResponseWriter, req *http.Request) Controller {
	return Controller{
		rw:  rw,
		req: req,
		Ctx: req.Context(),
	}
}

// WriteResponse ...
func (c *Controller) WriteResponse(statusCode int, response interface{}) {
	acceptedType := c.req.Header.Get("Accept")

	switch acceptedType {
	// For now we will only process application/json Accept headers
	case "*/*", "application/json", "":
		c.toJSON(statusCode, response)
		break
	default:
		c.emptyReponse(http.StatusNotAcceptable)
	}
}

// func (c *BaseController) ErrorResponse(statusCode int, message string) {
// 	http.Error(c.rw, message, statusCode)
// }

// UnmarshallBody ...
func (c *Controller) UnmarshallBody(dest interface{}) error {
	contentType := c.req.Header.Get("Content-Type")

	switch contentType {
	case "application/json":
		if err := c.unmarshallJSON(dest); err != nil {
			return err
		}
	default:
		msg := fmt.Sprintf("Content-Type %s is not supported", contentType)
		return errors.New(msg)
	}

	return nil
}

/*
 * JSON
 */

func (c *Controller) toJSON(statusCode int, response interface{}) {
	rw := c.rw
	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rw.Header().Set("X-Content-Type-Options", "nosniff")

	rw.WriteHeader(statusCode)

	if err := json.NewEncoder(rw).Encode(response); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

// Sends only a status code to client
func (c *Controller) emptyReponse(statusCode int) {
	rw := c.rw
	rw.WriteHeader(statusCode)
}

func (c *Controller) unmarshallJSON(dest interface{}) error {
	if err := json.NewDecoder(c.req.Body).Decode(dest); err != nil {
		return err
	}

	defer c.req.Body.Close()

	return nil
}
