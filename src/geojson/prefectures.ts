import {pgClient} from "../pgClient";

const fs = require("fs");

const path = './geojson/prefectures';

if (!fs.existsSync(path)) {
    fs.mkdirSync(path, {recursive: true, mode: 0o777});
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

const dumpUnionJsonSQL = `
SELECT json_build_object(
    'type', 'FeatureCollection',
    'features', json_agg(ST_AsGeoJSON(pref.*)::json)
    )
FROM (
	select
	    code as id,
		ST_Collect(geom)
	from
		japan
	where
	    pref = $1
	group by
		pref, code
) as pref;
`;


const main = async (): Promise<void> => {

    const client = await pgClient();

    const prefectures = await client.query('select code, name from prefectures group by code, name order by code');

    for (let i in prefectures.rows) {
        const code = prefectures.rows[i].code;
        const name = prefectures.rows[i].name;
        const json = await client.query(dumpJsonSQL, [code]);
        const filepath = `${path}/${code}.json`;
        console.log(code, name, filepath);
        fs.writeFileSync(filepath, JSON.stringify(json.rows[0].json_build_object),{mode: 0o777});
        const json2 = await client.query(dumpUnionJsonSQL, [name]);
        const filepath2 = `${path}/${code}_union.json`;
        fs.writeFileSync(filepath2, JSON.stringify(json2.rows[0].json_build_object),{mode: 0o777});
    }

    await client.end()
}

main().then(() => {
    console.log(`${__filename} finished`);
});