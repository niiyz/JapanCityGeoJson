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
    const res = await client.query('SELECT json_build_object(\n' +
        '    \'type\', \'FeatureCollection\',\n' +
        '    \'features\', json_agg(ST_AsGeoJSON(pref.*)::json)\n' +
        '    )\n' +
        'FROM (\n' +
        '\tselect\n' +
        '\t\tcode as id, name, (ST_Dump(ST_Union(geom))).geom\n' +
        '\tfrom\n' +
        '\t\tprefectures\n' +
        '\twhere\n' +
        '\t\tcode = \'01\'\n' +
        '\tgroup by\n' +
        '\t\tcode, name\n' +
        ') as pref;');

    const fs = require("fs");

    try {
        fs.writeFileSync("test.json", JSON.stringify(res.rows[0].json_build_object));
    } catch(e: any){
        console.log(e.message);
    }

    await client.end()
}

main().then(() => {
    console.log("index.ts finished");
});