import { FC } from 'react';
import { classNames } from 'shared/lib/classNames/classNames';
import { Sidebar } from 'widgets/Sidebar/ui/Sidebar/Sidebar';

import cls from './ContentWrapper.module.scss';

export enum PageAlign {
    DEFAULT = 'default',
    CENTER = 'center',
    CENTER_HORIZONTAL = 'center_horizontal',
    CENTER_VERTICAL = 'center_vertical',
}

interface ContentWrapperProps {
    children?: React.ReactNode;
    fullscreen?: boolean;
    align?: PageAlign;
    className?: string;
}

export const ContentWrapper: FC<ContentWrapperProps> = (props) => {
    const {
        children,
        fullscreen = false,
        className,
        align = PageAlign.DEFAULT,
    } = props;

    return (
        <div className={classNames('content-wrapper', {}, [className])}>
            {!fullscreen ? <Sidebar /> : null}
            <div className={classNames('page-wrapper', {}, [cls[align]])}>
                {children}
            </div>
        </div>
    );
};
