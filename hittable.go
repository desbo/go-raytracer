package main

type hittable interface {
	hit(r ray) float32 // the point along the ray that hits the object. < 0 means no hit
}

type hittables []hittable

func (hs hittables) hit(r ray) float32 {
	for _, h := range hs {
		if t := h.hit(r); t >= 0.0 {
			println("HIT")
			return t
		}
	}

	return -1
}
