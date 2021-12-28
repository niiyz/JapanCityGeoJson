const path = require('path');

module.exports = {

    devServer: {
        static: {
            directory: path.join(__dirname, 'public'),
        },
        compress: true,
        port: 9000,
    },

    mode: 'development', // "production" | "development" | "none"

    entry: './src/index.ts',

    output: {
        path: path.join(__dirname, "public"),
        filename: "index.js"
    },

    module: {
        rules: [{
            test: /\.ts$/,
            use: 'ts-loader'
        }]
    },

    resolve: {
        modules: [
            "node_modules", // node_modules 内も対象とする
        ],
        extensions: [
            '.ts',
            '.js' // node_modulesのライブラリ読み込みに必要
        ]
    }
};