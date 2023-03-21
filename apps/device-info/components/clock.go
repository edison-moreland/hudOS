package components

import (
	"time"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
)

type clock struct {
	material.LabelStyle
}

func Clock(th *material.Theme) *clock {
	return &clock{
		LabelStyle: MonoLabel(th, 110, time.Now().Format("15:04 PM")),
	}
}

func (cs *clock) Layout(gtx layout.Context) layout.Dimensions {
	// Ask to be redrawn when the second ticks over
	op.InvalidateOp{
		At: gtx.Now.Add(time.Minute - (time.Second * time.Duration(gtx.Now.Second()))),
	}.Add(gtx.Ops)

	return cs.LabelStyle.Layout(gtx)
}
