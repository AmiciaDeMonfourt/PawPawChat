import { StateSchema } from 'app/providers/StoreProvider/config/StateSchema';
import { SignUpRequest } from '../../types/SignUpRequest';

export const getSignUpFields = (state: StateSchema): SignUpRequest => {
    const { signUp } = state;
    const { Email, Password, Username } = signUp;

    return {
        Username,
        Email,
        Password,
    };
};
