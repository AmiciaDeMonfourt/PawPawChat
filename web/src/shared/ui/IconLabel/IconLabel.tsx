import { classNames } from 'shared/lib/classNames/classNames';
import cls from './IconLabel.module.scss';
import { FC, ReactNode } from 'react';
import { Text } from '../Text/Text';

interface IconLabelProps {
    className?: string;
    icon: FC<React.SVGProps<SVGSVGElement>>;
    text: string;
}

export const IconLabel = ({ className, icon: Icon, text }: IconLabelProps) => {
    return (
        <div className={classNames(cls.IconLabel, {}, [className])}>
            <Icon className={cls.icon} />
            <Text>{text}</Text>
        </div>
    );
};

