package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type myTheme struct{}

var _ fyne.Theme = (*myTheme)(nil)

func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(name, variant)
}

func (m myTheme) Font(style fyne.TextStyle) fyne.Resource {
	if style.Bold {
		return resourceMsyhbdTtc
	}
	if style.Italic {
		return theme.DefaultTheme().Font(style)
	}
	if style.Monospace {
		return theme.DefaultTheme().Font(style)
	}
	return resourceMsyhTtc
}

func (m myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(&myTheme{})
	myWindow := myApp.NewWindow("中文")

	l1 := widget.NewLabel("娃哈哈")
	l2 := widget.NewLabel("粗娃哈哈")
	l2.TextStyle.Bold = true
	container := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), l1, l2)
	myWindow.SetContent(container)

	myWindow.ShowAndRun()
}
