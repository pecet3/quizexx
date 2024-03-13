import React, { createContext, useContext, ReactNode, useState, useEffect } from "react";
import { TGameState, TRoomSettings } from "../types/event";
export type TUser = {
    name: string;
}
export type TAppState = {
    settings: TRoomSettings,
    gameState: TGameState,
    user: TUser,
};
export interface IAppStateProps {
    appState: TAppState;
    setAppState: React.Dispatch<React.SetStateAction<TAppState>>; // Adjust the type accordingly
}
const AppStateContext = createContext<IAppStateProps | undefined>(undefined);

export const AppStateProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
    const [appState, setAppState] = useState<TAppState>({
        settings: {
            roomName: "",
            category: "",
            difficulty: "easy",
            maxRounds: "5",
        },
        gameState: {
            round: 1,
            question: "",
            answers: [],
            actions: [],
            score: [],
            playersFinished: [],
        },
        user: {
            name: "",
        }
    });
    useEffect(() => {

    }, [appState.settings])
    return (
        <AppStateContext.Provider value={{ appState, setAppState }}>
            {children}
        </AppStateContext.Provider>
    );
};

export const useAppStateContext = (): IAppStateProps => {
    const context = useContext(AppStateContext);
    if (!context) {
        throw new Error("useAppStateContext must be used within an AppStateProvider");
    }
    return context;
};