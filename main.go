package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

func main() {
	nx := 800
	ny := 400
	ns := 100

	fmt.Println("P3")
	fmt.Println(strconv.Itoa(nx) + " " + strconv.Itoa(ny))
	fmt.Println("255")

	lowerLeft := Vec3{-2.0, -1.0, -1.0}
	horizontal := Vec3{4.0, 0.0, 0.0}
	vertical := Vec3{0.0, 2.0, 0.0}
	origin := Vec3{0.0, 0.0, 0.0}

	cam := Camera{Origin: origin, LowerLeft: lowerLeft, Vertical: vertical,
		Horizontal: horizontal}

	sphere1 := Sphere{Center: Vec3{0, 0, -1}, Radius: 0.5}
	sphere2 := Sphere{Center: Vec3{0, -100.5, -1}, Radius: 100}
	world := ObjectList{Items: []Hitable{sphere1, sphere2}}

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			col := &Vec3{0, 0, 0}
			for s := 0; s < ns; s++ {
				u := (float64(i) + rand.Float64()) / float64(nx)
				v := (float64(j) + rand.Float64()) / float64(ny)
				r := cam.GetRay(u, v)
				col = col.Add(color(&r, world))
			}
			col = col.DivParts(float64(ns))

			ir := int(255.99 * col.X)
			ig := int(255.99 * col.Y)
			ib := int(255.99 * col.Z)

			fmt.Println(strconv.Itoa(ir) + " " + strconv.Itoa(ig) + " " +
				strconv.Itoa(ib))
		}
	}
}

func color(r *Ray, world Hitable) *Vec3 {
	if h := world.Hit(r, 0.0, math.MaxFloat64); h != nil {
		temp := Vec3{h.Normal.X + 1, h.Normal.Y + 1, h.Normal.Z + 1}
		return temp.MultParts(0.5)
	} else {
		unitDirection := r.Direction().UnitVector()
		t := 0.5 * (unitDirection.Y + 1.0)
		one := &Vec3{1.0, 1.0, 1.0}
		colors := &Vec3{0.5, 0.7, 1.0}
		return one.MultParts(1.0 - t).Add(colors.MultParts(t))
	}
}

func hitSphere(center *Vec3, radius float64, r *Ray) float64 {
	oc := r.Origin().Sub(center)
	a := r.Direction().Dot(r.Direction())
	b := 2.0 * oc.Dot(r.Direction())
	c := oc.Dot(oc) - radius*radius
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return -1.0
	} else {
		return (-b - math.Sqrt(discriminant)) / (2.0 * a)
	}
}
