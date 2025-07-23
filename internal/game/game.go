package game

import "github.com/somewhat9/pong-go/internal/config"

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
	x float32
	w float32
	h float32
}

type ball struct {
	x float32
	y float32
	dx float32
	dy float32
	r float32
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

	g.midline.w = float32(g.Cfg.Screen.Width) / 64

	g.ball.r = float32(g.Cfg.Screen.Width) / 64
	g.ball.x = float32(g.Cfg.Screen.Width) / 2
	g.ball.y = float32(g.Cfg.Screen.Height) / 2
	
	return g
}
