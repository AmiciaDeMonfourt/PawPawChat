import { Configuration } from "webpack";
import { BuildOptions } from "../types/config";
import { buildPlugins } from "./buildPlugins";
import { buildLoaders } from "./buildLoaders";
import { buildResolves } from "./buildResolves";
import { buildDevServer } from "./buildDevServer";

export function buildConfig(options : BuildOptions) : Configuration {
    const {mode, paths, isDev} = options;

    return {
        mode,
        entry: paths.entry,
        plugins: buildPlugins(options),
        module: {
            rules: buildLoaders(options),
        },
        devtool: isDev ? 'inline-source-map' : undefined,
        devServer: isDev ? buildDevServer(options) : undefined,
        resolve: buildResolves(options),
        output: {
            path: paths.output,
            filename: "[name][contenthash].bundle.js",
            clean: true,
        },
    }
}