package render_test

import (
	"github.com/sasswart/gin-in-a-can/render"
	"testing"
)

var Renderer *render.Engine

func resetTestRenderer() {
	Renderer = render.Engine{}.New(render.GinRenderer{}, render.Config{})
}

func Test_Render_Render(t *testing.T) {
	resetTestRenderer()

	t.Errorf("TODO")
}
