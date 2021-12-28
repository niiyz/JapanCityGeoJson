import HOKKAIDO from "../../../geojson/prefectures/01.json";
import AOMORI from "../../../geojson/prefectures/02.json";

let map: google.maps.Map;

async function loadMap() {
    map = new google.maps.Map(document.getElementById("map") as HTMLElement, {
        center: { lat: 35.652832, lng: 139.839478 },
        zoom: 8,
    });
    console.log("loadMap");
}

loadMap().then(() => {
    map.data.addGeoJson(HOKKAIDO);
    map.data.addGeoJson(AOMORI);
    map.data.setStyle((feature) => {
        let fillColor;
        switch (feature.getProperty('name')) {
            case '北海道':
                fillColor = 'red';
                break;
            case '青森県':
                fillColor = 'blue';
                break;
        }
        return /** @type {google.maps.Data.StyleOptions} */ {
            fillColor,
            strokeWeight: 1,
        };
    });
});