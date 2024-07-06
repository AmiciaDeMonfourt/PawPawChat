import { createAsyncThunk } from '@reduxjs/toolkit';
import { SignInResponse } from '../types/SignInResponse';
import { getSignInFields } from '../selectors/getSignInFields/getSignInFields';
import { StateSchema } from 'app/providers/StoreProvider/config/StateSchema';
import { userActions } from 'entities/User';
import { $api } from 'shared/config/axoisConfig/axiosConfig';
import { AxiosError } from 'axios';

export const SignIn = createAsyncThunk<
    SignInResponse,
    void,
    { state: StateSchema }
>('signIn/userSignIn', async (_, thunkAPI) => {
    try {
        const signInData = getSignInFields(thunkAPI.getState());

        const response = await $api.post<SignInResponse>('signin', signInData);
        thunkAPI.dispatch(userActions.setUserData(response.data));
        return response.data;
    } catch (error) {
        if (error instanceof AxiosError) {
            thunkAPI.rejectWithValue(error.status);
        } else {
            thunkAPI.rejectWithValue(error);
        }
    }
});
