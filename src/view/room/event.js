class Event {
    constructor(type, payload) {
        this.type = type
        this.payload = payload
    }
}
function routeEvent(event) {
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
            gameState = {}
            updateVirtualDom({
                entryDashboard: true,
                waitingRoomDashboard: false,
                gameDashboard: false,
            })
            break
        case "room_settings":
            updateRoomSettings(event)
            break
        case "players_answered":
            updatePlayersAnswered(event)
            break
        case "chat_message":
            console.log(event)
            updateDomChatMessages(event.payload)
            break
        default:
            alert("invalid event type: ", event.type)

            break;
    }
}

//////////////////// SERVER EVENT FUNCTIONS //////////////////////

function updateGameState(event) {
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

function updatePlayersAnswered(event) {
    const playersAnswered = event.payload
    updateDomScore(gameState.score, playersAnswered)
    return
}

function updatePlayers(event) {
    const newPlayersState = event.payload
    gameState.players = newPlayersState
    updateDomGameState()
    return
}

function updateReadyStatus(event) {
    const players = event.payload.clients
    updateDomReadyStatus(players)
    return
}

function updateRoomSettings(event) {
    const data = event.payload
    roomSettings = data
    updateDomSettings(data)
    return
}

function updateServerMessage(event) {
    const data = event.payload.message
    updateDomServerMessage(data)
    return
}

function sendEvent(eventName, payload) {
    const event = new Event(eventName, payload)
    conn.send(JSON.stringify(event))
    return
}

//////////////////// CLIENT EVENT FUNCTIONS ////////////////////

function sendReadines() {
    const payload = {
        name: userName,
        isReady: true,
    }
    sendEvent("ready_player", payload)
    return
}


function sendAnswer(answer) {
    if (isAnswerSent) {
        return
    }
    const payload = {
        name: userName,
        round: gameState.round,
        answer,
        points: 0,
    }
    sendEvent("send_answer", payload)
    isAnswerSent = true
    return
}

function sendChatMessage(message) {
    const payload = {
        name: userName,
        time: getCurrentDateTimeString(),
        message,
    }
    sendEvent("chat_message", payload)
    return
}