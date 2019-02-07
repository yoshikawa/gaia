const path = require("path");
const HtmlWebpackPlugin = require("html-webpack-plugin");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");

module.exports = {
  entry: {
    index: path.resolve(__dirname, "./src/index.tsx")
  },
  output: {
    path: path.resolve(__dirname, "./dist"),
    filename: "./js/[name].js"
  },
  plugins: [
    new HtmlWebpackPlugin({
      filename: "index.html",
      template: "src/index.html"
    }),
    new MiniCssExtractPlugin({
      filename: "./css/style.css"
    })
  ],
  resolve: {
    extensions: [
      ".ts", // for ts-loader
      ".tsx", // for ts-loader
      ".js",
      ".jsx"
    ]
  },
  module: {
    rules: [
      {
        test: /\.tsx?$/,
        use: "ts-loader"
      },
      {
        test: /\.scss$/,
        use: [MiniCssExtractPlugin.loader, "css-loader", "sass-loader"]
      },
      {
        test: /\.(jpg|png|gif)$/,
        use: {
          loader: "file-loader",
          options: {
            name: "./images/[name].[ext]",
            outputPath: "./",
            publicPath: path => "." + path
          }
        }
      },
      {
        test: /\.html$/,
        use: [
          {
            loader: "html-loader",
            options: { minimize: true }
          }
        ]
      }
    ]
  },
  devServer: {
    contentBase: "./dist",
    port: 8081,
    inline: true,
    host: "0.0.0.0"
  }
};
