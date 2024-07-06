import { User } from 'entities/User';

export interface SignInResponse {
    user: User;
    tokenStr: string;
}
