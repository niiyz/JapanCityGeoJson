import {Client} from "pg";

export const pgClient = async (): Promise<Client> => {
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