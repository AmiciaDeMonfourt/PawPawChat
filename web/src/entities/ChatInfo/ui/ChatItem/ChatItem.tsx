import { classNames } from 'shared/lib/classNames/classNames';
import cls from './ChatItem.module.scss';
import { ChatInfo } from '../../model/types/ChatInfoSchema';
import { Text, TextSize, TextTheme } from 'shared/ui/Text/Text';
import { trimMessage } from 'entities/ChatInfo/lib/trimMessage';
import { formatTimeAgo } from 'shared/lib/formatTime/formatTimeAgo';
import { AppLink } from 'shared/ui/AppLink/AppLink';

interface ChatItemProps {
    className?: string;
    chatInfo: ChatInfo;
}

export const ChatItem = ({ className, chatInfo }: ChatItemProps) => {
    return (
        <li className={classNames(cls.ChatItem, {}, [className])}>
            <AppLink to={`/chats/${chatInfo.id}`} className={cls.link}>
                <div className={cls.chatInfo}>
                    <div className={cls.avatar}></div>
                    <div className={cls.chatTitle}>
                        <Text textSize={TextSize.SMALL}>{chatInfo?.name}</Text>
                        <Text
                            textSize={TextSize.SMALL}
                            textTheme={TextTheme.TRANSPARENT}
                        >
                            {trimMessage(chatInfo?.last_message)}
                        </Text>
                    </div>
                </div>
                <div className={cls.rightBlock}>
                    <Text
                        textSize={TextSize.SMALL}
                        textTheme={TextTheme.TRANSPARENT}
                    >
                        {formatTimeAgo(Date.parse(chatInfo?.last_message_time))}
                    </Text>
                    {!!chatInfo.unread && (
                        <div className={cls.unread}>
                            <Text textSize={TextSize.SMALL}>
                                {chatInfo?.unread}
                            </Text>
                        </div>
                    )}
                </div>
            </AppLink>
        </li>
    );
};
