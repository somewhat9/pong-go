package config

type Config struct {
	Window Window     `yaml:"window"`
	Screen Screen     `yaml:"screen"`
	Paddle Paddle     `yaml:"paddle"`
	Ball Ball         `yaml:"ball"`
	Line Line         `yaml:"line"`
}

type Screen struct {
	Width int         `yaml:"width"`
	Height int        `yaml:"height"`
	Color HexColor    `yaml:"color"`
}

type Window struct {
	Title string      `yaml:"title"`
	Width int         `yaml:"width"`
	Height int        `yaml:"height"`
	Resizable bool    `yaml:"resizable"`
}

type Paddle struct {
	Color HexColor    `yaml:"color"`
}

type Ball struct {
	Color HexColor    `yaml:"color"`
}

type Line struct {
	Color HexColor    `yaml:"color"`
} 

func Default() *Config {
	return &Config {
		Window{Title: "Pong Go", Width: 800, Height: 600},
		Screen{Width: 640, Height: 480, Color: HexColor{5, 5, 5, 255}},
		Paddle{Color: HexColor{255, 255, 255, 255}},
		Ball{Color: HexColor{255, 255, 255, 255}},
		Line{Color: HexColor{175, 175, 175, 175}},
	}
}
