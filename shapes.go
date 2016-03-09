package main

import "math"

type HitRecord struct {
	T         float64
	P, Normal Vec3
}

type Hitable interface {
	Hit(r *Ray, t_min, t_max float64) *HitRecord
}

type Sphere struct {
	Center Vec3
	Radius float64
}

func (s Sphere) Hit(r *Ray, t_min, t_max float64) *HitRecord {
	oc := r.Origin().Sub(&s.Center)
	a := r.Direction().Dot(r.Direction())
	b := oc.Dot(r.Direction())
	c := oc.Dot(oc) - s.Radius*s.Radius
	discriminant := b*b - a*c
	if discriminant > 0 {
		temp := (-b - math.Sqrt(b*b-a*c)) / a
		if temp < t_max && temp > t_min {
			h := HitRecord{}
			h.T = temp
			h.P = *r.PointAtParameter(temp)
			h.Normal = *h.P.Sub(&s.Center).DivParts(s.Radius)
			return &h
		}
		temp = (-b + math.Sqrt(b*b-a*c)) / a
		if temp < t_max && temp > t_min {
			h := HitRecord{}
			h.T = temp
			h.P = *r.PointAtParameter(temp)
			h.Normal = *h.P.Sub(&s.Center).DivParts(s.Radius)
			return &h
		}
	}
	return nil
}

type ObjectList struct {
	Items []Hitable
}

func (l ObjectList) Hit(r *Ray, t_min, t_max float64) *HitRecord {
	var h *HitRecord = nil
	closest := t_max
	for i := 0; i < len(l.Items); i++ {
		result := l.Items[i].Hit(r, t_min, closest)
		if result != nil {
			closest = result.T
			h = result
		}
	}
	return h
}
