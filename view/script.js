
const displayRoundElement = document.getElementById('displayRound');
const displayQuestionElement = document.getElementById('displayQuestion');
const gameFormElement = document.getElementById('gameForm');
const answerAElement = document.querySelector('.answerA');
const answerBElement = document.querySelector('.answerB');
const answerCElement = document.querySelector('.answerC');
const answerDElement = document.querySelector('.answerD');
let conn;
let userName = "tester"
connectWs()

let gameState = {
    isGame: false,
    category: "",
    round: 0,
    question: "",
    players: [{ name: "", answer: null, points: 0 }],
    prevRoundWinner: [""]
}


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
        case "start_game":
            startTheGame(event)
            break;
        case "new_round":
            startNewRound(event)
            break
        case "update_players":
            updatePlayers(event)
            break
        default:
            alert("unsupporting message type")
            break;
    }
}

function sendEvent(eventName, payload) {
    const event = new Event(eventName, payload)

    conn.send(JSON.stringify(event))
}


function connectWs() {
    if (window.WebSocket) {
        conn = new WebSocket(`ws://localhost:8080/ws?room=room1&name=tester`)
        conn.onopen = (e) => {
            addQuery("room", "room1")
            sendReadines()
        }

        conn.onclose = (e) => {
            alert("closed connection with ws server ")
        }

        conn.onmessage = (e) => {
            console.log(e.data)
            routeEvent(e)
        }
    } else {
        alert("your browser doesn't support websockets")
    }
}
//////////////////////////////////////////

function addQuery(param, value) {
    const url = new URL(window.location.href)
    url.searchParams.set(param, value)
    history.replaceState(null, null, url.href)
    return
}

///////////////////// SERVER EVENT FUNCTIONS //////////////////////

function startTheGame(event) {
    console.log(event)
    if (event.payload.isGame === true) {
        gameState.isGame = true
        gameState.round = event.payload.round
    }
    if (event.payload.isGame === false) {
        gameState.isGame = false
        gameState.round = 0
    }
    return
}

function startNewRound(event) {
    console.log(event)
    const newRound = event.payload.round
    if (!isGame) return
    if ((newRound - 1) !== round) return
    round = newRound
    return
}

function updatePlayers(event) {
    console.log(event)
    const newPlayersState = event.payload
    gameState.players = newPlayersState
}

//////////////////// CLIENT EVENT FUNCTIONS ////////////////////

function sendReadines() {
    const payload = {
        userName,
        isReady: true,
    }
    console.log(conn)
    conn.send(JSON.stringify({ type: "ready_player", payload }))
}

function sendAnswer() {
    const payload = {
        name,
        round,
        answer: 1,
    }
    conn.send(JSON.stringify({ type: "send_answer", payload }))
}


