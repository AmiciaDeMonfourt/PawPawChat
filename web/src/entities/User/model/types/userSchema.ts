export interface User {
    ID: number;
    Username: string;
    Email: string;
}

export interface UserShema {
    UserData?: User;
}