package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy) //allocate the outer slice, dy = number of rows
	for y := 0; y < dy; y++ {
		pic[y] = make([]uint8, dx) // allocate each inner slice
		for x := 0; x < dx; x++ {
			// Example function: bitwise XOR
			pic[y][x] = uint8((x ^ y) / 2) 
		}
	}
	return pic
}

// The range form of the for loop iterates over a slice or map.
// When ranging over a slice, two values are returned for each iteration. 
// The first is the index, and the second is a copy of the element at that index.
func PicRangeIteration(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy) //allocate the outer slice, dy = number of rows
	// range iteration without using the value, only the index 
	for i := range dy {
		pic[i] = make([]uint8, dx) // allocate each inner slice
		for j := range dx {
			// Example function: bitwise XOR
			pic[i][j] = uint8((i ^ j) / 2) 
		}
	}
	return pic
}


type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func createMap() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	m["Athens"] = Vertex{
		128.68433, -44.39967,
	}
	fmt.Println(m["Bell Labs"]) // {40.68433 -74.39967}
	fmt.Println(m["Bell Labs"].Long) // -74.39967

	fmt.Println(m) // map[Athens:{128.68433 -44.39967} Bell Labs:{40.68433 -74.39967}]

	var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m) // map[Google:{37.42202 -122.08408} Bell Labs:{40.68433 -74.39967}]
}

func mapOperations() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"]) // The value: 42

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"]) // The value: 48

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"]) // The value: 0

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok) // The value: 0 Present? false
}

func WordCount(s string) map[string]int {
	counts := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		// if _, ok := counts[word]; ok {
		// 	counts[word] = counts[word] + 1
		// } else {
		// 	counts[word] = 1
		// }
		counts[word]++
	}
	return counts

}

func main() {
	pic.Show(Pic)
	pic.Show(PicRangeIteration)
	createMap();
}
