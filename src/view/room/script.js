
const gameForm = document.getElementById('gameForm');
const connectButton = document.getElementById("connectButton")
const readyButton = document.getElementById("readyButton")
const readyUsersList = document.getElementById('readyUsersList')
const answerAElement = document.getElementById('answerA');
const answerBElement = document.getElementById('answerB');
const answerCElement = document.getElementById('answerC');
const answerDElement = document.getElementById('answerD');


const displayRoundElement = document.getElementById('displayRound');
const displayCategoryElement = document.getElementById('displayCategory');
const displayQuestionElement = document.getElementById('displayQuestion');
const displayPlayers = document.getElementById('displayPlayersInGame')
const displayReadyCount = document.getElementById('displayReadyCount')
const displayServerMessage = document.getElementById('displayServerMessage')

const roomName = getRoomName();
let userName = "";

let ready = true;
let isAnswerSent = false;
let isPlayerReady = false;

let gameState = {
    isGame: false,
    round: 1,
    question: "Test",
    answers: ["Lorem ipsum", "Lorem ipsum", "Lorem ipsum", "Lorem ipsum"],
    actions: [{ name: "", answer: null, round: 0 }],
    score: [{ name: "", points: 0, roundsWon: [] }],
};

const entryDashboard = document.getElementById("entryDashboard")
const waitingRoomDashboard = document.getElementById("waitingRoomDashboard")
const gameDashboard = document.getElementById("gameDashboard")

let virtualDom = {
    entryDashboard: true,
    waitingRoomDashboard: false,
    gameDashboard: false,
}

handleVirtualDom()
//////////////// Listeners /////////

connectButton.addEventListener("click", () => {
    const nameInput = document.getElementById("userNameInput")
    const name = nameInput.value
    userName = name

    if (name !== "" && roomName !== "") {
        connectWs()
    }
})
readyButton.addEventListener("click", () => {
    if (isPlayerReady === true) return
    sendReadines()
    readyButton.disabled = true
    return
})

gameForm.addEventListener("submit", (e) => {
    e.preventDefault();
    if (!ready) return alert("you are not ready")
    let formData = new FormData(gameForm);

    //// to fix reset value in the future

    let answerValue = formData.get('q1');
    const answer = Number(answerValue)
    if (answerValue !== null) {
        sendAnswer(answer);
        formData = null
        return
    } else {
        console.log("Nie wybrano odpowiedzi");
    }
});

//////////////////////////////////////////////////////////////////////////////

class Event {
    constructor(type, payload) {
        this.type = type
        this.payload = payload
    }
}
function routeEvent(event) {
    console.log(event.type)
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
            console.log(event)
            updateRoomSettings(event)
            break
        default:
            alert("unsupporting message type")
            break;
    }
}

function updateRoomSettings(event) {
    const data = event.payload
    console.log(data)
    updateDomSettings(data)
}

function updateServerMessage(event) {
    const data = event.payload.message

    updateDomServerMessage(data)
}

function sendEvent(eventName, payload) {
    const event = new Event(eventName, payload)

    conn.send(JSON.stringify(event))
}

function connectWs() {
    if (window.WebSocket) {
        const wsUrl = getWsUrl()
        conn = new WebSocket(wsUrl)
        conn.onopen = (e) => {
            console.log("ok")
            updateVirtualDom({
                entryDashboard: false,
                waitingRoomDashboard: true,
                gameDashboard: false,
            })
        }

        conn.onclose = (e) => {
            alert("closed connection with ws server ", e.data)
        }

        conn.onmessage = (e) => {
            const event = JSON.parse(e.data)
            routeEvent(event)
        }
    } else {
        alert("your browser doesn't support websockets")
    }
}
//////////////////////////////////////////

function getRoomName() {
    const queryString = window.location.search;
    const urlParams = new URLSearchParams(queryString);
    return urlParams.get('roomName') || '';
}

function getWsUrl() {
    const baseUrl = "ws://localhost:8090/ws"
    const queryString = window.location.search;
    const urlParams = new URLSearchParams(queryString);
    const isNewGame = urlParams.get('newGame') === 'true';

    if (isNewGame) {
        const gameSettings = {
            difficulty: urlParams.get('difficulty') || '',
            maxRounds: urlParams.get('maxRounds') || '',
            category: urlParams.get('category') || '',
        }
        console.log(gameSettings)
        return `${baseUrl}?new=true&room=${roomName}&name=${userName}&difficulty=${gameSettings.difficulty}&maxRounds=${gameSettings.maxRounds}&category=${gameSettings.category}`
    } else {
        return `${baseUrl}?room=${roomName}&name=${userName}`
    }
}

/////////// DOM /////////

function updateVirtualDom(newVirtualDom) {
    virtualDom = newVirtualDom
    console.log(virtualDom)
    handleVirtualDom()
}

function handleVirtualDom() {
    entryDashboard.classList.remove("hidden");
    waitingRoomDashboard.classList.remove("hidden");
    gameDashboard.classList.remove("hidden");
    if (virtualDom.entryDashboard) {
        waitingRoomDashboard.classList.add("hidden");
        gameDashboard.classList.add("hidden");
    } else if (virtualDom.waitingRoomDashboard) {
        entryDashboard.classList.add("hidden");
        gameDashboard.classList.add("hidden");
    } else if (virtualDom.gameDashboard) {
        entryDashboard.classList.add("hidden");
        waitingRoomDashboard.classList.add("hidden");
    }
}
function updateDomServerMessage(message) {
    displayServerMessage.innerHTML = message
}
function updateDomGameState() {
    updateVirtualDom({
        entryDashboard: false,
        waitingRoomDashboard: false,
        gameDashboard: true,
    })
    answerAElement.innerHTML = gameState.answers[0]
    answerBElement.innerHTML = gameState.answers[1]
    answerCElement.innerHTML = gameState.answers[2]
    answerDElement.innerHTML = gameState.answers[3]

    displayRoundElement.innerHTML = gameState.round
    displayQuestionElement.innerHTML = gameState.question
    displayCategoryElement.innerHTML = gameState.category
}

function updateDomScore(playerList) {
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

function updateDomReadyStatus(playerList) {
    let readyCounter = 0;

    readyUsersList.innerHTML = '';

    playerList.forEach(player => {
        const elementHTML = `
        <li id="${player.name + true}" class="text-black font-bold">
        ${player.name}
        ${player.isReady
                ? `✔`
                : `❌`}
        </li>
      `
        readyUsersList.insertAdjacentHTML("beforeend", elementHTML)

        if (player.isReady) {

            readyCounter++
        }

    });

    displayReadyCount.innerHTML = `${readyCounter}/${playerList.length}`
}


function updateDomSettings(roomSettings) {
    displayCategoryElement.innerHTML = roomSettings.category
}

///////////////////// SERVER EVENT FUNCTIONS //////////////////////

function updateGameState(event) {
    if (gameState.isGame === true) {
        updateVirtualDom({
            entryDashboard: false,
            waitingRoomDashboard: false,
            gameDashboard: true,
        })
    }
    gameState = event.payload
    console.log(gameState, "update")

    updateDomScore(gameState.score)
    updateDomGameState()
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