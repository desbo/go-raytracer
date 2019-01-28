package main

type sphere struct {
	Centre vec3
	Radius float32
}

func newSphere(centre vec3, radius float32) sphere {
	return sphere{centre, radius}
}

func (s sphere) hit(r ray) float32 {
	oc := r.origin().Subtract(s.Centre)
	a := r.direction().Dot(r.direction())
	b := 2 * oc.Dot(r.direction())
	c := oc.Dot(oc) - s.Radius*s.Radius

	return (b*b - 4*a*c)
}
