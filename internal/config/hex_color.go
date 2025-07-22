package config

import (
	"encoding/hex"
	"fmt"
	"image/color"
	"strings"
)

type HexColor color.RGBA

func (h HexColor) RGBA() (r, g, b, a uint32) {
	r = uint32(h.R) * 0x101
	g = uint32(h.G) * 0x101
	b = uint32(h.B) * 0x101
	a = uint32(h.A) * 0x101
	return
}

func (h *HexColor) UnmarshalText(text []byte) error {
	clr, err := parseHexColor(string(text))
	if err != nil {
		return err
	}
	*h = HexColor{R: clr.R, G: clr.G, B: clr.B, A: clr.A}
	return nil
}

func parseHexColor(s string) (color.RGBA, error) {
	s = strings.TrimPrefix(s, "#")
	var c color.RGBA
	c.A = 0xFF

	if len(s) == 3 {
		s = fmt.Sprintf("%c[1]%c[1]%c[2]%c[2]%c[3]%c[3]", s[0], s[1], s[2])
	}

	b, err := hex.DecodeString(s)
	if err != nil {
		return c, fmt.Errorf("invalid hex data: %w", err)
	}
	
	switch len(b) {
	case 3:
		c.R, c.G, c.B = b[0], b[1], b[2]
	case 4:
		c.R, c.G, c.B, c.A = b[0], b[1], b[2], b[3]
	default:
		return c, fmt.Errorf("invalid hex length: got %d bytes, want 3 or 4 bytes", len(b))
	}

	return c, nil
}
