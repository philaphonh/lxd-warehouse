export interface User {
    id: number;
    username: string;
    firstname: string;
    lastname: string;
    email: string;
    role: string;
    imageUrl: string;
}

export interface UserRaw {
    id: number;
    username: string;
    firstname: string;
    lastname: string;
    email: string;
    roleId: number;
    imageUrl: string;
}

export interface UserRole {
    id: number;
    roleName: string;
}

export interface AuthData {
    userId: string;
    username: string;
    role: number;
    imageUrl: string;
    token: string;
}