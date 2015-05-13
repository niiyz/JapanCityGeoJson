# JapanCityGeoJson

国土交通省のデータから市町村の形を作るための緯度経度データを抽出しました。

国土数値情報 (JPGIS2.1(GML)準拠及びSHAPE形式データ)　国土交通省

http://nlftp.mlit.go.jp/ksj/gml/datalist/KsjTmplt-N03.html

~~~
ogr2ogr -f GeoJSON places.json N03-14_140401.shp
~~~

~~~
ruby geo_to_json.rb
~~~

# サンプル ex.GoogleMapAPIv3

富山県氷見市表示
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

# Demo

テキストボックスに県、市、町、村、郡まで入力してください。

http://japancityshape.niiyz.com/

有る。

・「富山県」

・「上新川郡」

・「上市町」

・「足立区」

無い

・「富山県氷見市」

