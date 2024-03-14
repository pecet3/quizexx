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
    const [socket, setSocket] = useState<WebSocket | null>(null);
    const roomName = settings.roomName
    const userName = user.name

    function getWsUrl(isNewGame: boolean) {
        const baseUrl = "ws://127.0.0.1:8090/ws"

        console.log(settings, " sdadsa", user.name)
        if (isNewGame) {
            const gameSettings = {
                difficulty: settings.difficulty || '',
                maxRounds: settings.maxRounds || '',
                category: settings.category || '',
            }
            return `${baseUrl}?new=true&room=${roomName}&name=${userName}&difficulty=${gameSettings.difficulty}&maxRounds=${gameSettings.maxRounds}&category=${gameSettings.category}`
        } else {
            return `${baseUrl}?room=${roomName}&name=${userName}`
        }
    }


    const createSocket = async (isNewGame: boolean) => {
        if (settings.roomName === "" || user.name === "") {
            return
        }
        try {
            const url = await getWsUrl(isNewGame)
            console.log(url, " <=url")
            const ws = new WebSocket(url);

            ws.onopen = () => {
                setSocket(ws);
                console.log("WebSocket connection opened successfully!");
            };

            ws.onerror = (error) => {
                console.error("WebSocket error:", error);

            };

            ws.onclose = () => {

                setSocket(null);
                console.log("WebSocket connection closed.!!!!!!!!!");
            };

            // Cleanup function to close the WebSocket connection on unmount
            return () => ws.close();
        } catch (error) {
            console.error("Error creating WebSocket:", error);
        }
    };


    return { socket, createSocket };
};
