import HtmlWebpackPlugin from "html-webpack-plugin";
import { BuildOptions, BuildPaths } from "../types/config";
import { ProgressPlugin, WebpackPluginInstance, HotModuleReplacementPlugin, DefinePlugin } from "webpack";
import  MiniCssExtractPlugin from "mini-css-extract-plugin";


HtmlWebpackPlugin
export function buildPlugins(options : BuildOptions): WebpackPluginInstance[] {

    const {paths, isDev} = options;

    return [
        new ProgressPlugin(),
        new HtmlWebpackPlugin({
            template: paths.html
        }),
        new MiniCssExtractPlugin({
            filename: 'css/[name].[contenthash:8].css',
            chunkFilename:'css/[name].[contenthash:8].css',
        }),
        new DefinePlugin({
            __IS_DEV__: JSON.stringify(isDev),
        }),
        new HotModuleReplacementPlugin(),
    ]
}