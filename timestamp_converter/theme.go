package timestamp_converter

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

var (
	orange = &color.NRGBA{R: 255, G: 125, B: 30, A: 255}
)

type myTheme struct {
	variant string
}

func lightPaletColorNamed(name fyne.ThemeColorName) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	case theme.ColorNameButton:
		return color.NRGBA{R: 0xf5, G: 0xf5, B: 0xf5, A: 0xff}
	case theme.ColorNameDisabled:
		return color.NRGBA{R: 0xe3, G: 0xe3, B: 0xe3, A: 0xff}
	case theme.ColorNameDisabledButton:
		return color.NRGBA{R: 0xf5, G: 0xf5, B: 0xf5, A: 0xff}
	case theme.ColorNameError:
		return color.NRGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0xff}
	case theme.ColorNameForeground:
		return color.NRGBA{R: 0x56, G: 0x56, B: 0x56, A: 0xff}
	case theme.ColorNameHover:
		return color.NRGBA{A: 0x0f}
	case theme.ColorNameInputBackground:
		return color.NRGBA{R: 0xf3, G: 0xf3, B: 0xf3, A: 0xff}
	case theme.ColorNameInputBorder:
		return color.NRGBA{R: 0xe3, G: 0xe3, B: 0xe3, A: 0xff}
	case theme.ColorNameMenuBackground:
		return color.NRGBA{R: 0xf5, G: 0xf5, B: 0xf5, A: 0xff}
	case theme.ColorNameOverlayBackground:
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	case theme.ColorNamePlaceHolder:
		return color.NRGBA{R: 0x88, G: 0x88, B: 0x88, A: 0xff}
	case theme.ColorNamePressed:
		return color.NRGBA{A: 0x19}
	case theme.ColorNameScrollBar:
		return color.NRGBA{A: 0x99}
	case theme.ColorNameSeparator:
		return color.NRGBA{R: 0xf5, G: 0xf5, B: 0xf5, A: 0xff}
	case theme.ColorNameShadow:
		return color.NRGBA{A: 0x33}
	case theme.ColorNameSuccess:
		return color.NRGBA{R: 0x43, G: 0xf4, B: 0x36, A: 0xff}
	case theme.ColorNameWarning:
		return color.NRGBA{R: 0xff, G: 0x98, B: 0x00, A: 0xff}
	}

	return color.Transparent
}

func darkPaletColorNamed(name fyne.ThemeColorName) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.NRGBA{R: 0x14, G: 0x14, B: 0x15, A: 0xff}
	case theme.ColorNameButton:
		return color.NRGBA{R: 0x28, G: 0x29, B: 0x2e, A: 0xff}
	case theme.ColorNameDisabled:
		return color.NRGBA{R: 0x39, G: 0x39, B: 0x3a, A: 0xff}
	case theme.ColorNameDisabledButton:
		return color.NRGBA{R: 0x28, G: 0x29, B: 0x2e, A: 0xff}
	case theme.ColorNameError:
		return color.NRGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0xff}
	case theme.ColorNameForeground:
		return color.NRGBA{R: 0xf3, G: 0xf3, B: 0xf3, A: 0xff}
	case theme.ColorNameHover:
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x0f}
	case theme.ColorNameInputBackground:
		return color.NRGBA{R: 0x20, G: 0x20, B: 0x23, A: 0xff}
	case theme.ColorNameInputBorder:
		return color.NRGBA{R: 0x39, G: 0x39, B: 0x3a, A: 0xff}
	case theme.ColorNameMenuBackground:
		return color.NRGBA{R: 0x28, G: 0x29, B: 0x2e, A: 0xff}
	case theme.ColorNameOverlayBackground:
		return color.NRGBA{R: 0x18, G: 0x1d, B: 0x25, A: 0xff}
	case theme.ColorNamePlaceHolder:
		return color.NRGBA{R: 0xb2, G: 0xb2, B: 0xb2, A: 0xff}
	case theme.ColorNamePressed:
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x66}
	case theme.ColorNameScrollBar:
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x99}
	case theme.ColorNameSeparator:
		return color.NRGBA{R: 0x0, G: 0x0, B: 0x0, A: 0xff}
	case theme.ColorNameShadow:
		return color.NRGBA{A: 0x66}
	case theme.ColorNameSuccess:
		return color.NRGBA{R: 0x43, G: 0xf4, B: 0x36, A: 0xff}
	case theme.ColorNameWarning:
		return color.NRGBA{R: 0xff, G: 0x98, B: 0x00, A: 0xff}
	}

	return color.Transparent
}

func (t *myTheme) Color(c fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	if c == theme.ColorNamePrimary {
		return orange
	}

	if t.variant == "light" {
		return lightPaletColorNamed(c)
	}

	return darkPaletColorNamed(c)
}

func (t *myTheme) Font(s fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(s)
}

func (t *myTheme) Icon(i fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(i)
}

func (t *myTheme) Size(s fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(s)
}
