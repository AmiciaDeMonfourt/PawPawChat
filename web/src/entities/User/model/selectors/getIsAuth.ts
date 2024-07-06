import { StateSchema } from 'app/providers/StoreProvider/config/StateSchema';

export const getIsAuth = (state: StateSchema): boolean => state.user.isAuth;
