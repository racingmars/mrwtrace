package main

type Camera struct {
	Origin     Vec3
	LowerLeft  Vec3
	Vertical   Vec3
	Horizontal Vec3
}

func (c *Camera) GetRay(u, v float64) Ray {
	return Ray{c.Origin, *c.LowerLeft.Add(c.Horizontal.MultParts(u)).Add(
		c.Vertical.MultParts(v)) /*.Sub(&c.Origin)*/}
}
