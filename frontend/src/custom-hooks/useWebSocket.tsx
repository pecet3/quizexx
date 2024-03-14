import { useEffect, useState, useContext } from "react";

import { TAppState, TUser, useAppStateContext } from "./useAppContext";
import { IEvent, TRoomSettings } from "../types/event";

export interface IWSSettings {
    settings: TRoomSettings;
    user: TUser;
}

export const useWebSocket = () => {
    const { appState, setAppState } = useAppStateContext();
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

    function routeEvent(event: IEvent) {
        if (event.type === undefined) {
            alert("no type field in the event")
        }
        switch (event.type) {
            case "update_gamestate":
                updateGameState(event)
                break
            case "update_players":
                updatePlayers(event)
                break
            case "server_message":
                updateServerMessage(event)
                break
            case "ready_status":
                updateReadyStatus(event)
                break
            case "finish_game":

                break
            case "room_settings":
                updateRoomSettings(event)
                break
            default:
                alert("unsupporting message type")
                break;
        }
    }

    function updateGameState(event: IEvent) {
        if (gameState.isGame) {
            updateVirtualDom({
                entryDashboard: false,
                waitingRoomDashboard: false,
                gameDashboard: true,
            })
        }
        if (event.payload.round > gameState.round) {
            isAnswerSent = false
        }
        gameState = event.payload

        updateDomScore(gameState.score)
        updateDomGameState()
        return
    }

    function updatePlayers(event: IEvent) {
        const newPlayersState = event.payload
        gameState.players = newPlayersState
        updateDomGameState()

        return
    }

    function updateReadyStatus(event: IEvent) {
        const players = event.payload.clients
        updateDomReadyStatus(players)
    }

    function updateRoomSettings(event: IEvent) {
        const data = event.payload
        roomSettings = data
        updateDomSettings(data)
    }

    function updateServerMessage(event: IEvent) {
        const data = event.payload.message

        updateDomServerMessage(data)
    }

    function sendEvent(eventName, payload) {
        const event = new Event(eventName, payload)

        ws.send(JSON.stringify(event))
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

            ws.onmessage = (e) => {
                const event = JSON.parse(e.data) as IEvent
                routeEvent(event)
            }
            return () => ws.close();
        } catch (error) {
            console.error("Error creating WebSocket:", error);
        }
    };


    return { socket, createSocket };
};
