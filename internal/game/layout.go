package game

func (g *Game) Layout(outsideWidth, oustideHeight int) (screenWidth, screenHeight int) {
	return g.Cfg.Screen.Width, g.Cfg.Screen.Height
}
