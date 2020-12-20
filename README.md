# JapanCityGeoJson 2020

47都道府県の県・市・町・村・郡・区の形を作るための[GeoJsonデータ](/geojson)、[TopoJsonデータ](/topojson)です。

国土数値情報 (JPGIS2.1(GML)準拠及びSHAPE形式データ)　国土交通省
https://nlftp.mlit.go.jp/ksj/gml/datalist/KsjTmplt-N03-v2_4.html#prefecture00

GeoJson http://geojson.org/

https://tex2e.github.io/rfc-translater/html/rfc7946.html

TopoJson https://github.com/mbostock/topojson

### Make JAPAN2020-MULTIPOLYGON.json

```
// Download N03-20200101_GML.zip
// https://nlftp.mlit.go.jp/ksj/gml/datalist/KsjTmplt-N03-v2_4.html#prefecture00

// 
// Docker
docker-compose up

// Unzip Download File
docker-compose exec app unzip N03-20200101_GML.zip N03-20_200101.dbf N03-20_200101.shp N03-20_200101.shx N03-20_200101.xml

// Shapefile shit_jis -> UTF-8 , Input dataset open option shift_jis, Layer creation option UTF-8
docker-compose exec app ogr2ogr -f "ESRI Shapefile" -lco ENCODING=UTF-8 -oo ENCODING=shift_jis N03-20_200101sjis.shp N03-20_200101.shp

// Shapefile -> GeoJson
docker-compose exec app ogr2ogr -f GeoJSON -nlt MULTIPOLYGON JAPAN2020-MULTIPOLYGON.json N03-20_200101sjis.shp
```

### Make city json
```
// make city and pref, output to geojson directory
docker-compose exec app go run main.go JAPAN2020-MULTIPOLYGON.json
// make custom tokyo23, output geojson/custom/tokyo23.json
docker-compose exec app go run main.go JAPAN2020-MULTIPOLYGON.json custom tokyo23 13101 13102 13103 13104 13105 13106 13107 13108 13109 13110 13111 13112 13113 13114 13115 13116 13117 13118 13119 13120 13121 13122 13123
```
