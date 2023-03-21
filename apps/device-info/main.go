package main

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	giosystem "gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
	"github.com/edison-moreland/nreal-hud/apps/device-info/components"
	"github.com/edison-moreland/nreal-hud/go-sdk/system/dbus"
)

func main() {
	go func() {
		w := app.NewWindow(
			app.Decorated(false),
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

	networkInfo := components.NetworkInfo(theme)
	batteryDevice := systemDBus.UPowerDaemon().DisplayDevice()

	var ops op.Ops
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case giosystem.DestroyEvent:
			return e.Err
		case giosystem.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			// Color background
			components.Fill(color.NRGBA{
				R: 0,
				G: 0,
				B: 0,
				A: 255,
			}).Layout(gtx, gtx.Constraints.Max)

			layout.UniformInset(10).Layout(gtx,
				func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
								layout.Rigid(components.Clock(theme).Layout),
								layout.Flexed(1, networkInfo.Layout),
							)
						}),
						layout.Flexed(1, components.Battery(batteryDevice).Layout),
					)
				},
			)

			e.Frame(gtx.Ops)
		}
	}
}
