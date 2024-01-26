package creational_patterns

import "fmt"

type iGun interface {
	getName() string
	setName(string)
	getPower() int
	setPower(int)
}

type gun struct {
	name string
	power int
}

func (g *gun) getName() string {
	return g.name
}

func (g *gun) setName(name string) {
	g.name = name
}

func (g *gun) getPower() int {
	return g.power
}

func (g *gun) setPower(power int) {
	g.power = power
}

type ak47 struct {
	gun
}

func newAk47() iGun {
	return &ak47{
		gun: gun{
			name: "AK 47",
			power: 44,
		},
	}
}

type musket struct {
	gun
}

func newMusket() iGun {
	return &musket{
		gun: gun{
			name: "Musket",
			power: 23,
		},
	}
}

type gunFactory struct {}

func (gf *gunFactory) getGun(gunType string) (iGun, error) {
	if gunType == "AK47" {
		return newAk47(), nil
	}

	if gunType == "Musket" {
		return newMusket(), nil
	}

	return nil, fmt.Errorf("invalid gun type %v", gunType)
} 

func Factory_Pattern() {
	gunFty := &gunFactory {}
	gunType := "AK47"

	myGun, _ := gunFty.getGun(gunType)

	fmt.Printf("My gun is %v", myGun)
}