# JapanCityGeoJson 2016

47都道府県の県・市・町・村・郡・区の形を作るための[GeoJsonデータ](/geojson)、[TopoJsonデータ](/topojson)です。

国土数値情報 (JPGIS2.1(GML)準拠及びSHAPE形式データ)　国土交通省
https://nlftp.mlit.go.jp/ksj/gml/datalist/KsjTmplt-N03-v2_4.html#prefecture00

GeoJson http://geojson.org/

TopoJson https://github.com/mbostock/topojson

```
docker-compose up

docker-compose exec app unzip N03-20200101_GML.zip N03-20_200101.geojson

docker-compose exec app go run main.go
```
