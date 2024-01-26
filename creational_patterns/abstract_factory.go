package creational_patterns

import "fmt"

type iSportFactory interface {
	makeShoe() iShoe
	makeShirt() iShirt
}

func getSportFactory(brand string) (iSportFactory, error) {
	if brand == "adidas" {
		return &adidas{}, nil
	}

	if brand == "nike" {
		return &nike{}, nil
	}

	return nil, fmt.Errorf("unknown brand %v passed", brand)
}

type adidas struct {}

func (fac *adidas) makeShoe() iShoe {
	return &adidasShoe {
		shoe: shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (fac *adidas) makeShirt() iShirt {
	return &adidasShirt {
		shirt: shirt{
			logo: "adidas",
			size: 32,
		},
	}
}

type nike struct {}

func (fac *nike) makeShoe() iShoe {
	return &nikeShoe {
		shoe: shoe{
			logo: "nike",
			size: 14,
		},
	}
}

func (fac *nike) makeShirt() iShirt {
	return &nikeShirt {
		shirt: shirt{
			logo: "nike",
			size: 32,
		},
	}
}

type iShoe interface {
	getLogo() string
	setLogo(string)
	getSize() int
	setSize(int)
}

type shoe struct {
	logo string
	size int
}

func (s *shoe) getLogo() string {
	return s.logo
}

func (s *shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *shoe) getSize() int {
	return s.size
}

func (s *shoe) setSize(size int) {
	s.size = size
}

type adidasShoe struct {
	shoe
}

type nikeShoe struct {
	shoe
}

type iShirt interface {
	getLogo() string
	setLogo(string)
	getSize() int
	setSize(int)
}

type shirt struct {
	logo string
	size int
}

func (s *shirt) getLogo() string {
	return s.logo
}

func (s *shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *shirt) getSize() int {
	return s.size
}

func (s *shirt) setSize(size int) {
	s.size = size
}

type adidasShirt struct {
	shirt
}

type nikeShirt struct {
	shirt
}

func Abstract_Factory_Pattern() {
	/**
	* this should be set on runtime based on the platform
	*/
	brand := "adidas"

	factory, _ := getSportFactory(brand)

	shoe := factory.makeShoe()
	shirt := factory.makeShirt()

	fmt.Printf("My shoe is %v\n", shoe)
	fmt.Printf("My shirt is %v\n", shirt)

}