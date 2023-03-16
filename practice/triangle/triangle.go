package triangle

type Triangle struct {
	Base float32
	Tall float32
}

func (triangle Triangle) Wide() float32 {
	return 0.5 * triangle.Base * triangle.Tall
}
