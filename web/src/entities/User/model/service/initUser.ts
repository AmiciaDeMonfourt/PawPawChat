import { createAsyncThunk } from '@reduxjs/toolkit';
import { $api } from 'shared/config/axoisConfig/axiosConfig';
import { LOCAL_STORAGE_TOKEN_KEY } from 'shared/consts/localstorage';
import { User } from '../types/userSchema';
import { InitUserResponse } from '../types/InitUserResponse';

export const initUser = createAsyncThunk<User>(
    'user/initUser',
    async (_, thunkApi) => {
        try {
            if (localStorage.getItem(LOCAL_STORAGE_TOKEN_KEY)) {
                const response = await $api.get<InitUserResponse>('api/user');
                return response.data.user;
            } else {
                thunkApi.rejectWithValue('');
            }
        } catch (e) {
            thunkApi.rejectWithValue('');
        }
    },
);
