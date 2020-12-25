package main

import (
    "image/color"
    "math/rand"
    "github.com/fogleman/gg"
)

func main() {
    dc := gg.NewContext(1024, 1024)
    dc.SetColor(color.Black)
    dc.Clear()

    for i := 0; i < 360; i += 3 {

        rotateDots(dc, 6, i, 256)
        rotateDots(dc, 12, i, 64)
        rotateDots(dc, 8, i, 128)
        rotateDots(dc, 4, i, 32)

    }

    dc.SavePNG("colorcircles.png")
}

func rotateDots (dc *gg.Context, y float64, degree int, radius float64) {
        const S = 1024

        dc.Push()

        dc.RotateAbout(gg.Radians(float64(degree)), S/2, S/2)

        dc.SetRGBA(rand.Float64(), rand.Float64(), rand.Float64(), 0.6)
        dc.DrawCircle(S/2, S/y, S/radius)
        dc.Fill()

        dc.Pop()
}
