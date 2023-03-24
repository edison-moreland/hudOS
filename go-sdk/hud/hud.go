package hud

import (
	"image/color"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/widget/material"
	"github.com/rs/zerolog"
)

type HudMain func(window *app.Window, theme *material.Theme, logger zerolog.Logger) error

func App(appName string, appMain HudMain) {
	go func() {
		window := app.NewWindow(
			app.Decorated(false),
			app.Title(appName),
		)

		theme := material.NewTheme(gofont.Collection()).WithPalette(material.Palette{
			Fg: color.NRGBA{R: 255, G: 255, B: 255, A: 255},
			Bg: color.NRGBA{R: 0, G: 0, B: 0, A: 255},
		})

		logger := zerolog.New(zerolog.NewConsoleWriter())

		if err := appMain(window, &theme, logger); err != nil {
			logger.
				Fatal().
				Err(err).
				Msg("App main returned an error")
		}
	}()
	app.Main()
}
