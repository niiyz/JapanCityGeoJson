package main

import (
	"./geo"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Not Exists GeoJsonPath")
		os.Exit(1)
	}
	geoJsonFilePath := os.Args[1]
	raw, err := ioutil.ReadFile(geoJsonFilePath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var fc geo.FeatureCollection

	json.Unmarshal(raw, &fc)

	var coordinates []geo.LatLng
	for _, ft := range fc.Features {
		for _, g1 := range ft.Geometry.Coordinates {
			for _, g2 := range g1 {
				for _, g3 := range g2 {
					coordinates = append(coordinates, g3)
				}
			}
		}
	}
	x0, y0, x1, y1 := math.Inf(0), math.Inf(0), math.Inf(-1), math.Inf(-1)
	for _, coordinate := range coordinates {
		if coordinate[0] < x0 {
			x0 = coordinate[0]
		}
		if coordinate[1] < y0 {
			y0 = coordinate[1]
		}
		if coordinate[0] > x1 {
			x1 = coordinate[0]
		}
		if coordinate[1] > y1 {
			y1 = coordinate[1]
		}
	}
	fmt.Println(x0, y0, x1, y1)
}