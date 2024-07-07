import { classNames } from 'shared/lib/classNames/classNames';
import cls from './ChatList.module.scss';

interface ChatListProps {
    className?: string;
}

export const ChatList = ({ className }: ChatListProps) => {
    return <div className={classNames(cls.ChatList, {}, [className])}></div>;
};
