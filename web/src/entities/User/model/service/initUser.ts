import { createAsyncThunk } from "@reduxjs/toolkit";
import { $api } from "shared/config/axoisConfig/axiosConfig";
import { LOCAL_STORAGE_TOKEN_KEY } from "shared/consts/localstorage";
import { User } from "../types/userSchema";

export const initUser = createAsyncThunk(
    "user/initUser",
    async (_, thunkApi) => {
        try {
            if (localStorage.getItem(LOCAL_STORAGE_TOKEN_KEY)) {
                $api.get("");
            }
            else {
                thunkApi.rejectWithValue("");
            }
        }
        catch (e) {
            thunkApi.rejectWithValue("");
        }
    }
);