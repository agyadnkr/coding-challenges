package helpers

import "github.com/labstack/echo/v4"

func ReturnLog(c echo.Context, typelog int, data any, logmessage ...string) error {

	var logs map[string]any
	var response string

	if len(logmessage) > 0 {
		response = logmessage[0]
	} else {
		switch typelog {
		case 200:
			response = "OK"
		case 201:
			response = "CREATED"
		case 203:
			response = "FORBIDDEN"
		case 400:
			response = "BAD_REQUEST"
		case 404:
			response = "DATA_NOT_FOUND"
		case 500:
			response = "INTERNAL_SERVER_ERROR"
		}
	}

	logs = map[string]any{
		"status_sode": typelog,
		"Message":     response,
		"data":        data,
	}

	return c.JSON(typelog, logs)
}
