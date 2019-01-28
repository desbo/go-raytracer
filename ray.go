package main

// A ray is a function p(t) = A + t*B. Here p is a 3D position along a line
// in 3D.
type ray struct {
	A vec3 // origin
	B vec3 // direction
}

func newRay(a, b vec3) ray {
	return ray{a, b}
}

// the p function mentioned above
func (r ray) pointAt(t float32) vec3 {
	return r.A.Add(r.B.Multiply(t))
}

func (r ray) origin() vec3 {
	return r.A
}

func (r ray) direction() vec3 {
	return r.B
}
