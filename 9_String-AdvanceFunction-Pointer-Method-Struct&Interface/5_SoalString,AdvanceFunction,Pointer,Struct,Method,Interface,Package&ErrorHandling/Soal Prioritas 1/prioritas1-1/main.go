package main

import "fmt"

type Car struct {
	Type   string
	FuelIn float64
}

func (c Car) EstimateDistance() float64 {
	// Menghitung perkiraan jarak yang bisa ditempuh berdasarkan bahan bakar
	// yang sedang terisi dengan asumsi 1.5 L / mill
	return c.FuelIn * 1.5
}

func main() {
	// Membuat instance Car
	car := Car{
		Type:   "Sedan",
		FuelIn: 20.0,
	}

	// Menghitung perkiraan jarak yang bisa ditempuh
	distance := car.EstimateDistance()

	// Menampilkan hasil perhitungan
	fmt.Printf("Car Type: %s\n", car.Type)
	fmt.Printf("Fuel In: %.2f L\n", car.FuelIn)
	fmt.Printf("Estimated Distance: %.2f mill\n", distance)
}
