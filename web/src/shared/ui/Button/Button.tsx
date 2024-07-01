import { classNames } from "shared/lib/classNames/classNames";
import cls from "./Button.module.scss";
import { ButtonHTMLAttributes, FC, ReactNode } from "react";


export enum ButtonTheme {
    CLEAR = "clear",
    CLASSIC = "classic"
}

export enum ButtonFont {
    SMALL = "small",
    MEDIUM = "medium",
    LARGE = "large",
    XL = "xl",
}

interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement>{
    className?: string;
    theme?: ButtonTheme;
    fontSize?: ButtonFont;
}

export const Button : FC<ButtonProps> = (props) => {
    const {
        className,
        children,
        theme = ButtonTheme.CLEAR,
        fontSize = ButtonFont.MEDIUM,
        ...otherProps
    } = props;

    return (
        <button
            className={classNames(cls.Button, {}, [className, cls[theme], cls[fontSize]])}
            {...otherProps}
        >
            {children}
        </button>
    )
}