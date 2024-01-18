package behavioral_patterns

import "fmt"

type Shape interface {
	getType() string
	accept(Visitor)
}

type Rectangle struct {
	l, b int
}

func (r *Rectangle) accept(v Visitor) {
	v.visitRectangle(r)
}

func (r *Rectangle) getType() string {
	return "Rectangle"
}

type Circle struct {
	r int
}

func (c *Circle) accept(v Visitor) {
	v.visitCircle(c)
}

func (c *Circle) getType() string {
	return "Circle"
}

type Visitor interface {
	visitRectangle(*Rectangle)
	visitCircle(*Circle)
}

type AreaCalc struct {
	area int
}

func (a *AreaCalc) visitRectangle(r *Rectangle) {
	a.area = r.b * r.l
}

func (a *AreaCalc) visitCircle(c *Circle) {
	a.area = int(3.14 * float64(c.r * c.r))
}

type PerimeterCalc struct {
	perimeter int
}

func (p *PerimeterCalc) visitRectangle(r *Rectangle) {
	p.perimeter = 2 * (r.b + r.l)
}

func (p *PerimeterCalc) visitCircle(c *Circle) {
	p.perimeter = int(2 * 3.14 * float64(c.r))
}

func VisitorPattern() {
	fmt.Println("Visitor pattern")

	arr := []Shape{&Rectangle{l: 3, b: 4}, &Circle{r: 4}}

	fmt.Println("Calculation area of:")
	areaCalc := &AreaCalc{}
	for _, s := range arr {
		s.accept(areaCalc)
		fmt.Printf("  %s is %v\n", s.getType(), areaCalc.area)
	}

	fmt.Println("Calculation perimeter of:")
	perimeterCalc := &PerimeterCalc{}
	for _, s := range arr {
		s.accept(perimeterCalc)
		fmt.Printf("  %s is %v\n", s.getType(), perimeterCalc.perimeter)
	}
}