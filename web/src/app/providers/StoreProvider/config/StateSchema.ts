import { ChatListSchema } from 'entities/ChatInfo';
import { UserSchema } from 'entities/User';
import { SignInSchema } from 'features/SignIn';
import { SignUpSchema } from 'features/SignUp';

export interface StateSchema {
    user: UserSchema;
    signIn: SignInSchema;
    signUp: SignUpSchema;
    chatList: ChatListSchema;
}
