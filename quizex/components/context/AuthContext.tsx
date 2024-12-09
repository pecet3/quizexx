
// context/AuthContext.tsx
import React, { createContext, useContext, useState, ReactNode } from 'react';
import { User } from '../types/auth';

const AuthContext = createContext<AuthContextType | null>(null);

interface AuthProviderProps {
    children: ReactNode;
}
export interface LoginResponse {
    user: User;
    token: string;
}

export interface AuthContextType {
    user: User | null;
    jwt: string | null;
    setUser: (u: User) => void;
    setJWT: (token: string) => void;
}



export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
    const [user, setUser] = useState<User | null>(null);
    const [jwt, setJWT] = useState<string | null>(null);

    const login = (userData: User, token: string): void => {
        setUser(userData);
        setJWT(token);
    };

    const logout = () => {
        setUser(null);
        setJWT(null);
    };

    const updateUser = (userData: Partial<User>) => {
        setUser((prevUser) => {
            if (!prevUser) return null;
            return {
                ...prevUser,
                ...userData,
            };
        });
    };

    const value: AuthContextType = {
        user,
        jwt,
        setUser,
        setJWT,
    };

    return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};

// Dodajemy osobną funkcję do obsługi logiki logowani
export const useAuth = (): AuthContextType => {
    const context = useContext(AuthContext);
    if (!context) {
        throw new Error('useAuth must be used within an AuthProvider');
    }
    return context;
};