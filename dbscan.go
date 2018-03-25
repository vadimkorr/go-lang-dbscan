package main


import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"math"
)

type Point struct {
	Id		int
	Label	string
	X		int `json:"x"`
	Y 		int `json:"y"`
}

func Read(file string) []Point {
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return nil
	} else {
		byteValue, _ := ioutil.ReadAll(jsonFile)
		var points []Point
		json.Unmarshal([]byte(byteValue), &points)
		defer jsonFile.Close()
		return points
	}
}

func Init(points *[]Point) {
	for i:=0; i < len(*points); i++ {
		(*points)[i].Id = i;
		(*points)[i].Label = "undefined";		
	}
}

func Dist(p1, p2 Point) float64 {
	var res float64 = math.Sqrt(math.Pow(float64(p2.X - p1.X), 2) + math.Pow(float64(p2.Y - p1.Y), 2));
	return res;
}

func RangeQuery(db *[]Point, q Point, eps float64) []Point {
	var neighbors []Point
	for i := 0; i < len(*db); i++ {
		if (Dist((*db)[i], q) <= eps) {
			neighbors = append(neighbors, (*db)[i])
		}
	}
	return neighbors;
}

func DBSCAN(db *[]Point, eps float64, minPts int) {
	var c int = 0; // Cluster counter
	for i := 0; i < len(*db); i++ {
		if (*db)[i].Label != "undefined" {
			continue
		} // Previously processed in inner loop
		var neighbors = RangeQuery(db, (*db)[i], eps); // Find neighbors
		if (len(neighbors) < minPts) { // Density check
			(*db)[i].Label = "Noise" // Label as Noise
			continue
		}
		c++ // next cluster label
		(*db)[i].Label = strconv.Itoa(c) // Label initial point
		var seeds = neighbors // .filter(n => n.id !== p.id); // Neighbors to expand
		for j := 0; j < len(seeds); j++ { // Process every seed point
			if (seeds[j].Id == (*db)[i].Id) {
				continue
			}
			if (seeds[j].Label == "Noise") {
				seeds[j].Label = strconv.Itoa(c)
			} // Change Noise to border point
			if (seeds[j].Label != "undefined") {
				continue
			} // Previously processed
			seeds[j].Label = strconv.Itoa(c); // Label neighbor
			var seedNeighbors = RangeQuery(db, seeds[j], eps); // Find neighbors
			if (len(seedNeighbors) >= minPts) { // Density check
				for k := 0; k < len(seedNeighbors); k++ {
					seeds = append(seeds, seedNeighbors[k]); // Add new neighbors to seed set
				}
				fmt.Println("new len" + strconv.Itoa(len(seeds)))
			}
		}
	}

	fmt.Println("Done")

}

func main() {
	
	var points []Point = Read("input1.json")
	pointer := &points

	Init(&points)

	var eps float64 = 2.5
	var minPts int = 4
	DBSCAN(&points, eps, minPts)

	for i := 0; i < len(*pointer); i++ {
		fmt.Println("Id: " + strconv.Itoa((*pointer)[i].Id)) // strconv.Itoa(points[i].X))
		fmt.Println("Label: " + (*pointer)[i].Label) //strconv.Itoa(points[i].Y))
	}

	fmt.Println("Successfully Opened data.json")
	// defer the closing of our jsonFile so that we can parse it later on
	
}