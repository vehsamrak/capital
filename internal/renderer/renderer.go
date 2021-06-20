package renderer

import "github.com/vehsamrak/capital/internal/app"

type Renderer interface {
	Render(result *app.CapitalResult)
}
