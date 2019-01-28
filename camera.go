package main

type camera struct {
	Origin, LowerLeft, Horizontal, Vertical vec3
}

func newCamera() camera {
	return camera{
		newVec3(0, 0, 0),
		newVec3(-2, -1, -1),
		newVec3(4, 0, 0),
		newVec3(0, 2, 0),
	}
}

func (c camera) ray(u, v float32) ray {
	return newRay(
		c.Origin,
		c.LowerLeft.Add(c.Horizontal.Multiply(u)).Add(c.Vertical.Multiply(v)))
}
