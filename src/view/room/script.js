
const displayRoundElement = document.getElementById('displayRound');
const displayCategoryElement = document.getElementById('displayCategory');
const displayQuestionElement = document.getElementById('displayQuestion');

const gameFormElement = document.getElementById('gameForm');

const nameInput = document.getElementById("nameInput")
const connectButton = document.getElementById("connectButton")

const answerAElement = document.getElementById('answerA');
const answerBElement = document.getElementById('answerB');
const answerCElement = document.getElementById('answerC');
const answerDElement = document.getElementById('answerD');

const roomName = getRoomName();
let userName = "";


connectButton.addEventListener("click", (e) => {
    const nameInput = document.getElementById("nameInput")
    const name = nameInput.value
    userName = name

    if (name !== "" && roomName !== "") {
        connectWs()
    }
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
        const wsLink = getWsLink()
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

function getRoomName() {
    const queryString = window.location.search;
    const urlParams = new URLSearchParams(queryString);
    return urlParams.get('roomName') || '';
}

function getWsLink() {
    const gameSettings = getGameSettings();

    if (gameSettings) {
        return `ws://localhost:8080/ws?room=${roomName}&name=${userName}&difficulty=${gameSettings.difficulty}&maxRounds=${gameSettings.maxRounds}&category=${gameSettings.category}               `
    } else {
        return `ws://localhost:8080/ws?room=${roomName}&name=${userName}`
    }
}

function getGameSettings() {
    const queryString = window.location.search;
    const urlParams = new URLSearchParams(queryString);
    const isNewGame = urlParams.get('newGame') === 'true';

    if (isNewGame) {
        return {
            difficulty: urlParams.get('difficulty') || '',
            maxRounds: parseInt(urlParams.get('maxRounds')) || 0,
            category: urlParams.get('category') || '',
        };
    } else {
        return false;
    }
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