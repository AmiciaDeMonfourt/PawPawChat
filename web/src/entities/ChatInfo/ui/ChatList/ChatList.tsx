import { classNames } from 'shared/lib/classNames/classNames';
import cls from './ChatList.module.scss';
import { useSelector } from 'react-redux';
import { getChats } from 'entities/ChatInfo/model/selectors/getChats';
import { ChatItem } from '../ChatItem/ChatItem';

interface ChatListProps {
    className?: string;
}

export const ChatList = ({ className }: ChatListProps) => {
    const chats = useSelector(getChats);

    return (
        <div className={classNames(cls.ChatList, {}, [className])}>
            {chats.map((data) => (
                <ChatItem key={data.id} chatInfo={data} />
            ))}
        </div>
    );
};
