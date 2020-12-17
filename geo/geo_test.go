package geo

import (
	"reflect"
	"testing"
)

func TestSplitCity(t *testing.T) {
	json := `
{
  "type": "FeatureCollection",
  "features": [
    {
      "type": "Feature",
      "properties": {
        "N03_001": "富山県",
        "N03_002": "",
        "N03_003": "",
        "N03_004": "氷見市",
        "N03_007": "11111"
		
      },
      "geometry": {
        "type": "Polygon",
        "coordinates": [
          [
            [
              136.111111,
              36.111111
            ],
            [
              136.222222,
              36.222222
            ]
          ]
        ]
      }
    }]
}`

	actual := Split(SPLIT_TYPE_CITY, []byte(json))
	expected := map[string][]Feature{
		"氷見市": []Feature{{"Feature",
			Properties{"富山県", "", "", "氷見市", "11111"},
			Geometry{"Polygon", [][]LatLng{{{136.111111, 36.111111}, {136.222222, 36.222222}}}}},
		},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestSplitPref(t *testing.T) {
	json := `
{
  "type": "FeatureCollection",
  "features": [
    {
      "type": "Feature",
      "properties": {
        "N03_001": "富山県",
        "N03_002": "",
        "N03_003": "",
        "N03_004": "氷見市",
        "N03_007": "11111"
		
      },
      "geometry": {
        "type": "Polygon",
        "coordinates": [
          [
            [
              136.111111,
              36.111111
            ],
            [
              136.222222,
              36.222222
            ]
          ]
        ]
      }
    }]
}`

	actual := Split(SPLIT_TYPE_PREF, []byte(json))
	expected := map[string][]Feature{
		"富山県": []Feature{{"Feature",
			Properties{"富山県", "", "", "氷見市", "11111"},
			Geometry{"Polygon", [][]LatLng{{{136.111111, 36.111111}, {136.222222, 36.222222}}}}},
		},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
