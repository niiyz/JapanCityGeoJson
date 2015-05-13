# JapanCityGeoJson

国土交通省のデータから市町村の形を作るためのGeoJsonデータを小分けしました。

国土数値情報 (JPGIS2.1(GML)準拠及びSHAPE形式データ)　国土交通省

国土交通省国土政策局GISHP http://nlftp.mlit.go.jp/ksj/gml/datalist/KsjTmplt-N03.html

GeoJson http://geojson.org/

~~~
ogr2ogr -f GeoJSON places.json N03-14_140401.shp
~~~

~~~
ruby geo_to_json.rb
~~~

# サンプル ex.GoogleMapAPIv3

GoogleMapAPIveはGeoJsonデータに対応しているので富山県氷見市を表示してみます。

~~~
window.onload = function() {
    var latLng = new google.maps.LatLng(36.786897, 136.892720);
    var mapOptions = {
        zoom: 9,
        center: latLng
    };
    var div = document.getElementById("canvas");
    div.style.width = div.style.height = '100%';
    var map = new google.maps.Map(div, mapOptions);
    map.data.loadGeoJson('http://japancityshape.niiyz.com/geojson/氷見市.json');
    map.data.setStyle({fillColor: 'orange'});
};
~~~

![Screencast](https://github.com/niiyz/JapanCityGeoJson/blob/master/screenshot.png)

~~~
map.data.loadGeoJson('http://japancityshape.niiyz.com/geojson/神奈川県.json');
~~~

~~~
map.data.loadGeoJson('http://japancityshape.niiyz.com/geojson/横浜市.json');
~~~

~~~
map.data.loadGeoJson('http://japancityshape.niiyz.com/geojson/金沢区.json');
map.data.loadGeoJson('http://japancityshape.niiyz.com/geojson/旭区.json');
~~~

# Demo

http://japancityshape.niiyz.com/
