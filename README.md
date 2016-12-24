# JapanCityGeoJson 2016

47都道府県の県・市・町・村・郡・区の形を作るための[GeoJsonデータ](/geojson)、[TopoJsonデータ](/topojson)です。

# サンプル 例:GoogleMapAPIv3

GoogleMapAPIv3はGeoJsonデータに対応しているので富山県氷見市を表示してみます。

## ローカルに設置したGeoJsonを追加

~~~ js
var data = {<-- GeoJson -->};
var mapOptions = {
    zoom: 9,
    center: new google.maps.LatLng(36.786897, 136.892720)
};
var map = new google.maps.Map(document.getElementById("canvas"), mapOptions);
map.data.addGeoJson(data);
map.data.setStyle({fillColor: 'green'});
~~~

![Screencast](https://github.com/niiyz/JapanCityGeoJson/blob/master/screenshot.png)


### 市町村郡区シェイプ確認デモ(GoogleMap)

http://geojson.niiyz.com/

# サンプル 例: D3.js

TopoJsonをD3.jsで使用してみます。

~~~ html
<style>
    .氷見市 { fill: red;}
</style>
<script src="http://d3js.org/d3.v3.min.js"></script>
<script src="http://d3js.org/topojson.v1.min.js"></script>
<script>
var width = 800, height = 500;
var svg = d3.select("body").append("svg")
    .attr("width", width)
    .attr("height", height);

d3.json("topojson/富山県/富山県.topojson", function(error, json) {

  var toyama = topojson.feature(json, json.objects["富山県"]);
  var bounds = d3.geo.bounds(toyama);
  var centerX = d3.sum(bounds, function(d) {return d[0];}) / 2,
      centerY = d3.sum(bounds, function(d) {return d[1];}) / 2;
  var projection = d3.geo.mercator()
    .scale(10000)
    .center([centerX, centerY]);

  svg.selectAll("path")
      .data(toyama.features).enter().append("path")
      .attr("d", d3.geo.path().projection(projection))
      .attr("class", function(d) { return d.id; });
});
</script>
~~~

![Screencast](https://github.com/niiyz/JapanCityGeoJson/blob/master/screenshot2.png)

国土数値情報 (JPGIS2.1(GML)準拠及びSHAPE形式データ)　国土交通省

国土交通省国土政策局GISHP http://nlftp.mlit.go.jp/ksj/gml/datalist/KsjTmplt-N03.html

GeoJson http://geojson.org/

TopoJson https://github.com/mbostock/topojson
