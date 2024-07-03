export interface User {
    id: number;
    username: string;
    email: string;
}

export interface UserShema {
    UserData?: User;
    isInit: boolean;
    isAuth: boolean;
}