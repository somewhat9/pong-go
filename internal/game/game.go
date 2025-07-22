package game

import "github.com/somewhat9/pong-go/internal/config"

type Game struct {
	Cfg *config.Config
	p1 paddle
	p2 paddle
	ball
}

type paddle struct {
	score int
	y float32
	dy float32
}

type ball struct {
	x float32
	y float32
	dx float32
	dy float32
}
