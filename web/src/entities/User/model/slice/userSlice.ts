import { PayloadAction, createSlice } from "@reduxjs/toolkit";
import { UserShema } from "../types/userSchema";
import { SignInResponse } from "features/SignIn/model/types/SignInResponse";
import { LOCAL_STORAGE_TOKEN_KEY } from "shared/consts/localstorage";

const initialState: UserShema = {
    UserData: null,
    isInit: false
};

export const userSlice = createSlice({
    name: 'user',
    initialState,
    reducers: {
        setUserData: (state , action : PayloadAction<SignInResponse>) => {
            state.UserData = action.payload.user;
            localStorage.setItem(LOCAL_STORAGE_TOKEN_KEY, action.payload.tokenString);
        }
    },
});

export const { actions: userActions} = userSlice;
export const { reducer: userReducer} = userSlice;