import { PayloadAction, createSlice } from "@reduxjs/toolkit";
import { UserShema } from "../types/userSchema";
import { SignInResponse } from "features/SignIn/model/types/SignInResponse";

const initialState: UserShema = {
    UserData: undefined,
};

export const userSlice = createSlice({
    name: 'user',
    initialState,
    reducers: {
        setUserData: (state , action : PayloadAction<SignInResponse>) => {
            console.log(action);
            console.log("fff");
            const {Username, ID, Email} = action.payload.user;

            state.UserData.Email = "f";
            state.UserData.Username = "f";
            state.UserData.ID = 1;

            localStorage.setItem("token", action.payload.tokenString);
        }
    },
});

export const { actions: userActions} = userSlice;
export const { reducer: userReducer} = userSlice;