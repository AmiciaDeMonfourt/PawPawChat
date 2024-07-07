import { classNames } from 'shared/lib/classNames/classNames';
import cls from './Text.module.scss';
import { ReactNode } from 'react';

export enum TextSize {
    SMALL = 'small',
    MEDIUM = 'medium',
    LARGE = 'large',
    XL = 'xl',
}

export enum TextTheme {
    TRANSPARENT = 'transparent',
}

interface TextProps {
    className?: string;
    children: ReactNode;
    textSize?: TextSize;
    textTheme?: TextTheme;
}

export const Text = ({
    className,
    children,
    textSize = TextSize.MEDIUM,
    textTheme,
}: TextProps) => {
    return (
        <p
            className={classNames(cls.Text, {}, [
                className,
                textSize,
                cls[textTheme],
            ])}
        >
            {children}
        </p>
    );
};

