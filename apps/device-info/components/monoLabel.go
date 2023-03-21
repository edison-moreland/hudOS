package components

import (
	"image/color"

	"gioui.org/unit"
	"gioui.org/widget/material"
)

func MonoLabel(th *material.Theme, size unit.Sp, text string) material.LabelStyle {
	label := material.Label(th, size, text)
	label.Font.Variant = "Mono"
	label.Color = color.NRGBA{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	}

	return label
}
