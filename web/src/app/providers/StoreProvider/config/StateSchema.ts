import { UserShema } from "entities/User";
import { SignInSchema } from "features/SignIn";
import { SignUpSchema } from "features/SignUp";

export interface StateSchema {
    user: UserShema;
    signIn: SignInSchema;
    signUp: SignUpSchema;
}