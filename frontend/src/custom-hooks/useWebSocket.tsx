import { useEffect, useState, useContext } from "react";

import { TAppState, TUser, useAppStateContext } from "./useAppContext";
import { TRoomSettings } from "../types/event";

export interface IWSSettings {
    settings: TRoomSettings;
    user: TUser;
}

export const useWebSocket = () => {
    const { appState } = useAppStateContext();

    const settings = appState.settings;
    const user = appState.user;
    user.name = "test"
    const [socket, setSocket] = useState<WebSocket | null>(null);
    const [isConnected, setIsConnected] = useState(false);

    useEffect(() => {
        const createSocket = async () => {
            try {
                const isNewGame = settings.name === "";
                let url = `ws://127.0.0.1:8090/ws?room=${settings.name}&name=${user.name}`;

                if (isNewGame) {
                    url += `&new=true&difficulty=${settings.difficulty}&maxRounds=${settings.maxRounds}&category=${settings.category}`;
                }

                const ws = new WebSocket(url);

                ws.onopen = () => {
                    setIsConnected(true);
                    setSocket(ws);
                    console.log("WebSocket connection opened successfully!");
                };

                ws.onerror = (error) => {
                    console.error("WebSocket error:", error);
                };

                ws.onclose = () => {
                    setIsConnected(false);
                    setSocket(null);
                    console.log("WebSocket connection closed.");
                };

                // Cleanup function to close the WebSocket connection on unmount
                return () => ws.close();
            } catch (error) {
                console.error("Error creating WebSocket:", error);
            }
        };

        createSocket();
    }, [appState.settings, appState.user]);

    return { socket, isConnected, setSocket };
};
