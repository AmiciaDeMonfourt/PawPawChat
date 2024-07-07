import {
    EnhancedStore,
    ReducersMapObject,
    configureStore,
} from '@reduxjs/toolkit';
import { StateSchema } from './StateSchema';
import { userReducer } from 'entities/User';
import { signInReducer } from 'features/SignIn';
import { signUpReducer } from 'features/SignUp';
import { chatListReducer } from 'entities/ChatInfo';

export function createReduxStore(initialState?: StateSchema) {
    const rootReducer: ReducersMapObject<StateSchema> = {
        user: userReducer,
        signIn: signInReducer,
        signUp: signUpReducer,
        chatList: chatListReducer,
    };

    return configureStore<StateSchema>({
        reducer: rootReducer,
        devTools: __IS_DEV__,
    });
}

export type AppDispatch = ReturnType<typeof createReduxStore>['dispatch'];
