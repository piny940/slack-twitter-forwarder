package handler

import (
	"encoding/json"
	"os"

	"github.com/labstack/echo/v4"
)

type slackHandler struct{}

func NewSlackHandler() *slackHandler {
	return &slackHandler{}
}

func (h *slackHandler) Create(ctx echo.Context) error {
	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(ctx.Request().Body).Decode(&jsonBody)
	if err != nil {
		return ctx.JSON(400, echo.Map{
			"message": "Body must be a valid JSON object",
		})
	}
	if jsonBody["type"] == "url_verification" {
		return ctx.JSON(200, echo.Map{
			"challenge": jsonBody["challenge"],
		})
	}
	if jsonBody["api_app_id"] != os.Getenv("SLACK_APP_ID") || jsonBody["token"] != os.Getenv("SLACK_VERIFICATION_TOKEN") {
		return ctx.JSON(401, echo.Map{
			"message": "Unauthorized",
		})
	}
	event, ok := jsonBody["event"].(map[string]interface{})
	if !ok {
		return ctx.JSON(400, echo.Map{
			"message": "Event must be a valid JSON object",
		})
	}
	if event["type"] != "message" {
		return ctx.JSON(200, echo.Map{
			"message": "Ok",
		})
	}

	return ctx.JSON(200, echo.Map{
		"message": "Ok",
	})
}
