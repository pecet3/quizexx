import { useEffect, useState } from "react";

import { TAppState, TUser, useAppStateContext } from "./useAppContext";
import { TRoomSettings } from "../types/event";

export interface IWSSettings {
    settings: TRoomSettings,
    user: TUser
}
export const useWebSocket = () => {
    const { appState } = useAppStateContext();

    const settings = appState.settings
    const user = appState.user
    console.log("ws hook: ", appState)
    let isNewGame = false;
    if (settings.name === "") {
        isNewGame = true
    }
    const urlNewGame = `http://127.0.0.1:8090/ws?room=${settings.name}&name=${user.name}&new=true&difficulty=${settings.difficulty}&maxRounds=${settings.maxRounds}&category=${settings.category}`
    const urlNotNewGame = `http://127.0.0.1:8090/room/?room=${settings.name}&name=${user.name}`
    console.log(urlNewGame, urlNotNewGame)
    const [socket, setSocket] = useState<WebSocket | null>(null);

    let url = urlNotNewGame
    if (isNewGame) {
        url = urlNewGame
    }
    // useEffect(() => {
    //     const ws = new WebSocket(settings.name);
    //     setSocket(ws);

    //     return () => {
    //         ws.close();
    //     };
    // }, []);

    return socket;
};
