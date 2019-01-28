package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
)

const (
	width   = 1200
	height  = 600
	samples = 16
)

func colour(ctx *hitContext, world hittables, r ray) vec3 {
	if world.hit(ctx, r) {
		target := ctx.P.Add(ctx.Normal).Add(randomPointInUnitSphere())
		return colour(ctx, world, newRay(ctx.P, target.Subtract(ctx.P))).Multiply(0.5)
	}

	unitDirection := r.direction().Unit()
	t := 0.5 * (unitDirection.Y + 1)
	return newVec3(1, 1, 1).Multiply(1 - t).Add(newVec3(0.5, 0.7, 1.0).Multiply(t))
}

func main() {
	rect := image.Rect(width, height, 0, 0)
	img := image.NewRGBA(rect)
	ctx := newHitContext()

	world := hittables{
		newSphere(newVec3(0, 0, -1), 0.5),
		newSphere(newVec3(0, -100.5, -1), 100),
	}

	cam := newCamera()

	for y := height - 1; y >= 0; y-- {
		for x := 0; x < width; x++ {
			col := newVec3(0, 0, 0)

			for s := 0; s < samples; s++ {
				u := (float32(x) + rand.Float32()) / float32(width)
				v := (float32(y) + rand.Float32()) / float32(height)
				r := cam.ray(u, v)
				col = col.Add(colour(ctx, world, r))
			}

			col = col.DivideFloat(float32(samples))
			col = newVec3(
				float32(math.Sqrt(float64(col.X))),
				float32(math.Sqrt(float64(col.Y))),
				float32(math.Sqrt(float64(col.Z))))

			img.Set(x, height-y, color.RGBA{
				R: uint8(255.99 * col.X),
				G: uint8(255.99 * col.Y),
				B: uint8(255.99 * col.Z),
				A: uint8(255),
			})
		}
	}

	png.Encode(os.Stdout, img)
}
