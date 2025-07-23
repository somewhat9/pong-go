package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Update() error {
	
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyArrowUp) && g.p1.dy > -g.p1.v_max:
		g.p1.dy -= 1
	case ebiten.IsKeyPressed(ebiten.KeyArrowDown) && g.p1.dy < g.p1.v_max:
		g.p1.dy += 1
	}

	if g.p1.dy < 0 {
		g.p1.dy += 0.25
	} else if g.p1.dy > 0 {
		g.p1.dy -= 0.25
	}

	if (g.p1.y + g.p1.dy < 0 || g.p1.y + g.p1.dy > float32(g.Cfg.Screen.Height)-g.p1.h) {
		g.p1.dy = -g.p1.dy
	} 
	g.p1.y += g.p1.dy
	
	g.ball.x += g.ball.dx
	g.ball.y += g.ball.dy
	g.collide()

	return nil
}
