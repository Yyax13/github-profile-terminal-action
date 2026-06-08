package theme

import "image/color"

type Theme struct {
	Foreground color.RGBA
	Background color.RGBA
	Highlight  color.RGBA
}

var darkTheme = Theme{
	Foreground: color.RGBA{
		R: 0xad,
		G: 0xba,
		B: 0xc7,
		A: 0xff,
	},
	Background: color.RGBA{ // 0d1117
		R: 0x0d,
		G: 0x11,
		B: 0x17,
		A: 0xff,
	},
	Highlight: color.RGBA{
		R: 0x50,
		G: 0x94,
		B: 0xf0,
		A: 0xff,
	},
}

func ByName(name string) Theme {
	switch name {
	case "dark":
		return darkTheme
	case "light":
		panic("oh no")
	default:
		return darkTheme
	}
}
