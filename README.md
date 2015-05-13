# JapanCityGeoJson

県、市町村郡区の形を作るためのGeoJsonデータです。

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

GoogleMapAPIv3はGeoJsonデータに対応しているので富山県氷見市を表示してみます。

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
    map.data.loadGeoJson('http://geojson.niiyz.com/geojson/富山県/氷見市.json');
    map.data.setStyle({fillColor: 'orange'});
};
~~~

![Screencast](https://github.com/niiyz/JapanCityGeoJson/blob/master/screenshot.png)


県
~~~ 
map.data.loadGeoJson('http://geojson.niiyz.com/geojson/神奈川県/神奈川県.json');
~~~

市
~~~
map.data.loadGeoJson('http://geojson.niiyz.com/geojson/神奈川県/横浜市.json');
~~~

区
~~~
map.data.loadGeoJson('http://geojson.niiyz.com/geojson/神奈川県/横浜市金沢区.json');
map.data.loadGeoJson('http://geojson.niiyz.com/geojson/神奈川県/横浜市旭区.json');
~~~

郡全域
~~~
map.data.loadGeoJson('http://geojson.niiyz.com/geojson/高知県/高岡郡.json');
~~~

町・村
~~~
map.data.loadGeoJson('http://geojson.niiyz.com/geojson/高知県/高岡郡中土佐.json');
~~~

# Demo

http://geojson.niiyz.com/
