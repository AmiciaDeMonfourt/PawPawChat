import { StateSchema } from "app/providers/StoreProvider/config/StateSchema";
import { User } from "../model/types/userSchema";

export const getUser = (state : StateSchema) : User => state.user.UserData; 