import { StateSchema } from 'app/providers/StoreProvider/config/StateSchema';

export const getChats = (state: StateSchema) => state.chatList.chats;
