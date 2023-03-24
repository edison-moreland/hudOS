package main

import (
	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
	"github.com/rs/zerolog"

	"github.com/edison-moreland/nreal-hud/go-sdk/hud"
	"github.com/edison-moreland/nreal-hud/go-sdk/hud/components"
	"github.com/edison-moreland/nreal-hud/go-sdk/system/dbus"
)

func main() {
	hud.App("clock", hudMain)
}

func hudMain(window *app.Window, theme *material.Theme, logger zerolog.Logger) error {
	systemDBus, err := dbus.NewSystemDbus()
	if err != nil {
		return err
	}
	defer systemDBus.Close()

	batteryDevice := systemDBus.UPowerDaemon().DisplayDevice()

	var ops op.Ops
	for {
		event := <-window.Events()
		switch event := event.(type) {
		case system.DestroyEvent:
			return event.Err

		case system.FrameEvent:
			gtx := layout.NewContext(&ops, event)
			components.Fill(theme.Bg).
				Layout(gtx, gtx.Constraints.Max)

			layout.Flex{}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Center.Layout(gtx, components.Clock(theme, 96).Layout)
				}),
				layout.Flexed(1, components.Battery(batteryDevice).Layout),
			)

			event.Frame(gtx.Ops)
		}
	}
}
