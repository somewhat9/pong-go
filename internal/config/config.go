package config

type Config struct {
	Window Window     `yaml:"window"`
	Screen Screen     `yaml:"screen"`
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

func Default() *Config {
	return &Config {
		Window{Title: "Pong Go", Width: 800, Height: 600},
		Screen{Width: 640, Height: 480, Color: HexColor{5, 5, 5, 255}},
	}
}
