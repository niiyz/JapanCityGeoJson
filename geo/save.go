package geo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func Save(path string, fts []Feature) {

	// Create FeatureCollection
	var fc = FeatureCollection{"FeatureCollection", fts}
	// Convert Struct To Json
	b, err := json.Marshal(&fc)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var out bytes.Buffer
	json.Indent(&out, b, "    ", "    ")

	// Create New File
	file, err := os.Create(path)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// File Close
	defer file.Close()

	// Wraite Bytes
	_, err = file.Write(b)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
