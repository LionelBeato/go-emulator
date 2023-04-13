package main

import (
	"image"
	"image/color"

	"gioui.org/app"
	// "gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	// "gioui.org/text"
	// "gioui.org/widget/material"
)

func run(w *app.Window) error {
	// th := material.NewTheme(gofont.Collection())
	var ops op.Ops
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			black := color.NRGBA{R: 0, G: 0, B: 0, A: 255}

			defer clip.Rect{Max: image.Pt(10, 10)}.Push(&ops).Pop()
			paint.ColorOp{Color: black}.Add(&ops)
			paint.PaintOp{}.Add(&ops)

			// title := material.H1(th, "Hello, Gio")
			// title.Color = maroon
			// title.Alignment = text.Middle
			// title.Layout(gtx)

			e.Frame(gtx.Ops)
		}
	}
}