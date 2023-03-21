module github.com/edison-moreland/nreal-hud

go 1.20

require (
	gioui.org v0.0.0-20230224004350-5f818bc5e7f9
	github.com/go-ble/ble v0.0.0-20230130210458-dd4b07d15402
	github.com/godbus/dbus v4.1.0+incompatible
	github.com/pkg/errors v0.9.1
)

require (
	gioui.org/cpu v0.0.0-20210817075930-8d6a761490d2 // indirect
	gioui.org/shader v1.0.6 // indirect
	github.com/JuulLabs-OSS/cbgo v0.0.1 // indirect
	github.com/benoitkugler/textlayout v0.3.0 // indirect
	github.com/go-text/typesetting v0.0.0-20221214153724-0399769901d5 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/mgutz/ansi v0.0.0-20170206155736-9520e82c474b // indirect
	github.com/mgutz/logxi v0.0.0-20161027140823-aebf8a7d67ab // indirect
	github.com/raff/goble v0.0.0-20190909174656-72afc67d6a99 // indirect
	github.com/sirupsen/logrus v1.5.0 // indirect
	golang.org/x/exp v0.0.0-20221012211006-4de253d81b95 // indirect
	golang.org/x/exp/shiny v0.0.0-20220827204233-334a2380cb91 // indirect
	golang.org/x/image v0.0.0-20220722155232-062f8c9fd539 // indirect
	golang.org/x/sys v0.0.0-20220825204002-c680a09ffe64 // indirect
	golang.org/x/text v0.7.0 // indirect
)

replace gioui.org => ./.build/vendor/go/gio
