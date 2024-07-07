import { AppDispatch } from 'app/providers/StoreProvider/config/store';
import { ChatList } from 'entities/ChatInfo';
import { getChats } from 'entities/ChatInfo/model/selectors/getChats';
import { fetchChatList } from 'entities/ChatInfo/model/service/getChats';
import { useEffect } from 'react';
import { useDispatch } from 'react-redux';
import { ContentWrapper } from 'shared/ui/ContentWrapper/ContentWrapper';
import { useTranslation } from 'react-i18next';

const ChatsPage = () => {
    const dispatch = useDispatch<AppDispatch>();
    const { t } = useTranslation();
    useEffect(() => {
        dispatch(fetchChatList());
    }, []);

    return (
        <ContentWrapper>
            <select>
                <option value="all">{t('All chats')}</option>
                <option value="unread">{t('Unread')}</option>
            </select>
            <ChatList />
        </ContentWrapper>
    );
};

export default ChatsPage;
