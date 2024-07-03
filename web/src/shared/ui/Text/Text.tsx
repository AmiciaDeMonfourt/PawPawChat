import { classNames } from "shared/lib/classNames/classNames";
import cls from "./Text.module.scss";
import { ReactNode } from "react";

interface TextProps {
    className?: string;
    children: ReactNode
}

export const Text = ({className, children} : TextProps) => {
    return (
        <p className={classNames(cls.Text, {}, [className])}>
            {children}
        </p>
    )
}