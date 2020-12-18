package main

import (
    "image/color"
    "math/rand"
    "github.com/fogleman/gg"
)

func main() {

    dc := gg.NewContext(1400, 800)
    dc.SetColor(color.Black)
    dc.Clear()

    for x := 0; x < 1400; x += 40 {
        for y := 0; y < 800; y += 40 {
            dc.DrawRectangle(float64(x), float64(y), 25, 25)
            dc.SetRGB(rand.Float64() * 10, rand.Float64(), rand.Float64())
            dc.Fill()
        }
    }

    dc.SavePNG("xy-rectangle.png")
}
