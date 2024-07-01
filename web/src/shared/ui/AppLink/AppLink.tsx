import { classNames } from "shared/lib/classNames/classNames";
import cls from "./AppLink.module.scss";
import { LinkProps } from "react-router-dom";
import { FC } from "react";
import { Link } from "react-router-dom";

export enum AppLinkTheme {
    CLASSIC = "classic",
}

interface AppLinkProps extends LinkProps {
    className?: string,
    theme?: AppLinkTheme,
}

export const AppLink: FC<AppLinkProps> = (props : AppLinkProps) => {
    const {
        className,
        to,
        children,
        theme = AppLinkTheme.CLASSIC,
        ...otherProps
    } = props;

    return (
        <Link 
            className={classNames(cls.AppLink, {}, [className, cls[theme]])}
            to={to}
            {...otherProps}
        >
            {children}
        </Link>
    )
}