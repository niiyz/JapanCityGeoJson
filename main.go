package main

import (
	"./geo"
	"fmt"
	"io/ioutil"
	"os"
)

func makeCityTest(targetPref string, targetCity string, raw []byte) {
	// Split Data
	cityMap := geo.Split(geo.SPLIT_TYPE_CITY, raw)

	// Loop
	for city, fts := range cityMap {
		// Pref Name
		pref := fts[0].GetPref()
		if targetPref != pref || targetCity != targetCity {
			return
		}
		// fmt.Println(pref, city, len(fts))
		dir := "test"
		if !isExist(dir) {
			os.MkdirAll(dir, 0777)
		}
		// Save Path
		path := dir + "/" + city + ".json"
		// Save Json
		geo.Save(path, fts)
	}
}

// City
func makeCityGeoJson(raw []byte) {
	// Split Data
	cityMap := geo.Split(geo.SPLIT_TYPE_CITY, raw)

	// Loop
	for city, fts := range cityMap {
		// Pref Name
		pref := fts[0].GetPref()
		fmt.Println(pref, city, cap(fts))
		// Save Dir
		dir := "geojson/" + pref
		if !isExist(dir) {
			os.MkdirAll(dir, 0777)
		}
		// Save Path
		path := dir + "/" + city + ".json"
		// Save Json
		geo.Save(path, fts)
	}
}

// County
func makeCountyGeoJson(raw []byte) {
	// Split Data
	cityMap := geo.Split(geo.SPLIT_TYPE_COUNTY, raw)

	// Loop
	for county, fts := range cityMap {
		// Pref Name
		pref := fts[0].GetPref()
		fmt.Println(county, cap(fts))
		// Save Dir
		dir := "geojson/" + pref
		if !isExist(dir) {
			os.MkdirAll(dir, 0777)
		}
		// Save Path
		path := dir + "/" + county + ".json"
		// Save Json
		geo.Save(path, fts)
	}
}

// Pref
func makePrefGeoJson(raw []byte) {
	// Split Data
	cityMap := geo.Split(geo.SPLIT_TYPE_PREF, raw)
	// Loop
	for pref, fts := range cityMap {
		fmt.Println(pref, cap(fts))
		// Save Dir
		dir := "geojson/prefectures"
		if !isExist(dir) {
			os.MkdirAll(dir, 0777)
		}
		// Save Path
		path := dir + "/" + pref + ".json"
		// Save Json
		geo.Save(path, fts)
	}
}

func reset() {
	os.RemoveAll("geojson")
}

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
	// reset()
	// City
	makeCityTest("富山県", "氷見市", raw)
	// Pref
	// makePrefGeoJson(raw)
	// County
	// makeCountyGeoJson(raw)
}

// Check Dir Exist
func isExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
