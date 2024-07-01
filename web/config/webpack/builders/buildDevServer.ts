import { BuildOptions } from "../types/config";
import {Configuration as DevServerConfiguration} from "webpack-dev-server";

export function buildDevServer(options: BuildOptions) : DevServerConfiguration {
    const {
        port
    } = options;

    return {
          open: true,
          port,
          historyApiFallback: true,
          hot: true,
        //   proxy: [
        //     {
        //         context: ['/api'],
        //         target: 'http://localhost:8080',
        //         changeOrigin: true,
        //     },
        // ],
    }
}