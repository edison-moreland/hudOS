package components

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

type fill struct {
	col color.NRGBA
}

func Fill(c color.NRGBA) fill {
	return fill{
		col: c,
	}
}

func (f fill) Layout(gtx layout.Context, sz image.Point) layout.Dimensions {
	defer clip.Rect(image.Rectangle{Max: sz}).Push(gtx.Ops).Pop()
	paint.ColorOp{Color: f.col}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return layout.Dimensions{Size: sz}
}
