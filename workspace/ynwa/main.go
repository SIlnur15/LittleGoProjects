package main

func main() {
	filteredStrings := func(slc []string, callback func(string) bool) []string {
		result := make([]string, 0)
		for _, elem := range slc {
			if callback(elem) {
				ret = append(ret, elem)
			}
		}
		return result
	}
}
