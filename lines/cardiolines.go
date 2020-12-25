package main

import (
    "image/color"
    "math/rand"
    "github.com/fogleman/gg"
)

func main() {

    dc := gg.NewContext(1000, 600)
    dc.SetColor(color.Black)
    dc.Clear()

    dc.SetRGB(0.1, 0.2, 0.7)
    dc.SetLineWidth(10)
    dc.DrawLine(10, 100, 1000, 100)
    dc.Stroke()

    initx := 10.0
    inity := 100.0

    for x := 10; x <= 1000; x += 10 {
        dc.SetRGB(0.1, 0.5, 0.1)
        dc.SetLineWidth(5)
        y := rand.Float64() * 200
        dc.DrawLine(float64(x), y, initx, inity)
        dc.Stroke()
        initx = float64(x)
        inity = y
    }

    dc.SavePNG("cardiolines.png")
}
