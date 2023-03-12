package main

import (
	"image"
	"image/color"
	"log"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget/material"
)

type fill struct {
	col color.NRGBA
}

func (f fill) Layout(gtx layout.Context, sz image.Point) layout.Dimensions {
	defer clip.Rect(image.Rectangle{Max: sz}).Push(gtx.Ops).Pop()
	paint.ColorOp{Color: f.col}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return layout.Dimensions{Size: sz}
}

type clock struct {
	material.LabelStyle
	background color.NRGBA
}

func Clock(th *material.Theme) clock {
	label := clock{
		LabelStyle: material.Label(
			th,
			th.TextSize*6,
			"12:00 PM",
		),
		background: th.Bg,
	}
	label.Font.Variant = "Mono"
	label.Color = color.NRGBA{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	}

	return label
}

func (cs *clock) Layout(gtx layout.Context) layout.Dimensions {
	cs.LabelStyle.Text = gtx.Now.Format("15:04 PM")

	op.InvalidateOp{
		At: gtx.Now.Add((time.Second * 60) - (time.Second * time.Duration(gtx.Now.Second()))),
	}.Add(gtx.Ops)

	fill{color.NRGBA{
		R: 0,
		G: 0,
		B: 0,
		A: 255,
	}}.Layout(gtx, gtx.Constraints.Max)

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
	{
		fg := theme.Fg
		bg := theme.Bg
		theme.Fg = fg
		theme.Bg = bg
	}

	clock := Clock(theme)

	var ops op.Ops
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			clock.Layout(gtx)

			e.Frame(gtx.Ops)
		}
	}
}
