export interface User {
    id: number;
    username: string;
    email: string;
}

export interface UserSchema {
    UserData?: User;
    isInit: boolean;
    isAuth: boolean;
}

