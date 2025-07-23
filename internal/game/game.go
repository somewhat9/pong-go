package game

import (
	"math/rand"

	"github.com/somewhat9/pong-go/internal/config"
)

type Game struct {
	Cfg *config.Config
	p1 paddle
	p2 paddle
	ball
	midline
}

type paddle struct {
	score int
	y float32
	dy float32
	v_max float32
	x float32
	w float32
	h float32
}

type ball struct {
	x float32
	y float32
	dx float32
	dy float32
	v_max float32
	r float32
}


func (g *Game) collide() {
	if g.ball.y < g.ball.r || g.ball.y > float32(g.Cfg.Screen.Height)-g.ball.r {
		g.ball.dy = -g.ball.dy
		g.ball.dx = -g.ball.dx
	}
	if g.ball.x < g.ball.r {
		g.p2.score++
		g.resetBall()
	}
	if g.ball.x > float32(g.Cfg.Screen.Width)-g.ball.r {
		g.p1.score++
		g.resetBall()
	}
	if g.p1.collision(&g.ball) || g.p2.collision(&g.ball) {
		g.ball.dy = -g.ball.dy
		g.ball.dx = -g.ball.dx
	}
}

func (g *Game) resetBall() {
	g.ball.x = float32(g.Cfg.Screen.Width) / 2
	g.ball.y = float32(g.Cfg.Screen.Height) / 2
	g.ball.dx = (rand.Float32()*g.ball.v_max)+1
	g.ball.dy = (rand.Float32()*g.ball.v_max)+1
}

func (p *paddle) collision(b *ball) bool {
	closest_x := clampToInterval(b.x, p.x, p.x+p.w)
	closest_y := clampToInterval(b.y, p.y, p.y+p.h)

	x_dist := b.x - closest_x
	y_dist := b.y - closest_y

	return x_dist*x_dist+y_dist*y_dist <= b.r*b.r
}

func clampToInterval(v, min, max float32) float32 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v

}


type midline struct {
	w float32
}

func NewGame(cfg *config.Config) *Game {
	g := &Game{}
	g.Cfg = cfg

	g.p1.w = float32(g.Cfg.Screen.Width) / 64
	g.p2.w = float32(g.Cfg.Screen.Width) / 64

	g.p1.h = float32(g.Cfg.Screen.Height) / 6
	g.p2.h = float32(g.Cfg.Screen.Height) / 6

	g.p1.x = g.p1.w*2
	g.p2.x = float32(g.Cfg.Screen.Width) - g.p2.w*3

	g.p1.y = (float32(g.Cfg.Screen.Height) - g.p1.h) / 2
	g.p2.y = (float32(g.Cfg.Screen.Height) - g.p2.h) / 2

	g.p1.v_max = float32(5)
	g.p2.v_max = float32(5)
	g.ball.v_max = float32(5)

	g.midline.w = float32(g.Cfg.Screen.Width) / 64

	g.ball.r = float32(g.Cfg.Screen.Width) / 64
	g.resetBall()
	
	return g
}
