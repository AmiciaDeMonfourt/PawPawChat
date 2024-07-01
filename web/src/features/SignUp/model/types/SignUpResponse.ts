import { User } from "entities/User";

export interface SignUpResponse {
    user: User;
    tokenString: string;
}
