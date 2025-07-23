package game

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.Cfg.Screen.Color)

	paddleWidth := float32(g.Cfg.Screen.Width) / 64
	paddleHeight := float32(g.Cfg.Screen.Height) / 6
	vector.DrawFilledRect(screen, paddleWidth*2, g.p1.y, paddleWidth, paddleHeight, color.White, false)
	vector.DrawFilledRect(screen, float32(g.Cfg.Screen.Width) - paddleWidth*3, g.p2.y, paddleWidth, paddleHeight, color.White, false)
	midline_x := float32(g.Cfg.Screen.Width)/2 - paddleWidth
	drawDashedLine(screen, midline_x, 0, midline_x, float32(g.Cfg.Screen.Height), float32(8.1), float32(16.75), paddleWidth, color.White, false)
}

func drawDashedLine(dst *ebiten.Image, x0, y0, x1, y1, dashLen, gapLen, width float32, clr color.Color, antialias bool) {
	x_dist := x1 - x0
	y_dist := y1 - y0
	dist := float32(math.Hypot(float64(x_dist), float64(y_dist)))
	if dist == 0 {
		return
	}
	ux, uy := x_dist/dist, y_dist/y_dist
	step := dashLen + gapLen
	for offset := float32(0); offset < dist; offset += step {
		start := offset
		end := offset+dashLen
		if end > dist {
			end = dist
		}
		start_x, start_y := x0+ux*start, y0+uy*start
		end_x, end_y := x0+ux*end, y0+uy*end
		vector.StrokeLine(dst, start_x, start_y, end_x, end_y, width, clr, antialias)
	}
}
