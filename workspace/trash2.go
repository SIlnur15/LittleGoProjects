package main

import (
	"fmt"
	"math"
	"sort"
)

// getAverage вычисляет среднее значение элементов слайса
func getAverage(numbers ...int) int {
	sum := 0
	for _, numb := range numbers {
		sum += numb
	}
	return sum / len(numbers)
}

// getMedian вычисляет медиану элементов слайса
// медиана – центральный элемент слайса (отсортированного)
func getMedian(numbers []int) int {
	sort.Ints(numbers) // сортировка слайса
	if len(numbers)%2 == 0 {
		return int((numbers[len(numbers)/2] + numbers[len(numbers)/2-1]) / 2)
	}
	return numbers[len(numbers)/2]
}

// getDistribution вычисляет размах (макс - мин) элементов слайса
func getDistribution(numbers []int) int {
	maxim := math.Inf(-1)
	minim := math.Inf(1)
	for _, numb := range numbers {
		if float64(numb) > maxim {
			maxim = float64(numb)
		}
		if float64(numb) < minim {
			minim = float64(numb)
		}
	}
	return int(maxim - minim)
}

// getMode вычисляет моду
// эта функция должна вычислять моду (значение, которое встречается чаще всего)
func getMode(numbers []int) int {
	frequencyMap := make(map[int]int)
	for _, numb := range numbers {
		frequencyMap[numb]++
	}
	maximFrequency := math.Inf(-1)
	for key := range frequencyMap {
		if float64(frequencyMap[key]) >= maximFrequency {
			maximFrequency = float64(frequencyMap[key])
		}
	}
	maximFrequencyKeys := make([]int, 0)
	for key := range frequencyMap {
		if float64(frequencyMap[key]) == maximFrequency {
			maximFrequencyKeys = append(maximFrequencyKeys, key)
		}
	}
	sort.Ints(maximFrequencyKeys)
	return maximFrequencyKeys[len(maximFrequencyKeys)-1]
}

// getVariance вычисляет дисперсию (среднее квадратичное отклонение от среднего значения)
func getVariance(numbers ...int) int {
	mean := getAverage(numbers...)
	var sum int
	for _, num := range numbers {
		sum += (num - mean) * (num - mean)
	}
	return int(sum / len(numbers))
}

// analyseSlice вызывает все функции
// используйте ваши функции getAverage, getMedian, getDistribution, getMode и getVariance
func analyseSlice(slice []int) (average int, median int, distribution int, mode int, variance int) {
	if len(slice) == 0 { // на случай пустого слайса
		return 0, 0, 0, 0, 0
	}
	average = getAverage(slice...)
	median = getMedian(slice)
	distribution = getDistribution(slice)
	mode = getMode(slice)
	variance = getVariance(slice...)
	return average, median, distribution, mode, variance
}

func main() {
	// не изменяйте функцию main!!
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	average, median, distribution, mode, variance := analyseSlice(nums)
	fmt.Println(average, median, distribution, mode, variance)
}
