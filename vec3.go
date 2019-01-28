package main

import (
	"fmt"
	"math"
	"math/rand"
)

type vec3 struct {
	X, Y, Z float32
}

func newVec3(x, y, z float32) vec3 {
	return vec3{x, y, z}
}

func randomVec3() vec3 {
	return newVec3(rand.Float32(), rand.Float32(), rand.Float32())
}

func (v vec3) Length() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z)))
}

func (v vec3) DivideVector(w vec3) vec3 {
	return vec3{v.X / w.X, v.Y / w.Y, v.Z / w.Z}
}

func (v vec3) DivideFloat(x float32) vec3 {
	return vec3{v.X / x, v.Y / x, v.Z / x}
}

func (v vec3) Unit() vec3 {
	return v.DivideFloat(v.Length())
}

func (v vec3) String() string {
	return fmt.Sprintf("(%0.4f, %0.4f, %0.4f)", v.X, v.Y, v.Z)
}

func (v vec3) Add(w vec3) vec3 {
	return vec3{
		v.X + w.X,
		v.Y + w.Y,
		v.Z + w.Z,
	}
}

func (v vec3) Subtract(w vec3) vec3 {
	return vec3{v.X - w.X, v.Y - w.Y, v.Z - w.Z}
}

func (v vec3) Multiply(x float32) vec3 {
	return vec3{
		v.X * x,
		v.Y * x,
		v.Z * x,
	}
}

func (v vec3) Dot(w vec3) float32 {
	return v.X*w.X + v.Y*w.Y + v.Z*w.Z
}
