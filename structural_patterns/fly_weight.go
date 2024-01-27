package structural_patterns

import "fmt"

type iDress interface {
	getColor() string
}

type terroristDress struct {
	color string
}

func newTerroristDress() *terroristDress {
	return &terroristDress{"red"}
}

func (tst *terroristDress) getColor() string {
	return tst.color
}

type counterTerroristDress struct {
	color string
}

func newCounterTerroristDress() *counterTerroristDress {
	return &counterTerroristDress{"green"}
}

func (cst *counterTerroristDress) getColor() string {
	return cst.color
}

const (
	terroristDressType = "tDress"
	counterTerroristDressType = "ctDress"
)

var (
	dressFactorySingleInstance = &dressFactory{
		dressMap: make(map[string]iDress),
	}
)

type dressFactory struct {
	dressMap map[string]iDress
}

func getDressFactorySingleInstance() *dressFactory {
	return dressFactorySingleInstance
}

func (df *dressFactory) getDressByType(dressType string) (iDress, error) {
	if df.dressMap[dressType] != nil {
		return df.dressMap[dressType], nil
	}

	if dressType == terroristDressType {
		df.dressMap[dressType] = newTerroristDress()
		return df.dressMap[dressType], nil
	}

	if dressType == counterTerroristDressType {
		df.dressMap[dressType] = newCounterTerroristDress()
		return df.dressMap[dressType], nil
	}

	return nil, fmt.Errorf("invalid dress type")
}

type player struct {
	dress iDress
	playerType string
	lat int
	long int 
}

func newPlayer(playerType, dressType string) *player {
	dress, _ := getDressFactorySingleInstance().getDressByType(dressType)

	return &player{
		dress: dress,
		playerType: playerType,
	}
}

func (p *player) movePlayer(lat, long int) {
	p.lat = lat
	p.long = long
}

type game struct {
	terrorist []*player
	counterTerrorist []*player 
}

func newGame() *game {
	return &game{
		terrorist: make([]*player, 1),
		counterTerrorist: make([]*player, 1),
	}
}

func (g *game) addTerrorist() {
	player := newPlayer("T", terroristDressType)
	g.terrorist = append(g.terrorist, player)
}

func (g *game) addCounterTerrorist() {
	player := newPlayer("CT", counterTerroristDressType)
	g.counterTerrorist = append(g.counterTerrorist, player)
}

func FlyWeightPattern() {
	game := newGame()

    //Add Terrorist
    game.addTerrorist()
    game.addTerrorist()
    game.addTerrorist()
    game.addTerrorist()

    //Add CounterTerrorist
    game.addCounterTerrorist()
    game.addCounterTerrorist()
    game.addCounterTerrorist()

    dressFactoryInstance := getDressFactorySingleInstance()

    for dressType, dress := range dressFactoryInstance.dressMap {
			fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
    }
}