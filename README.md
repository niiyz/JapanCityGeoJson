# JapanCityGeoJson 2016

47都道府県の県・市・町・村・郡・区の形を作るための[GeoJsonデータ](/geojson)、[TopoJsonデータ](/topojson)です。


国土数値情報 (JPGIS2.1(GML)準拠及びSHAPE形式データ)　国土交通省

国土交通省国土政策局GISHP http://nlftp.mlit.go.jp/ksj/gml/datalist/KsjTmplt-N03.html

GeoJson http://geojson.org/

TopoJson https://github.com/mbostock/topojson

# D3.js

D3.jsで京都のtopojsonを使用してみます。

![Screencast](https://github.com/niiyz/JapanCityGeoJson/blob/master/screenshot2.png)


~~~ html
<script src="http://d3js.org/d3.v3.min.js"></script>
<script src="http://d3js.org/topojson.v1.min.js"></script>
<script type="text/javascript" src="data/kyoto_pref_topo.js"></script>
<script>
var width = 800, height = 500;
var svg = d3.select("body").append("svg")
    .attr("width", width)
    .attr("height", height);
var kyoto = topojson.feature(json, json.objects["京都府"]);
var bounds = d3.geo.bounds(kyoto);
var centerX = d3.sum(bounds, function(d) {return d[0];}) / 2,
  centerY = d3.sum(bounds, function(d) {return d[1];}) / 2;
var projection = d3.geo.mercator()
    .scale(20000)
    .center([centerX, centerY]);

svg.selectAll("path")
    .data(kyoto.features).enter().append("path")
    .attr("d", d3.geo.path().projection(projection))
    .attr("class", function(d) { return d.id; });
</script>
~~~


### 市町村郡区シェイプ確認デモ(GoogleMap)

http://geojson.niiyz.com/
