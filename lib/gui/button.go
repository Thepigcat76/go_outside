package gui

import (
	"embed"

	"go_outside/lib/util"

	"github.com/veandco/go-sdl2/sdl"
)

type Button struct {
	Name        string
	Clicked     bool
	Visible     bool
	texture_std *sdl.Texture
	texture_sel *sdl.Texture
	button_rect *sdl.Rect
}

func Create_button(name string, texturePath string, renderer *sdl.Renderer, assets embed.FS, X, Y, W, H int32) Button {
	texture_std := util.Load_image(texturePath, renderer, assets, 1)
	texture_sel := util.Load_image(texturePath + "_selected", renderer, assets, 1)
	button_rect := sdl.Rect{X: X, Y: Y, W: W, H: H}

	return Button{Name: name, Clicked: false, Visible: true, texture_std: texture_std.Texture, texture_sel: texture_sel.Texture, button_rect: &button_rect}
}

func (b *Button) Draw_button(renderer *sdl.Renderer) {
	mouse_x, mouse_y, _ := sdl.GetMouseState()
	left_click, _ := util.Mouse_clicked()
	if b.Visible == true {
		if util.Collide_point(b.button_rect, mouse_x, mouse_y) {
			renderer.Copy(b.texture_sel, nil, b.button_rect)
			if left_click {
				b.Clicked = true
			} else {
				b.Clicked = false
			}
		} else {
			renderer.Copy(b.texture_std, nil, b.button_rect)
		}
	}
}
