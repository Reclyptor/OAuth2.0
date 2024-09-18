package util

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io"
)

func FromJSON[T interface{}](ctx echo.Context, t T) error {
	bytes, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, &t)
	if err != nil {
		return err
	}
	return nil
}
