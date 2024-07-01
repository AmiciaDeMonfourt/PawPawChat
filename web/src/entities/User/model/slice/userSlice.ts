import { PayloadAction, createSlice } from "@reduxjs/toolkit";
import { UserShema } from "../types/userSchema";
import { SignInResponse } from "features/SignIn/model/types/SignInResponse";

const initialState: UserShema = {};

export const userSlice = createSlice({
    name: 'user',
    initialState,
    reducers: {
        setUserData: (state , action : PayloadAction<SignInResponse>) => {
            const {Username, ID, Email} = action.payload.user;

            state.UserData.Email = Email;
            state.UserData.Username = Username;
            state.UserData.ID = ID;

            localStorage.setItem("token", action.payload.tokenString);
        }
    },
});

export const { actions: userActions} = userSlice;
export const { reducer: userReducer} = userSlice;