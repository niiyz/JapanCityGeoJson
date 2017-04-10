# JapanCityGeoJson 2016

47都道府県の県・市・町・村・郡・区の形を作るための[GeoJsonデータ](/geojson)、[TopoJsonデータ](/topojson)です。


国土数値情報 (JPGIS2.1(GML)準拠及びSHAPE形式データ)　国土交通省

国土交通省国土政策局GISHP http://nlftp.mlit.go.jp/ksj/gml/datalist/KsjTmplt-N03.html

GeoJson http://geojson.org/

TopoJson https://github.com/mbostock/topojson

# D3.js

D3.jsでtokyo23区のtopojsonを使用してみます。

![Screencast](https://github.com/niiyz/JapanCityGeoJson/blob/master/screenshot2.png)


~~~ html
<style>
    .tokyo23-13101 {fill: red;}
    .tokyo23-13102 {fill: blue;}
    .tokyo23-13103 {fill: green;}
    .tokyo23-13104 {fill: orange;}
    .tokyo23-13105 {fill: blueviolet;}
    .tokyo23-13106 {fill: azure;}
    .tokyo23-13107 {fill: forestgreen;}
    .tokyo23-13108 {fill: tomato;}
    .tokyo23-13109 {fill: lightyellow;}
    .tokyo23-13110 {fill: yellow;}
    .tokyo23-13111 {fill: crimson;}
    .tokyo23-13112 {fill: forestgreen;}
    .tokyo23-13113 {fill: red;}
    .tokyo23-13114 {fill: skyblue;}
    .tokyo23-13115 {fill: palegoldenrod;}
    .tokyo23-13116 {fill: red;}
    .tokyo23-13117 {fill: maroon;}
    .tokyo23-13118 {fill: royalblue;}
    .tokyo23-13119 {fill: lawngreen;}
    .tokyo23-13120 {fill: darkblue;}
    .tokyo23-13121 {fill: darkmagenta;}
    .tokyo23-13122 {fill: cornsilk;}
    .tokyo23-13123 {fill: aqua;}
</style>
<script src="http://d3js.org/d3.v3.min.js"></script>
<script src="http://d3js.org/topojson.v1.min.js"></script>
<script type="text/javascript" src="data/tokyo23_topojson.js"></script>
<script>
var width = 800, height = 500;
var svg = d3.select("body").append("svg")
    .attr("width", width)
    .attr("height", height);
var id = 'tokyo23';
var tokyo23 = topojson.feature(json, json.objects[id]);
var bounds = d3.geo.bounds(tokyo23);
var centerX = d3.sum(bounds, function(d) {return d[0];}) / 2,
  centerY = d3.sum(bounds, function(d) {return d[1];}) / 2;
var projection = d3.geo.mercator()
    .scale(70000)
    .center([centerX, centerY]);

svg.selectAll("path")
    .data(tokyo23.features).enter().append("path")
    .attr("d", d3.geo.path().projection(projection))
    .attr("class", function(d) { return 'tokyo23-' + d.id; });
</script>
~~~


### 市町村郡区シェイプ確認デモ(GoogleMap)

http://geojson.niiyz.com/

### データ更新手順
https://github.com/niiyz/JapanCityGeoJson/wiki/データ更新手順
