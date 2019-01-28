package main

import (
	"math"
)

type sphere struct {
	Centre vec3
	Radius float32
}

func newSphere(centre vec3, radius float32) sphere {
	return sphere{centre, radius}
}

func randomPointInUnitSphere() vec3 {
	for {
		if p := randomVec3().Multiply(2).Subtract(newVec3(1, 1, 1)); p.Length()*p.Length() >= 1 {
			return p
		}
	}
}

func (s sphere) hit(ctx *hitContext, r ray) bool {
	oc := r.origin().Subtract(s.Centre)
	a := r.direction().Dot(r.direction())
	b := 2 * oc.Dot(r.direction())
	c := oc.Dot(oc) - s.Radius*s.Radius
	disc := b*b - 4*a*c

	if disc > 0 {
		t := (-b - float32(math.Sqrt(float64(disc)))) / (2.0 * a)

		if ctx.Valid(t) {
			ctx.T = t
			ctx.P = r.pointAt(t)
			ctx.Normal = (ctx.P.Subtract(s.Centre)).DivideFloat(s.Radius)
			return true
		}
	}

	return false
}
