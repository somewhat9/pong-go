package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 255, 0, 255})
}
