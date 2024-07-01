import { User } from "entities/User";

export interface SignInResponse {
    user: User;
    tokenString: string;
}