import { StateSchema } from 'app/providers/StoreProvider/config/StateSchema';

export const getUnreadChats = (state: StateSchema) => {
    state.chatList.chats.filter((data) => data.unread !== 0);
};
