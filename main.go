package main

import (
	"./geo"
	"fmt"
	"io/ioutil"
	"os"
)

// City
func makeCityGeoJson(raw []byte) {
	// Split Data
	cityMap := geo.Split(geo.SPLIT_TYPE_CITY, raw)

	// Loop
	for cityCode, fts := range cityMap {
		prefCode := fts[0].GetPrefCode()
		fmt.Println(cityCode, prefCode, fts[0].GetPref(), fts[0].GetCounty(), fts[0].GetCity())
		// Save Dir
		dir := fmt.Sprintf("geojson/%s", prefCode)
		if !isExist(dir) {
			os.MkdirAll(dir, 0777)
		}
		// Save Path
		path := fmt.Sprintf("%s/%s.json", dir, cityCode)
		// Save Json
		geo.Save(path, fts)
	}
}

// Pref
func makePrefGeoJson(raw []byte) {
	// Split Data
	cityMap := geo.Split(geo.SPLIT_TYPE_PREF, raw)
	// Loop
	for prefCode, fts := range cityMap {
		fmt.Println(prefCode, fts[0].GetPref())
		// Save Dir
		dir := "geojson/prefectures"
		if !isExist(dir) {
			os.MkdirAll(dir, 0777)
		}
		// Save Path
		path := fmt.Sprintf("%s/%s.json", dir, prefCode)
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
	// makeCityGeoJson(raw)
	// Pref
	makePrefGeoJson(raw)
}

// Check Dir Exist
func isExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
