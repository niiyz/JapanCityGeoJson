import HOKKAIDO from "../../../geojson/prefectures/01.json";
import AOMORI from "../../../geojson/prefectures/02.json";
import TOYAMA from "../../../geojson/prefectures/16.json";

let map: google.maps.Map;

async function loadMap() {
    map = new google.maps.Map(document.getElementById("map") as HTMLElement, {
        center: { lat: 43.06417, lng: 141.34695 },
        zoom: 6,
    });
    console.log("loadMap");
}

loadMap().then(() => {
    map.data.addGeoJson(HOKKAIDO);
    map.data.addGeoJson(TOYAMA);
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