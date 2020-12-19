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
        "type": "MultiPolygon",
        "coordinates": [
          [
            [
              [
				136.111111,
              	36.111111
              ],
			  [
				136.222222,
                36.222222
			  ]
          	],
            [
              [
				136.333333,
              	36.333333
              ],
			  [
				136.444444,
                36.444444
			  ]
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
			Geometry{"MultiPolygon", [][][]LatLng{
				{
					{
						{136.111111, 36.111111}, {136.222222, 36.222222},
					},
					{
						{136.333333, 36.333333}, {136.444444, 36.444444},
					},
				},
			},
			},
		},
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
        "type": "MultiPolygon",
        "coordinates": [
          [
            [
              [
				136.111111,
              	36.111111
              ],
			  [
				136.222222,
                36.222222
			  ]
          	],
            [
              [
				136.333333,
              	36.333333
              ],
			  [
				136.444444,
                36.444444
			  ]
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
			Geometry{"MultiPolygon", [][][]LatLng{
						{
							{
								{136.111111, 36.111111}, {136.222222, 36.222222},
							},
							{
								{136.333333, 36.333333}, {136.444444, 36.444444},
							},
						},
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
