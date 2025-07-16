package main

import ("fmt"; "os"; "strconv"; "log"; "math")
// import ("fmt"; "math")

func main() {
	maximum := math.Inf(-1)
	inputing := os.Args[1:]
	for _, value := range inputing {
		value, error := strconv.ParseFloat(value, 64)
		if error != nil {
			log.Fatal(error)
		}
		if maximum <= value {
			maximum = value
		}
	}
	fmt.Println(maximum)
}