package main

import (
	"image"
	"image/color"

	"gioui.org/app"
	"gioui.org/op/clip"
	"gioui.org/op/paint"

	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
)

// 64 x 32 multidimensional array representing the screen
var screen [32][64]int

// function for GUI down the line
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
			white := color.NRGBA{R: 255, G: 255, B: 255, A: 255}

			paint.Fill(&ops, black)

			for x, row := range screen {
				for y, col := range row {
					if col > 0 {
						print("▅▅")
						off := op.Offset(image.Pt(y*10, x*10)).Push(&ops)
						cl := clip.Rect{Max: image.Pt(10, 10)}.Push(&ops)
						paint.ColorOp{Color: white}.Add(&ops)
						paint.PaintOp{}.Add(&ops)
						off.Pop()
						cl.Pop()
					} else {
						print("  ")
					}
				}
				print("\n")
			}

			// defer clip.Rect{Max: image.Pt(10, 10)}.Push(&ops).Pop()
			// paint.ColorOp{Color: white}.Add(&ops)
			// paint.PaintOp{}.Add(&ops)

			// title := material.H1(th, "Hello, Gio")
			// maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
			// title.Color = maroon
			// title.Alignment = text.Middle
			// title.Layout(gtx)

			e.Frame(gtx.Ops)
		}
	}
}

// func run(w *app.Window) error {
// 	th := material.NewTheme(gofont.Collection())
// 	var ops op.Ops
// 	for {
// 		e := <-w.Events()
// 		switch e := e.(type) {
// 		case system.DestroyEvent:
// 			return e.Err
// 		case system.FrameEvent:
// 			gtx := layout.NewContext(&ops, e)

// 			title := material.H1(th, "Hello, Gio")
// 			maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
// 			title.Color = maroon
// 			title.Alignment = text.Middle
// 			title.Layout(gtx)

// 			e.Frame(gtx.Ops)
// 		}
// 	}
// }
