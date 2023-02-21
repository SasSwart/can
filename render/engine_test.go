package render_test

import (
	"github.com/sasswart/gin-in-a-can/config"
	"github.com/sasswart/gin-in-a-can/render"
	"testing"
)

var e *render.Engine

func resetTestRenderer() {
	cfg := config.Data{}
	err := cfg.Load()
	if err != nil {
		panic(err.Error())
	}
	e = render.Engine{}.New(render.GinRenderer{}, cfg)
}

func Test_Render_Render(t *testing.T) {
	resetTestRenderer()

	t.Errorf("TODO")
}
