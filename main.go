package main

type Car struct {
	Model string
	Color string
}

// para declarar, podemos fazer assim:
// xyz :  car

func (c Car) Start() {
	println("Starting the car", c.Model)
}

func soma(a, b int) int {
	return a + b
}

func main() {
	car := Car{ // declarando e definindo os valores
		Model: "Fiat",
		Color: "blue",
	}
	println(car.Model, car.Color)
	car.Model = "Ford"
	println(car.Model, car.Color)
	car.Start()
}
