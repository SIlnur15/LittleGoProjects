package main

import ("fmt"; "errors"; "log")

func paintCalculator(width, height float64) (float64, error) {
	if width <= 0 || height <= 0 {
		err := errors.New("entered data is not correct")
		return 0, err
	}
	area := width*height/10.0
	return area, nil
}

func main() {
	liters, state := paintCalculator(0.0, 15.0)
	if state != nil {
		log.Fatal(state)
	}
	fmt.Printf("You need %.2f liters\n", liters)
}
