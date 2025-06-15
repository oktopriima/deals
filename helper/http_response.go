package helper

import "github.com/labstack/echo/v4"

type HttpResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// ResponseOK sends a JSON response with HTTP status 200 (OK).
// It wraps the provided data in a standard HttpResponse structure
// with a success code and message.
//
// Parameters:
//
//	c    - echo.Context: the Echo context for the current request.
//	data - interface{}: the payload to include in the response.
//
// Returns:
//
//	error: any error encountered while sending the response.
func ResponseOK(c echo.Context, data interface{}) error {
	return c.JSON(200, HttpResponse{
		Code:    200,
		Message: "OK",
		Data:    data,
	})
}

// ResponseFailed sends a JSON response with the specified HTTP status code and error message.
// The response body contains a standardized structure with the code, message, and a nil data field.
// It is typically used to return error responses in Echo handlers.
//
// Parameters:
//
//	c      - the Echo context used to send the response
//	code   - the HTTP status code to set in the response
//	message - the error message to include in the response body
//
// Returns:
//
//	An error if the response could not be sent, otherwise nil.
func ResponseFailed(c echo.Context, code int, message string) error {
	return c.JSON(code, HttpResponse{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

// ResponseCreated sends a JSON response with HTTP status 201 (Created).
// It wraps the provided data in a standardized HttpResponse structure with a "Created" message.
// Parameters:
//   - c: echo.Context, the Echo context for the current request.
//   - data: interface{}, the payload to include in the response.
//
// Returns an error if the response could not be sent.
func ResponseCreated(c echo.Context, data interface{}) error {
	return c.JSON(201, HttpResponse{
		Code:    201,
		Message: "Created",
		Data:    data,
	})
}
