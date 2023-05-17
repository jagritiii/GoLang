package composition

// struct inside another structure
type engine struct {
	hp int
}

type wheel struct {
	wheelDimension int
}

type Car struct {
	CarName string
	Engine  engine
	Wheel   wheel
}

func NewCar(carName string, hp int, wd int) Car {
	return Car{
		CarName: carName,
		Engine: engine{
			hp,
		},
		Wheel: wheel{
			wheelDimension: wd,
		},
	}
}
