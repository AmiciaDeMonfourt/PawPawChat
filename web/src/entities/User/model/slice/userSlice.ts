import { PayloadAction, createSlice } from '@reduxjs/toolkit';
import { User, UserShema } from '../types/userSchema';
import { SignInResponse } from 'features/SignIn/model/types/SignInResponse';
import { LOCAL_STORAGE_TOKEN_KEY } from 'shared/consts/localstorage';
import { initUser } from '../service/initUser';

const initialState: UserShema = {
    UserData: null,
    isInit: false,
    isAuth: false,
};

export const userSlice = createSlice({
    name: 'user',
    initialState,
    reducers: {
        setUserData: (state, action: PayloadAction<SignInResponse>) => {
            console.log(action.payload);
            state.UserData = action.payload.user;

            if (state.UserData) {
                state.isAuth = true;
            }

            localStorage.setItem(
                LOCAL_STORAGE_TOKEN_KEY,
                action.payload.tokenStr,
            );
        },
        logout: (state) => {
            state.UserData = null;
            state.isAuth = false;
            localStorage.removeItem(LOCAL_STORAGE_TOKEN_KEY);
        },
    },
    extraReducers: (builder) => {
        builder
            .addCase(
                initUser.fulfilled,
                (state, action: PayloadAction<User>) => {
                    console.log(action);
                    state.isInit = true;
                    state.UserData = action.payload;

                    if (state.UserData) {
                        state.isAuth = true;
                    }
                },
            )
            .addCase(initUser.rejected, (state) => {
                state.isInit = true;
                state.isAuth = false;
                localStorage.removeItem(LOCAL_STORAGE_TOKEN_KEY);
            });
    },
});

export const { actions: userActions } = userSlice;
export const { reducer: userReducer } = userSlice;
