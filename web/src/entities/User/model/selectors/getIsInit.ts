import { StateSchema } from "app/providers/StoreProvider/config/StateSchema";

export const getIsInit = (state : StateSchema) : boolean => state.user.isInit;