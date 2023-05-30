package main

import "fmt"

const tempKelvin = 373

func main() {
	tempCelcius := tempKelvin - 273
	fmt.Printf("A temperatura de ebulição da água em ºK = %v, e a temperatura de ebulição da água em ºC = %v.", tempKelvin, tempCelcius)
}