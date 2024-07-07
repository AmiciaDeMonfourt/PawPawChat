import { PayloadAction, createSlice } from '@reduxjs/toolkit';
import { ChatListSchema } from '../types/ChatInfoSchema';
import { fetchChatList } from '../service/getChats';

const initialState: ChatListSchema = {
    chats: [],
    isLoading: false,
};

export const chatListSlice = createSlice({
    name: 'chatList',
    initialState,
    reducers: {
        clearChats: (state) => {
            state.chats = [];
        },
    },
    extraReducers: (builder) => {
        builder
            .addCase(fetchChatList.pending, (state) => {
                state.isLoading = true;
            })
            .addCase(fetchChatList.fulfilled, (state, action) => {
                state.isLoading = false;
                state.chats = action.payload;
            })
            .addCase(fetchChatList.rejected, (state, action) => {
                state.isLoading = false;
            });
    },
});

export const { actions: chatListActions } = chatListSlice;
export const { reducer: chatListReducer } = chatListSlice;
