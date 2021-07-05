package main

import (
	"./geo"
	"fmt"
	"io/ioutil"
	"os"
)

// Custom
func makeCustomGeoJson(raw []byte, title string, codes []string) {
	// Split Data
	cityMap := geo.Split(raw, geo.SPLIT_TYPE_CUSTOM, codes)
	// Loop
	var customFts []geo.Feature
	for _, fts := range cityMap {
		customFts = append(customFts, fts...)
	}
	// Save Dir
	dir := fmt.Sprintf("geojson/custom")
	if !isExist(dir) {
		os.MkdirAll(dir, 0777)
	}
	// Save Path
	path := fmt.Sprintf("%s/%s.json", dir, title)
	// Save Json
	geo.Save(path, customFts)
}

// City
func makeCityGeoJson(raw []byte) {

	// Split Data
	cityMap := geo.Split(raw, geo.SPLIT_TYPE_CITY, []string{})

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

	os.RemoveAll("geojson/prefectures")
	// Save Dir
	dir := "geojson/prefectures"
	if !isExist(dir) {
		os.MkdirAll(dir, 0777)
	}

	// Split Data
	cityMap := geo.Split(raw, geo.SPLIT_TYPE_PREF, []string{})
	// Loop
	for prefCode, fts := range cityMap {
		fmt.Println(prefCode, fts[0].GetPref())
		// Save Path
		path := fmt.Sprintf("%s/%s.json", dir, prefCode)
		// Save Json
		geo.Save(path, fts)
	}
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
	switch os.Args[2] {
	case "custom":
		makeCustomGeoJson(raw, os.Args[3], os.Args[4:])
	default:
		makeCityGeoJson(raw)
		makePrefGeoJson(raw)
	}
}

// Check Dir Exist
func isExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
