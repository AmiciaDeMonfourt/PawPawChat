import { createAsyncThunk } from '@reduxjs/toolkit';
import { $mockApi } from 'shared/config/axoisConfig/axiosConfig';
import { ChatInfo } from '../types/ChatInfoSchema';

export const fetchChatList = createAsyncThunk<ChatInfo[]>(
    'chats/fetchChatList',
    async (_, thunkApi) => {
        try {
            const response = await $mockApi.get<ChatInfo[]>('chats');
            return response.data;
        } catch (e) {
            thunkApi.rejectWithValue(e);
        }
    },
);
