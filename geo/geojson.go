package geo

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
