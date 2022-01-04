import {pgClient} from "../pgClient";

const fs = require("fs");

const path = './topojson/prefectures';

if (!fs.existsSync(path)) {
    fs.mkdirSync(path, {recursive: true, mode: 0o777});


const main = async (): Promise<void> => {

    const client = await pgClient();

    const prefectures = await client.query("select substring(code, 1, 2) as code, pref from japan where pref = '富山県' group by pref, substring(code, 1, 2)");

    for (let i in prefectures.rows) {
        const prefecture = prefectures.rows[i];
        const json = await client.query('select gid as arc_id, geom, code from japan where pref = $1 group by code', [prefecture.code]);
        console.log(prefecture.code, prefecture.name);
        fs.writeFileSync(`${path}/${prefecture.code}.topojson`, JSON.stringify(JSON.parse(json.rows[0].topojson)),{mode: 0o777});
    }
    await client.end();
}

main().then(() => {
    console.log(`${__filename} finished`);
});