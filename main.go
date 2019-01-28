package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

const (
	width   = 800
	height  = 400
	samples = 50
)

func main() {
	rect := image.Rect(width, height, 0, 0)
	img := image.NewRGBA(rect)

	world := hittables{
		newSphere(newVec3(0, 0, -1), 0.5),
		// newSphere(newVec3(0, -100.5, -1), 100),
	}

	cam := newCamera()

	colour := func(r ray) vec3 {
		t := world.hit(r)
		if t > 0 {
			N := r.pointAt(t).Subtract(newVec3(0, 0, -1)).Unit()
			return newVec3(N.X+1, N.Y+1, N.Z+1).Multiply(0.5)
		}

		unitDirection := r.direction().Unit()
		t = 0.5 * (unitDirection.Y + 1)
		return newVec3(1, 1, 1).Multiply(1 - t).Add(newVec3(0.5, 0.7, 1.0).Multiply(t))
	}

	for y := height - 1; y >= 0; y-- {
		for x := 0; x < width; x++ {
			u := float32(x) / float32(width)
			v := float32(y) / float32(height)
			r := cam.ray(u, v)
			col := colour(r)

			// col := newVec3(0, 0, 0)

			// for s := 0; s < samples; s++ {
			// 	u := (float32(x) + rand.Float32()) / float32(width)
			// 	v := (float32(y) + rand.Float32()) / float32(height)
			// 	r := cam.ray(u, v)
			// 	col = col.Add(colour(r))
			// }

			// col = col.DivideFloat(float32(samples))

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
