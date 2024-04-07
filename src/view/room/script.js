

const roomName = getRoomName();
let userName = "";

let ready = true;
let isAnswerSent = false;
let isPlayerReady = false;

let settings = {
    category: "",
    difficulty: "",
    maxRounds: "",
    name: "",
    lang: "",
}

let gameState = {
    isGame: false,
    round: 1,
    question: "Test",
    answers: ["Lorem ipsum", "Lorem ipsum", "Lorem ipsum", "Lorem ipsum"],
    actions: [{ name: "", answer: null, round: 0 }],
    score: [{ name: "", points: 0, roundsWon: [], isAnswered: false }],
};

// initial state "virtual dom"
let virtualDom = {
    entryDashboard: true,
    minorLogo: true,
    waitingRoomDashboard: false,
    gameDashboard: false,
}

if (gameState.isGame) {
    virtualDom = {
        entryDashboard: true,
        waitingRoomDashboard: false,
        gameDashboard: false,
    }
}
// handleVirtualDom()

//////////////// Listeners ///////////////////////

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
    if (answerValue !== null && !isAnswerSent) {
        sendAnswer(answer);
        isAnswerSent = true
        formData = null
        return
    } else {
        alert("no answer :(");
    }
});

chatForm.addEventListener("submit", (e) => {
    e.preventDefault()
    let formData = new FormData(chatForm)
    const chatInput = formData.get('chatInput')
    console.log(chatInput)
    chatForm.elements['chatInput'].value = "";
})

////////// W E B socket connection /////

function connectWs() {
    if (window.WebSocket) {
        const wsUrl = getWsUrl()
        conn = new WebSocket(wsUrl)
        conn.onopen = (e) => {
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
///////////// H E L P e r s /////////////

function getRoomName() {
    const queryString = window.location.search;
    const urlParams = new URLSearchParams(queryString);
    return urlParams.get('roomName') || '';
}

function getWsUrl() {
    const baseUrl = "ws://127.0.0.1:8090/ws"
    const queryString = window.location.search;
    const urlParams = new URLSearchParams(queryString);
    const isNewGame = urlParams.get('newGame') === 'true';

    if (isNewGame) {
        const gameSettings = {
            difficulty: urlParams.get('difficulty') || '',
            maxRounds: urlParams.get('maxRounds') || '',
            category: urlParams.get('category') || '',
            lang: urlParams.get("lang") || '',
        }
        return `${baseUrl}?new=true&room=${roomName}&name=${userName}&difficulty=${gameSettings.difficulty}&maxRounds=${gameSettings.maxRounds}&category=${gameSettings.category}&lang=${gameSettings.lang}`
    } else {
        return `${baseUrl}?room=${roomName}&name=${userName}`
    }
}


