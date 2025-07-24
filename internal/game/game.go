package game

import (
	"math"
	"math/rand"

	"github.com/somewhat9/pong-go/internal/config"
	"golang.org/x/image/font"
)

type Game struct {
	Cfg *config.Config
	Font font.Face
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
	speed float32
	speedUp float32
	maxBounceAngle float32
	v_max float32
	r float32
}

type midline struct {
	w float32
}

func (g *Game) collide() {
	if g.ball.y < g.ball.r || g.ball.y > float32(g.Cfg.Screen.Height)-g.ball.r {
		g.ball.dy *= -1
	}
	if g.ball.x < g.ball.r {
		g.p2.score++
		g.resetBall()
	}
	if g.ball.x > float32(g.Cfg.Screen.Width)-g.ball.r {
		g.p1.score++
		g.resetBall()
	}
	if g.p1.collision(&g.ball) {
		g.ball.bounceOffPaddle(&g.p1)
	} else if g.p2.collision(&g.ball) {
		g.ball.bounceOffPaddle(&g.p2)
	}
}

func (g *Game) resetBall() {
	g.ball.x = float32(g.Cfg.Screen.Width) / 2
	g.ball.y = float32(g.Cfg.Screen.Height) / 2
	
	angle := (rand.Float64()*2 - 1) * (math.Pi/4)
	if rand.Intn(2) == 0 {
		angle = math.Pi - angle
	}

	g.ball.speed = (rand.Float32()*5) + 3

	g.ball.dx = g.ball.speed * float32(math.Cos(angle))
	g.ball.dy = g.ball.speed * float32(math.Sin(angle))
}

func (b *ball) bounceOffPaddle(p *paddle)  {
	b.dx *= -1

	if b.dx > 0 {
		b.x = p.x+p.w+b.r
	} else {
		b.x = p.x-b.r
	}

	rel_y := (b.y - (p.y + p.h/2)) / (p.h/2)
	if rel_y < -1 { rel_y = -1 }
	if rel_y > 1 { rel_y = 1 }

	angle := float64(rel_y * b.maxBounceAngle)
	var dir float32 = 1.0
	if b.dx < 0 {dir = -1}
	b.dx = b.speed * float32(math.Cos(angle)) * dir
	b.dy = b.speed * float32(math.Sin(angle))

	b.speed += b.speedUp

	norm := b.speed / float32(math.Hypot(float64(b.dx), float64(b.dy)))
	b.dx *= norm
	b.dy *= norm

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

func (p *paddle) Automatic(b *ball) {
	target_y := b.y - p.h/2
	p.dy = target_y - p.y

	if float32(math.Abs(float64(p.dy))) > b.speed {
		p.dy = float32(math.Copysign(float64(b.speed)/2, float64(p.dy)))
	}
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
	g.ball.maxBounceAngle = 5*math.Pi/12
	g.speedUp = 1
	g.resetBall()
	return g
}
