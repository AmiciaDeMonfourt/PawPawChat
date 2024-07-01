declare module '*.scss' {
    interface IClassNames {
      [className: string]: string
    }
    const classes: IClassNames;
    export = classes; 
  }

declare module "*.svg" {
  import React from "react";
  const content: React.FC<React.SVGProps<SVGSVGElement>>;
  export default content;
}

declare module "*.png";
declare module "*.jpg";
declare module "*.jpeg";
declare module "*.ttf";

declare const __IS_DEV__: boolean;