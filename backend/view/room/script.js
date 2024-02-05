
const displayRoundElement = document.getElementById('displayRound');
const displayCategoryElement = document.getElementById('displayCategory');
const displayQuestionElement = document.getElementById('displayQuestion');

const gameFormElement = document.getElementById('gameForm');

const readyButton = document.getElementById("readyButton");
const enterButton = document.getElementById("enterButton")
const connectButton = document.getElementById("connectButton")
const createRoomButton = document.getElementById("createRoomButton")

const entry = document.getElementById("entry")
const gameDashboard = document.getElementById("gameDashboard")
const entryDashboard = document.getElementById("entryDashboard")
const roomSettignsDashboard = document.getElementById("roomSettignsDashboard")
const roomDashboard = document.getElementById("roomDashboard")

const nameInput = document.getElementById("nameInput")

const answerAElement = document.getElementById('answerA');
const answerBElement = document.getElementById('answerB');
const answerCElement = document.getElementById('answerC');
const answerDElement = document.getElementById('answerD');

let conn;
let userName;

//room
let roomName;

let roomSettings = {
    roomName,
    difficulty: "łatwy",
    maxRound: 5,
    category: ""
}

let ready = true;
let isAnswerSent = false;

let gameState = {
    isGame: false,
    category: "",
    round: 1,
    question: "Test",
    answers: ["Lorem ipsum", "Lorem ipsum", "Lorem ipsum", "Lorem ipsum"],
    actions: [{ name: "", answer: null, round: 0 }],
    score: [{ name: "kuba", points: 10, roundsWon: [] }]
};

nameInput.addEventListener("change", (e) => {
    console.log(e.target.value)
    userName = e.target.value
})

createRoomButton.addEventListener("click", (e) => {
    console.log("test")
    roomSettignsDashboard.classList.remove("hidden")
})
gameFormElement.addEventListener("submit", (e) => {
    e.preventDefault();
    if (!ready) return alert("you are not ready")
    const formData = new FormData(gameFormElement);

    const answerValue = formData.get('q1');
    const answer = Number(answerValue)
    if (answerValue !== null) {
        sendAnswer(answer);
    } else {
        console.log("Nie wybrano odpowiedzi");
    }
});

// enterButton.addEventListener("click", (e) => {
//     e.preventDefault()
//     const input = document.getElementById("nameInput")
//     userName = input.value
//     connectWs()
//     return
// })

readyButton.addEventListener("click", (e) => {
    e.preventDefault()
    sendReadines()
    ready = true
})

//////////////////////////////////////////////////////////////////////////////

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
        case "update_gamestate":
            updateGameState(event)
            break
        case "update_players":
            updatePlayers(event)
            break
        case "room_message":
            console.log(event)
            break
        default:
            alert("unsupporting message type")
            break;
    }
}

function connectWs() {
    if (window.WebSocket) {
        conn = new WebSocket(`ws://localhost:8080/ws?room=room1&name=${userName}`)
        conn.onopen = (e) => {
            addQuery("room", "room1")
            entryDashboard.classList.add("hidden")
            gameDashboard.classList.remove("hidden")
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

function sendEvent(eventName, payload) {
    const event = new Event(eventName, payload)

    conn.send(JSON.stringify(event))
}

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

function updateTable(playerList) {
    const tableBody = document.getElementById('scoreTableBody');

    tableBody.innerHTML = '';

    playerList.forEach(player => {
        const row = document.createElement('tr');
        const nameCell = document.createElement('td');
        const pointsCell = document.createElement('td');

        nameCell.textContent = player.name;
        pointsCell.textContent = player.points;

        row.appendChild(nameCell);
        row.appendChild(pointsCell);
        tableBody.appendChild(row);
    });
}


///////////////////// SERVER EVENT FUNCTIONS //////////////////////

function startTheGame(event) {
    console.log("start", event)
    if (event.payload.isGame === true) {
        gameState = event.payload
        roomDashboard.classList.add("hidden")
        gameDashboard.classList.remove("hidden")
        updateDom()

    }
    return
}

function updateGameState(event) {
    gameState = event.payload
    console.log(gameState, "update")

    updateTable(gameState.score)
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


document.addEventListener("DOMContentLoaded", function () {
    // Cały twój kod JavaScript tutaj
});