export interface SignUpSchema {
    Username: string;
    Password: string;
    Email: string;
    IsLoading: boolean;
    errors: string[];
}
