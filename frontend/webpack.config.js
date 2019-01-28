const path = require("path");
const HtmlWebpackPlugin = require("html-webpack-plugin");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");

module.exports = {
  entry: path.resolve(__dirname, "./src/main.tsx"),
  output: {
    path: path.resolve(__dirname, "./dist"),
    filename: "./js/main.js"
  },
  plugins: [
    new HtmlWebpackPlugin({
      template: "src/index.html"
    }),
    new HtmlWebpackPlugin({
      filename: "login.html",
      template: "src/view/login.html"
    }),
    new HtmlWebpackPlugin({
      filename: "users/index.html",
      template: "src/view/users/index.html"
    }),
    new HtmlWebpackPlugin({
      filename: "users/add.html",
      template: "src/view/users/add.html"
    }),
    new HtmlWebpackPlugin({
      filename: "users/edit.html",
      template: "src/view/users/edit.html"
    }),
    new HtmlWebpackPlugin({
      filename: "users/view.html",
      template: "src/view/users/view.html"
    }),
    new HtmlWebpackPlugin({
      filename: "units/index.html",
      template: "src/view/units/index.html"
    }),
    new HtmlWebpackPlugin({
      filename: "units/add.html",
      template: "src/view/units/add.html"
    }),
    new HtmlWebpackPlugin({
      filename: "units/edit.html",
      template: "src/view/units/edit.html"
    }),
    new HtmlWebpackPlugin({
      filename: "units/view.html",
      template: "src/view/units/view.html"
    }),
    new MiniCssExtractPlugin({
      filename: "./stylesheets/style.scss"
    })
  ],
  resolve: {
    extensions: [
      ".ts", // for ts-loader
      ".tsx", // for ts-loader
      ".js"
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
