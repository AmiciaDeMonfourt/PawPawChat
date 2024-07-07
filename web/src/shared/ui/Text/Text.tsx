import { classNames } from 'shared/lib/classNames/classNames';
import cls from './Text.module.scss';
import { ReactNode } from 'react';

export enum TextSize {
    SMALL = 'small',
    MEDIUM = 'medium',
    LARGE = 'large',
    XL = 'xl',
}

interface TextProps {
    className?: string;
    children: ReactNode;
    textSize?: TextSize;
}

export const Text = ({
    className,
    children,
    textSize = TextSize.MEDIUM,
}: TextProps) => {
    return (
        <p className={classNames(cls.Text, {}, [className, textSize])}>
            {children}
        </p>
    );
};

