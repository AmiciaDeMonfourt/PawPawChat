import path from "path";
import { Configuration } from "webpack";
import { BuildEnv, BuildPaths } from "./config/webpack/types/config";
import { buildConfig } from "./config/webpack/builders/buildConfig";

export default (env : BuildEnv) : Configuration => {
    const paths : BuildPaths = {
        entry: path.resolve(__dirname,  'src', 'index.tsx'),
        output: path.resolve(__dirname, 'build'),
        html: path.resolve(__dirname, "public", "index.html"),
        src: path.resolve(__dirname, "src"),
    }
    
    const mode = env.mode || "development";
    const isDev = mode === "development";
    const port = env.port || 3000;
    
    return buildConfig({
        mode,
        paths,
        isDev,
        port
    });
}
