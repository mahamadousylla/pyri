package baseControllers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"utilities/logger"
)

const fileName = "baseController.go"

type BaseController struct {
	rw  http.ResponseWriter
	req *http.Request
	Ctx context.Context
}

// NewBaseController ...
func NewBaseController(rw http.ResponseWriter, req *http.Request) BaseController {
	return BaseController{
		rw:  rw,
		req: req,
		Ctx: req.Context(),
	}
}

// ServiceResponse ...
func (c *BaseController) WriteResponse(statusCode int, response interface{}) {
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
func (c *BaseController) UnmarshallBody(dest interface{}) error {
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

func (c *BaseController) toJSON(statusCode int, response interface{}) {
	rw := c.rw
	headers := rw.Header()
	headers.Set("Content-Type", "application/json; charset=UTF-8")
	headers.Set("X-Content-Type-Options", "nosniff")

	rw.WriteHeader(statusCode)

	if err := json.NewEncoder(rw).Encode(response); err != nil {
		l := logger.GetLogger(c.Ctx, fileName, "json")
		l.Error("Error encoding response object")
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

// Sends only a status code to client
func (c *BaseController) emptyReponse(statusCode int) {
	rw := c.rw
	rw.WriteHeader(statusCode)
}

func (c *BaseController) unmarshallJSON(dest interface{}) error {
	if err := json.NewDecoder(c.req.Body).Decode(dest); err != nil {
		return err
	}

	defer c.req.Body.Close()

	return nil
}
