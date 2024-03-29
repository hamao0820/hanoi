package hsl

import "image/color"

type HSL struct {
	H, S, L float64
}

func NewHSL(h, s, l float64) HSL {
	return HSL{h, s, l}
}

func hueToRGB(v1, v2, h float64) float64 {
	if h < 0 {
		h += 1
	}
	if h > 1 {
		h -= 1
	}
	switch {
	case 6*h < 1:
		return (v1 + (v2-v1)*6*h)
	case 2*h < 1:
		return v2
	case 3*h < 2:
		return v1 + (v2-v1)*((2.0/3.0)-h)*6
	}
	return v1
}

func (c HSL) ToRGB() color.Color {
	h := c.H
	s := c.S
	l := c.L

	if s == 0 {
		return color.RGBA{uint8(l * 255), uint8(l * 255), uint8(l * 255), 255}
	}

	var v1, v2 float64
	if l < 0.5 {
		v2 = l * (1 + s)
	} else {
		v2 = (l + s) - (s * l)
	}

	v1 = 2*l - v2

	r := hueToRGB(v1, v2, h+(1.0/3.0))
	g := hueToRGB(v1, v2, h)
	b := hueToRGB(v1, v2, h-(1.0/3.0))

	return color.RGBA{uint8(r * 255), uint8(g * 255), uint8(b * 255), 255}
}
