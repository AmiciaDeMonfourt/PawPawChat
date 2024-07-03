import { createAsyncThunk } from "@reduxjs/toolkit";
import { StateSchema } from "app/providers/StoreProvider/config/StateSchema";
import { userActions } from "entities/User";
import { getSignUpFields } from "../selectors/getSignUpFields/getSignUpFields";
import { SignUpResponse } from "../types/SignUpResponse";
import { $api } from "shared/config/axoisConfig/axiosConfig";
import axios, { AxiosError } from "axios";

export const SignUp = createAsyncThunk<SignUpResponse, void, {state : StateSchema}>(
    "signUp/userSignUp",
    async (_, thunkAPI) => {
        try {
            const signInData = getSignUpFields(thunkAPI.getState());
            
            const response = await $api.post<SignUpResponse>("signup", signInData);
            console.log(response);
            thunkAPI.dispatch(userActions.setUserData(response.data));

            return response.data;
        }
        catch(error) {
            if (error instanceof AxiosError) {
                thunkAPI.rejectWithValue(error.status);
            } else {
                thunkAPI.rejectWithValue(error);
            }
        }
    }
);