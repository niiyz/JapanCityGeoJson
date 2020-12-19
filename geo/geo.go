package geo

import (
	"encoding/json"
)

const (
	SPLIT_TYPE_CITY   string = "CITY"
	SPLIT_TYPE_PREF   string = "PREF"
	SPLIT_TYPE_COUNTY string = "COUNTY"
)

type FeatureCollection struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}

type Feature struct {
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
	Geometry   Geometry   `json:"geometry"`
}

func (ft Feature) GetPrefCode() string {
	if ft.Properties.Code == "" {
		return GetPrefCodeByName(ft.Properties.Pref)
	}
	return ft.Properties.Code[:2]
}

func (ft Feature) GetCode() string {
	if ft.Properties.Code == "" {
		return "UNDECIDED_LAND"
	}
	return ft.Properties.Code
}

func (ft Feature) GetCity() string {
	return ft.Properties.City
}

func (ft Feature) GetCounty() string {
	return ft.Properties.County
}

func (ft Feature) GetPref() string {
	return ft.Properties.Pref
}

func (ft Feature) GetGeometryCoordinates() [][][]LatLng {
	return ft.Geometry.Coordinates
}

type Properties struct {
	Pref    string `json:"N03_001"` // "富山県"
	SubPref string `json:"N03_002"` // null 北海道のみ有効
	County  string `json:"N03_003"` // "下新川郡"
	City    string `json:"N03_004"` // "朝日町"
	Code    string `json:"N03_007"` // "16343"
}

type Geometry struct {
	Type        string       `json:"type"` // Polygon
	Coordinates [][][]LatLng `json:"coordinates"`
}

type LatLng [2]float64

func Split(splitType string, raw []byte) map[string][]Feature {

	var fc FeatureCollection

	json.Unmarshal(raw, &fc)

	ftsByKey := make(map[string][]Feature)

	for _, ft := range fc.Features {
		var key string
		switch splitType {
		case SPLIT_TYPE_CITY:
			key = ft.GetCode()
		case SPLIT_TYPE_PREF:
			key = ft.GetPref()
		case SPLIT_TYPE_COUNTY:
			county := ft.GetCounty()
			if county == "" {
				continue
			}
			key = county
		}
		ftsByKey[key] = append(ftsByKey[key], ft)
	}

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
