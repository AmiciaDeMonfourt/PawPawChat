import {StateSchema} from "app/providers/StoreProvider/config/StateSchema"
import {SignInRequest} from "../../types/SignInRequest"

export const getSignInFields = (state : StateSchema) : SignInRequest  => {
    const {signIn} = state;
    const {Email, Password} = signIn;

    return {
        Email,
        Password,
    }
}