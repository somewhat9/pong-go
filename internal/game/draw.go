package game

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.Cfg.Screen.Color)

	vector.DrawFilledRect(screen, g.p1.w*2, g.p1.y, g.p1.w, g.p1.h, g.Cfg.Paddle.Color, false)
	vector.DrawFilledRect(screen, float32(g.Cfg.Screen.Width) - g.p2.w*3, g.p2.y, g.p2.w, g.p2.h, g.Cfg.Paddle.Color, false)

	drawDashedLine(screen, float32(g.Cfg.Screen.Width)/2, 0, float32(g.Cfg.Screen.Width)/2, float32(g.Cfg.Screen.Height), float32(8.1), float32(16.75), g.midline.w, g.Cfg.Line.Color, false)
	
	b1 := text.BoundString(g.Font, fmt.Sprint(g.p1.score))
	b2 := text.BoundString(g.Font, fmt.Sprint(g.p2.score))
	text.Draw(screen, fmt.Sprint(g.p1.score), g.Font,  (g.Cfg.Screen.Width-b1.Dx())/4, 50, color.White)
	text.Draw(screen, fmt.Sprint(g.p2.score), g.Font, (g.Cfg.Screen.Width-b2.Dx())*3/4, 50, color.White)

	vector.DrawFilledCircle(screen, g.ball.x, g.ball.y, g.ball.r, g.Cfg.Ball.Color, false)
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
