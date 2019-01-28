package main

import (
	"math"
)

type hitContext struct {
	T, TMin, TMax float32
	P, Normal     vec3
}

func newHitContext() *hitContext {
	return &hitContext{
		TMin: 0.001,
		TMax: math.MaxFloat32,
	}
}

func (ctx *hitContext) Valid(t float32) bool {
	return t < ctx.TMax && t > ctx.TMin
}

type hittable interface {
	hit(ctx *hitContext, r ray) bool
}

type hittables []hittable

func (hs hittables) hit(ctx *hitContext, r ray) bool {
	for _, h := range hs {
		if h.hit(ctx, r) {
			return true
		}
	}

	return false
}
