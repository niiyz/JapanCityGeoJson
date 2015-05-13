# JapanCityShape

国土交通省のデータから市町村の形を作るための緯度経度データを抽出しました。

国土数値情報 (JPGIS2.1(GML)準拠及びSHAPE形式データ)　国土交通省

http://nlftp.mlit.go.jp/ksj/gml/datalist/KsjTmplt-N03.html

~~~
ogr2ogr -f GeoJSON places.json N03-14_140401.shp
~~~

~~~
ruby geo_to_json.rb
~~~

# Sample ex.GoogleMapAPIv3 Polygon

~~~
$.getJSON('data/富山県.json', function (data) {
    for (var city in data) {
        var latLngs = data[city];
        var coords = [];
        for (var i = 0; i < latLngs.length; i++) {
            var lat = latLngs[i][1];
            var lng = latLngs[i][0];
            coords.push(new google.maps.LatLng(lat, lng));
        }
        var polygon = new google.maps.Polygon({
            paths: coords,
            strokeColor: "red",
            strokeOpacity: 0.8,
            strokeWeight: 2,
            fillColor: 'orange',
            fillOpacity: 0.4
        });
        polygon.setMap(map);
    }
});
~~~

![Screencast](https://github.com/niiyz/JapanCityShape/blob/master/screenshot.png)

# Demo

http://toyamamap.niiyz.com/
