import { useEffect, useState } from "react";
import { TRoomSettings } from "../types/event";
import { TUser } from "./useAppContext";

export const useWebSocket = ({ settings, user }: { settings: TRoomSettings, user: TUser }) => {
    let isNewGame = false;
    if (settings.name === "") {
        isNewGame = true
    }
    const urlNewGame = `http://127.0.0.1:8090/ws?room=${settings.name}&name=${user.name}&new=true&difficulty=${settings.difficulty}&maxRounds=${settings.maxRounds}&category=${settings.category}`
    const urlNotNewGame = `http://127.0.0.1:8090/room/?room=${settings.name}&name=${user.name}`
    console.log(urlNewGame, urlNotNewGame)
    const [socket, setSocket] = useState<WebSocket | null>(null);

    useEffect(() => {
        const ws = new WebSocket(settings.name);
        setSocket(ws);

        return () => {
            ws.close();
        };
    }, []);

    return socket;
};
