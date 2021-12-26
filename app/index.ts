import { Client } from "pg";

const pgConnect = async (): Promise<Client> => {
    const client = new Client({
        host: "postgis",
        port: 5432,
        database: "postgis",
        user: "postgis",
        password: "postgis",
    })
    await client.connect()
    return client;
}

const main = async (): Promise<void> => {

    const client = await pgConnect();
    const sql = `
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
    const res = await client.query('select code, name from prefectures group by code, name order by code');

    const fs = require("fs");

    for (let i in res.rows) {
        const code = res.rows[i].code;
        const name = res.rows[i].name;
        console.log(code, name);
        const json = await client.query(sql, [code]);
        fs.writeFileSync(`${code}.json`, JSON.stringify(json.rows[0].json_build_object));
    }
    await client.end()
}

main().then(() => {
    console.log("index.ts finished");
});