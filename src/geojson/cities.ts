import {pgClient} from "../pgClient";

const fs = require("fs");

const mdOnly = false;

const path = './geojson/cities';

const dumpJsonSQL = `
SELECT json_build_object(
    'type', 'FeatureCollection',
    'features', json_agg(ST_AsGeoJSON(pref.*)::json)
    )
FROM (
	select
		code as id, ST_Union(geom)
	from
		japan
	where
	    code = $1
    and
        ST_Area(geom) > 0.0003495285322129
	group by
		code
) as pref;
`;

const writeReadme = (prefCode: string, content: string) => {
    const readme = `| 都道府県 | 行政区分 | 行政区分コード | GeoJson | TopoJson |
|-----------|--------- |--------------|------|------|
${content}`;

    fs.writeFileSync(`geojson/cities/${prefCode}/README.md`, readme);
    fs.writeFileSync(`topojson/cities/${prefCode}/README.md`, readme);
}


const main = async (): Promise<void> => {

    const client = await pgClient();

    const cities = await client.query("select code, pref, regional, city1, city2, count(*) as cnt from japan where ST_Area(geom) > 0.0003495285322129 group by code, pref, regional, city1, city2 order by code");

    let mdContent = "";
    let prevPrefCode = "";

    for (let i in cities.rows) {
        const city = cities.rows[i];
        if (!city.code) {
             console.log("所属未定地");
             break;
        }
        const prefCode = city.code.substring(0, 2);
        if (prevPrefCode === "") {
            prevPrefCode = prefCode;
        }
        const filepath = `${path}/${prefCode}`;
        if (!fs.existsSync(filepath)) {
            fs.mkdirSync(filepath, {recursive: true});
        }
        if (! mdOnly) {
            const json = await client.query(dumpJsonSQL, [city.code]);
            fs.writeFileSync(`${filepath}/${city.code}.json`, JSON.stringify(json.rows[0].json_build_object));
        }
        console.log(city.code, city.pref, city.regional, city.city1, city.city2, city.cnt);
        if (prevPrefCode !== prefCode) {
            writeReadme(prevPrefCode, mdContent);
            prevPrefCode = prefCode;
            mdContent = "";
        }
        const cityName = `${city.regional || ""}${city.city1 || ""}${city.city2 || ""}`;
        mdContent += `| ${city.pref} | ${cityName} | ${city.code} | [${cityName}](/geojson/cities/${prefCode}/${city.code}.json) | [${cityName}](/topojson/cities/${prefCode}/${city.code}.topojson) |\n`;
    }

    writeReadme(prevPrefCode, mdContent);

    await client.end();
}

main().then(() => {
    console.log(`${__filename} finished`);
});