import { PayloadAction, createSlice } from "@reduxjs/toolkit";
import { SignUpSchema } from "../types/SignUpSchema";
import { SignUp } from "../service/SignUp";
import { error } from "console";

const initialState: SignUpSchema = {
    Username: '',
    Password: '',
    Email: '',
    IsLoading: false,
    errors: [],
};

export const userSlice = createSlice({
    name: 'signUp',
    initialState,
    reducers: {
        setUsername: (state, action : PayloadAction<string>) => {
            state.Username = action.payload;
        },
        setPassword: (state, action : PayloadAction<string>) => {
            state.Password = action.payload;
        },
        setEmail: (state, action : PayloadAction<string>) => {
            state.Email = action.payload;
        }
    },
    extraReducers: (builder) => {
        builder
        .addCase(SignUp.pending, (state) => {
            state.IsLoading = true;
            console.log("ok");
        })
        .addCase(SignUp.fulfilled, (state) => {
            state.IsLoading = false;
        })
        .addCase(SignUp.rejected, (state, action) => {
            state.IsLoading = false;
            console.log(action);
        })
    }
});

export const { actions: signUpActions} = userSlice;
export const { reducer: signUpReducer} = userSlice;