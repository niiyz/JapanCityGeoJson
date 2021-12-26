# JapanCityGeoJson 2020

47都道府県の県・市・町・村・郡・区の形を作るための[GeoJsonデータ](/geojson)、[TopoJsonデータ](/topojson)です。

国土数値情報 (JPGIS2.1(GML)準拠及びSHAPE形式データ)　国土交通省
https://nlftp.mlit.go.jp/ksj/gml/datalist/KsjTmplt-N03-v2_4.html#prefecture00

GeoJson http://geojson.org/

https://tex2e.github.io/rfc-translater/html/rfc7946.html

TopoJson https://github.com/mbostock/topojson

### Make JAPAN2020-MULTIPOLYGON.json

```shell
docker-compose up --build
```

### Make city json

- GeoJson
```shell
// make city and pref, output to geojson directory
docker-compose exec app go run main.go JAPAN2020-MULTIPOLYGON.json
// make custom tokyo23, output geojson/custom/tokyo23.json
docker-compose exec app go run main.go JAPAN2020-MULTIPOLYGON.json custom tokyo23 13101 13102 13103 13104 13105 13106 13107 13108 13109 13110 13111 13112 13113 13114 13115 13116 13117 13118 13119 13120 13121 13122 13123
```

- TopoJson
https://github.com/topojson/topojson/wiki/Introduction

```shell
docker-compose exec node /bin/bash geojson_to_topojson.sh
```