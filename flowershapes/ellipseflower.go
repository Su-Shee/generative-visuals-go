package main

import (
    "image/color"
    "math/rand"
    "github.com/fogleman/gg"
)

func main() {
    const S = 1024
    dc := gg.NewContext(S, 980)
    dc.SetColor(color.Black)
    dc.Clear()

    for i := 0; i < 360; i += 15 {
        dc.Push()
        dc.RotateAbout(gg.Radians(float64(i)), S/2, S/2)
        dc.DrawEllipse(S/2, S/2, S*7/16, S/8)
        dc.SetRGBA(rand.Float64(), rand.Float64(), rand.Float64(), 0.2)
        //dc.SetRGBA(0.7, float64(i + 20), float64(i), 0.5)
        //dc.SetRGBA(float64(i), float64(i), float64(i), 0.5)
        dc.Fill()
        dc.Pop()
    }
    dc.SavePNG("flower.png")
}
