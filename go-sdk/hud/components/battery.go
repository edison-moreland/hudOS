package components

import (
	"fmt"
	"image"
	"image/color"
	"time"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"github.com/edison-moreland/nreal-hud/go-sdk/system/dbus"
)

var (
	batteryBorder = color.NRGBA{
		R: 130,
		G: 130,
		B: 130,
		A: 255,
	}
)

// Todo: Indicate when battey is charging

type battery struct {
	displayDevice *dbus.UPowerDevice
}

func Battery(device *dbus.UPowerDevice) *battery {
	return &battery{
		displayDevice: device,
	}
}

func (b *battery) Layout(gtx layout.Context) layout.Dimensions {
	batteryPercentage, err := b.displayDevice.GetPercentage()
	if err != nil {
		// TODO: logging?
		fmt.Printf("Could not get battery percentage! %s\n", err.Error())
		batteryPercentage = 0
	}

	// Update battery 4 times per minute
	op.InvalidateOp{
		At: gtx.Now.Add(time.Minute / 4),
	}.Add(gtx.Ops)

	return layout.UniformInset(20).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return drawBattery(gtx.Ops, gtx.Constraints.Min, float64(batteryPercentage)/100)
	})
}

func drawBattery(ops *op.Ops, size image.Point, chargeLevel float64) layout.Dimensions {
	// Charge level should from 0 to 1
	// Green when full, red when empty
	fillColor := color.NRGBA{
		R: uint8(255.0 * (1 + -chargeLevel)),
		G: uint8(255.0 * chargeLevel),
		A: 255,
	}

	outlineRrect := clip.UniformRRect(
		image.Rect(0, 0, size.X, size.Y), 10,
	)

	fillRrect := outlineRrect
	fillRrect.Rect.Min.Y = size.Y - int(float64(size.Y)*chargeLevel)

	paint.FillShape(ops, fillColor, fillRrect.Op(ops))

	paint.FillShape(ops, batteryBorder,
		clip.Stroke{
			Path:  outlineRrect.Path(ops),
			Width: 4,
		}.Op(),
	)

	return layout.Dimensions{Size: size}
}
