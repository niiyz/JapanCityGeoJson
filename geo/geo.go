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

func (ft Feature) GetCity() string {
	var city string
	if ft.Properties.City == "所属未定地" {
		city = ft.Properties.Pref + ft.Properties.City
	} else {
		city = ft.Properties.County + ft.Properties.City
	}
	return city
}

func (ft Feature) GetCounty() string {
	return ft.Properties.County
}

func (ft Feature) GetPref() string {
	return ft.Properties.Pref
}

type Properties struct {
	Pref    string `json:"N03_001"` // "富山県"
	SubPref string `json:"N03_002"` // null 北海道のみ有効
	County  string `json:"N03_003"` // "下新川郡"
	City    string `json:"N03_004"` // "朝日町"
	Code    string `json:"N03_007"` // "16343"
}

type Geometry struct {
	Type        string     `json:"type"` // Polygon
	Coordinates [][]LatLng `json:"coordinates"`
}

type LatLng [2]float64

func Split(splitType string, raw []byte) map[string][]Feature {

	var fc FeatureCollection

	json.Unmarshal(raw, &fc)

	cityMap := make(map[string][]Feature)

	for _, ft := range fc.Features {
		var key string
		switch splitType {
		case SPLIT_TYPE_CITY:
			key = ft.GetCity()
		case SPLIT_TYPE_PREF:
			key = ft.GetPref()
		case SPLIT_TYPE_COUNTY:
			county := ft.GetCounty()
			if county == "" {
				continue
			}
			key = county
		}
		cityMap[key] = append(cityMap[key], ft)
	}
	return cityMap
}
