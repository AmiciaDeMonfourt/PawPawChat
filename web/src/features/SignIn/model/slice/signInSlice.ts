import { PayloadAction, createSlice } from '@reduxjs/toolkit';
import { SignInSchema } from '../types/SignInSchema';
import { SignIn } from '../service/SignIn';

const initialState: SignInSchema = {
    Password: '',
    Email: '',
    IsLoading: false,
    errors: [],
};

export const userSlice = createSlice({
    name: 'signIn',
    initialState,
    reducers: {
        setPassword: (state, action: PayloadAction<string>) => {
            state.Password = action.payload;
        },
        setEmail: (state, action: PayloadAction<string>) => {
            state.Email = action.payload;
            console.log();
        },
    },
    extraReducers: (builder) => {
        builder
            .addCase(SignIn.pending, (state) => {
                state.IsLoading = true;
                console.log('ok');
            })
            .addCase(SignIn.fulfilled, (state) => {
                state.IsLoading = false;
            })
            .addCase(SignIn.rejected, (state, action) => {
                state.IsLoading = false;
                console.log(action);
            });
    },
});

export const { actions: signInActions } = userSlice;
export const { reducer: signInReducer } = userSlice;
