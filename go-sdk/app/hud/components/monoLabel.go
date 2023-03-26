package components

import (
	"gioui.org/unit"
	"gioui.org/widget/material"
)

func MonoLabel(th *material.Theme, size unit.Sp, text string) material.LabelStyle {
	label := material.Label(th, size, text)
	label.Font.Variant = "Mono"
	label.Color = th.Fg

	return label
}
