package main

import (
	"net"
	"strings"

	"gioui.org/layout"
	"gioui.org/widget/material"

	"github.com/edison-moreland/nreal-hud/go-sdk/app/hud/components"
)

// Information to display
// Ip address
// WiFi network

type networkInfo struct {
	layout.List
	theme *material.Theme
}

func NetworkInfo(th *material.Theme) networkInfo {
	return networkInfo{
		List: layout.List{
			Axis: layout.Vertical,
		},
		theme: th,
	}
}

func (n *networkInfo) Layout(gtx layout.Context) layout.Dimensions {
	ifaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	return n.List.Layout(gtx, len(ifaces), func(gtx layout.Context, i int) layout.Dimensions {
		iface := ifaces[i]

		var sb strings.Builder
		sb.WriteString(iface.Name)
		addrs, err := iface.Addrs()
		if err != nil {
			panic(err)
		}

		for _, addr := range addrs {
			ip, _, err := net.ParseCIDR(addr.String())
			if err != nil {
				panic(err)
			}

			if ip.To4() == nil {
				continue
			}

			sb.WriteString(" ")
			sb.WriteString(ip.To4().String())
		}

		return components.MonoLabel(n.theme, 40, sb.String()).Layout(gtx)
	})
}
