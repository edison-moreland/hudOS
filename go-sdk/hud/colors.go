package hud

import "image/color"

// TODO: Replace these with more pleasing colors
var (
	Black = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	White = color.NRGBA{R: 255, G: 255, B: 255, A: 255}

	Red   = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	Green = color.NRGBA{R: 0, G: 255, B: 0, A: 255}
	Blue  = color.NRGBA{R: 0, G: 0, B: 255, A: 255}

	Yellow  = color.NRGBA{R: 255, G: 255, B: 0, A: 255}
	Cyan    = color.NRGBA{R: 0, G: 255, B: 255, A: 255}
	Magenta = color.NRGBA{R: 255, G: 0, B: 255, A: 255}
)
