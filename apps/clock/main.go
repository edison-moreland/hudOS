package main

import (
	"image"
	"log"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
)

type clock struct {
	material.LabelStyle
}

func Clock(th *material.Theme) clock {
	label := clock{
		LabelStyle: material.Label(
			th,
			th.TextSize*6,
			"12:00 PM",
		),
	}
	label.Font.Variant = "Mono"

	return label
}

func (cs *clock) Layout(gtx layout.Context) layout.Dimensions {
	cs.LabelStyle.Text = gtx.Now.Format("15:04 PM")

	op.InvalidateOp{
		At: gtx.Now.Add((time.Second * 60) - (time.Second * time.Duration(gtx.Now.Second()))),
	}.Add(gtx.Ops)

	return cs.LabelStyle.Layout(gtx)
}

func NoMinimum(gtx layout.Context, w layout.Widget) layout.Dimensions {
	gtx.Constraints.Min = image.Point{}
	return w(gtx)
}

func main() {
	go func() {
		w := app.NewWindow(
			app.Decorated(false),
			app.Size(461, 112),
		)
		err := run(w)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(w *app.Window) error {
	theme := material.NewTheme(gofont.Collection())
	clock := Clock(theme)

	var ops op.Ops
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			NoMinimum(gtx, clock.Layout)

			e.Frame(gtx.Ops)
		}
	}
}
