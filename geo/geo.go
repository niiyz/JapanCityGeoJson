package geo

import (
	"encoding/json"
	"fmt"
)

const (
	SPLIT_TYPE_CUSTOM string = "CUSTOM"
	SPLIT_TYPE_CITY   string = "CITY"
	SPLIT_TYPE_PREF   string = "PREF"
	SPLIT_TYPE_COUNTY string = "COUNTY"
)

func inSlice(str string, cods []string) bool {
	for _, v := range cods {
		if v == str {
			return true
		}
	}
	return false
}

func mergeFeature(ftsByKey map[string][]Feature) map[string][]Feature {
	multiPolygonMergedFtsByKey := make(map[string][]Feature)
	for key, fts := range ftsByKey {
		var code [][]LatLng
		for _, ft := range fts {
			code = append(code, ft.GetGeometryCoordinates()[0][0])
		}
		fts[0].Geometry.Coordinates = [][][]LatLng{code}
		multiPolygonMergedFtsByKey[key] = []Feature{fts[0]}
	}

	return multiPolygonMergedFtsByKey
}

func Split(raw []byte, splitType string, codes []string) map[string][]Feature {

	var fc FeatureCollection

	json.Unmarshal(raw, &fc)

	ftsByKey := make(map[string][]Feature)

	for _, ft := range fc.Features {
		var key string
		switch splitType {
		case SPLIT_TYPE_CUSTOM:
			if !inSlice(ft.GetCode(), codes) {
				continue
			}
			fmt.Println(ft.GetCity())
			key = ft.GetCode()
		case SPLIT_TYPE_CITY:
			key = ft.GetCode()
		case SPLIT_TYPE_PREF:
			key = ft.GetPrefCode()
		case SPLIT_TYPE_COUNTY:
			county := ft.GetCounty()
			if county == "" {
				continue
			}
			key = county
		}
		ftsByKey[key] = append(ftsByKey[key], ft)
	}

	return mergeFeature(ftsByKey)
}
