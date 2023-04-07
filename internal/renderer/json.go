package renderer

import (
	"encoding/json"

	"github.com/vehsamrak/capital/internal/app"
)

type Json struct {
}

func (c Json) Render(capitalResult *app.CapitalResult) (render string, err error) {
	capitalResultJson, err := json.Marshal(capitalResult)
	if err != nil {
		return
	}

	return string(capitalResultJson), nil
}
