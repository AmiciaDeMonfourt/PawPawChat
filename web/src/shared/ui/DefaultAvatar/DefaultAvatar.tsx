import { classNames } from 'shared/lib/classNames/classNames';
import cls from './DefaultAvatar.module.scss';

interface DefaultAvatarProps {
    className?: string;
    name: string;
    size?: number;
}

export const DefaultAvatar = ({
    className,
    name,
    size = 200,
}: DefaultAvatarProps) => {
    return (
        <div
            className={classNames(cls.DefaultAvatar, {}, [className])}
            style={{
                width: `${size}px`,
                height: `${size}px`,
                fontSize: `${size / 3}px`,
            }}
        >
            {name[0]}
        </div>
    );
};
