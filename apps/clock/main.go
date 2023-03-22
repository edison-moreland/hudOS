package main

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
	"github.com/edison-moreland/nreal-hud/go-sdk/components"
	"github.com/edison-moreland/nreal-hud/go-sdk/system/dbus"
)

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

	systemDBus, err := dbus.NewSystemDbus()
	if err != nil {
		return err
	}
	defer systemDBus.Close()

	batteryDevice := systemDBus.UPowerDaemon().DisplayDevice()

	var ops op.Ops
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			components.Fill(color.NRGBA{
				R: 0,
				G: 0,
				B: 0,
				A: 255,
			}).Layout(gtx, gtx.Constraints.Max)

			layout.Flex{}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Center.Layout(gtx, components.Clock(theme, 96).Layout)
				}),
				layout.Flexed(1, components.Battery(batteryDevice).Layout),
			)

			e.Frame(gtx.Ops)
		}
	}
}
