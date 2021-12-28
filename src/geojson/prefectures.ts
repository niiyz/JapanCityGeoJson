import {pgClient} from "../pgClient";

const fs = require("fs");

const path = './geojson/prefectures';

if (!fs.existsSync(path)) {
    fs.mkdirSync(path, {recursive: true});
}

const dumpJsonSQL = `
SELECT json_build_object(
    'type', 'FeatureCollection',
    'features', json_agg(ST_AsGeoJSON(pref.*)::json)
    )
FROM (
	select
		code as id, name, ST_Union(geom)
	from
		prefectures
	where
	    code = $1
	group by
		code, name
) as pref;
`;

const writeReadme = (content: string) => {
    const readme = `| 都道府県 | 都道府県コード | GeoJson | TopoJson |
|-----------|--------------|------|------|
${content}`;

    fs.writeFileSync("geojson/prefectures/README.md", readme);
}

const main = async (): Promise<void> => {

    const client = await pgClient();

    const prefectures = await client.query('select code, name from prefectures group by code, name order by code');

    let mdContent = "";

    for (let i in prefectures.rows) {
        const code = prefectures.rows[i].code;
        const name = prefectures.rows[i].name;
        const json = await client.query(dumpJsonSQL, [code]);
        const filepath = `${path}/${code}.json`;
        console.log(code, name, filepath);
        fs.writeFileSync(filepath, JSON.stringify(json.rows[0].json_build_object));
        mdContent += `| ${name} | ${code} | [${name}](/geojson/prefectures/${code}.json) | [${name}](/topojson/prefectures/${code}.topojson) |\n`;
    }

    await client.end()

    writeReadme(mdContent);
}

main().then(() => {
    console.log(`${__filename} finished`);
});