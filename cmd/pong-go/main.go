package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/somewhat9/pong-go/internal/assets"
	"github.com/somewhat9/pong-go/internal/config"
	"github.com/somewhat9/pong-go/internal/game"
)

func main() {
	cfg, err := config.LoadYAML("settings")
	g := game.NewGame(cfg)
	if err != nil {
		log.Fatalf("could not load .yaml: %v", err)
	}

	if g.Cfg.Window.Resizable {
		ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	}
	ebiten.SetWindowSize(g.Cfg.Window.Width, g.Cfg.Window.Height)
	ebiten.SetWindowTitle(g.Cfg.Window.Title)
	g.Font = assets.LoadFont(g.Cfg.Screen.FontSize)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
