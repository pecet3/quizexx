
const displayRoundElement = document.getElementById('displayRound');
const displayQuestionElement = document.getElementById('displayQuestion');
const gameFormElement = document.getElementById('gameForm');
const answerAElement = document.querySelector('.answerA');
const answerBElement = document.querySelector('.answerB');
const answerCElement = document.querySelector('.answerC');
const answerDElement = document.querySelector('.answerD');
const readyButton = document.getElementById("readyButton");
const enterForm = document.getElementById("enterForm")

let conn;
let userName;

let ready = false;

enterForm.addEventListener("submit", (e) => {
    e.preventDefault()
    const input = document.getElementById("nameInput")
    userName = input.value
    connectWs()

})

gameFormElement.addEventListener("submit", (e) => {
    e.preventDefault();
    if (!ready) return alert("you are not ready")
    const formData = new FormData(gameFormElement);

    const answerValue = formData.get('q1');
    const answer = Number(answerValue)
    if (answerValue !== null) {
        console.log(answerValue);

        // Tutaj możesz użyć wartości answerValue
        sendAnswer(answer);
    } else {
        console.log("Nie wybrano odpowiedzi");
    }
});


readyButton.addEventListener("click", (e) => {
    e.preventDefault()
    sendReadines()
    ready = true
})


let gameState = {
    isGame: false,
    category: "",
    round: 1,
    question: "",
    answers: [""],
    actions: [{ name: "", answer: null, points: 0, round: 0 }],
    actionsHistory: [{ name: "", answer: null, points: 0, round: 0 }]
}


class Event {
    constructor(type, payload) {
        this.type = type
        this.payload = payload
    }
}

function routeEvent(event) {
    // if (event.type === undefined) {
    //     alert("no type field in the event")
    // }
    switch (event.type) {
        case "start_game":
            startTheGame(event)
            break;
        case "update_gamestate":
            updateGameState(event)
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
        conn = new WebSocket(`ws://localhost:8080/ws?room=room1&name=${userName}`)
        conn.onopen = (e) => {
            addQuery("room", "room1")
        }

        conn.onclose = (e) => {
            alert("closed connection with ws server ", e.data)
        }

        conn.onmessage = (e) => {
            console.log(e.data)
            const event = JSON.parse(e.data)
            routeEvent(event)
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

function updateDom() {
    answerAElement.innerHTML = gameState.answers[0]
    answerBElement.innerHTML = gameState.answers[1]
    answerCElement.innerHTML = gameState.answers[2]
    answerDElement.innerHTML = gameState.answers[3]

    displayRoundElement.innerHTML = gameState.round
    displayQuestionElement.innerHTML = gameState.question
}

///////////////////// SERVER EVENT FUNCTIONS //////////////////////

function startTheGame(event) {
    console.log("start", event)
    if (event.payload.isGame === true) {
        gameState = event.payload
        updateDom()
    }
    if (event.payload.isGame === false) {
        gameState.isGame = false
        gameState.round = 0
    }
    return
}

function updateGameState(event) {
    gameState = event.payload
    console.log(gameState, "update")
    updateDom()
    return
}

function updatePlayers(event) {
    const newPlayersState = event.payload
    gameState.players = newPlayersState
    updateDom()

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
    const payload = {
        name: userName,
        round: gameState.round,
        answer,
        points: 0,
    }
    sendEvent("send_answer", payload)
    return
}


