package main

type car struct {
	Model string
	Color string
}
// para declarar, podemos fazer assim:
// xyz :  car
func main() {
	car := car{ // declarando e definindo os valores
		Model: "Fiat",
		Color: "blue",
	}
	println(car.Model, car.Color)
	car.Model = "Ford"
	println(car.Model, car.Color)
}
