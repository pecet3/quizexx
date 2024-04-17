const gameForm = document.getElementById('gameForm');
const chatForm = document.getElementById('chatForm');

const connectButton = document.getElementById("connectButton")
const readyButton = document.getElementById("readyButton")
const readyUsersList = document.getElementById('readyUsersList')

const answerAElement = document.getElementById('answerA');
const answerBElement = document.getElementById('answerB');
const answerCElement = document.getElementById('answerC');
const answerDElement = document.getElementById('answerD');

const displayRoundElement = document.getElementById('displayRound');
const displayQuestionElement = document.getElementById('displayQuestion');
const displayPlayers = document.getElementById('displayPlayersInGame')
const displayReadyCount = document.getElementById('displayReadyCount')
const displayServerMessageWaiting = document.getElementById('displayServerMessageWaiting')
const displayServerMessageDashboard = document.getElementById('displayServerMessageDashboard')
const displayCountAnswered = document.getElementById('displayAnswered')

function updateVirtualDom(newVirtualDom) {
    virtualDom = newVirtualDom
    if (gameState.isGame) {
        virtualDom = {
            entryDashboard: false,
            waitingRoomDashboard: false,
            gameDashboard: true,
        }
    }
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
    if (virtualDom.waitingRoomDashboard) {
        displayServerMessageWaiting.innerHTML = message
    } else if (virtualDom.gameDashboard) {
        displayServerMessageDashboard.innerHTML = message
    }
}
function updateDomGameState() {
    updateVirtualDom({
        entryDashboard: false,
        waitingRoomDashboard: false,
        gameDashboard: true,
    })
    answerAElement.innerText = gameState.answers[0]
    answerBElement.innerText = gameState.answers[1]
    answerCElement.innerText = gameState.answers[2]
    answerDElement.innerText = gameState.answers[3]

    displayRoundElement.innerText = gameState.round
    displayQuestionElement.innerText = gameState.question
    // displayPlayersInGame.innerText = gameState.actions.length

    // displayCountAnswered.innerText = `${gameState.playersFinished.length}/${gameState.score.length}`
}

function updateDomScore(playerList, answeredList) {
    const tableBody = document.getElementById('scoreTableBody');

    tableBody.innerText = '';

    playerList.forEach(player => {
        const row = document.createElement('tr');
        const nameCell = document.createElement('td');
        const pointsCell = document.createElement('td');
        pointsCell.textContent = player.points;

        nameCell.textContent = player.name;

        if (answeredList !== undefined) {
            answeredList.forEach(a => {
                if (a === player.name) {
                    nameCell.textContent = player.name + "✔";
                }
            }
            )
        }

        row.appendChild(nameCell);
        row.appendChild(pointsCell);
        tableBody.appendChild(row);
    });
}

function updateDomReadyStatus(playerList) {
    let readyCounter = 0;

    readyUsersList.innerText = '';

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

    displayReadyCount.innerText = `${readyCounter}/${playerList.length}`
}

function updateDomSettings(data) {
    const displayCategoryElement = document.getElementById('displayCategory');
    displayCategoryElement.innerText = data.category

}

function updateDomChatMessages(data) {
    const chatMessages = document.getElementById("chatMessages")

    const elementHTML = `
      ${userName === data.name
            ? `<li class="flex flex-row-reverse p-0.5 w-full ">`
            : `<li class="flex p-0.5 w-full">`}


      ${userName === data.name
            ? `<div class="p-1 flex flex-row-reverse rounded-md bg-white border border-black break-words ">`
            : `<div class="p-1 flex bg-white border rounded-md border-black break-words">`}

        <div class="break-words flex flex-col justify-end items-center">

        ${data.name === userName
            ? `<a class="text-gray-700 ">(You)</a>`
            : `<a class="font-bold text-teal-500 underline"> ${data.name}: </a>`}

        ${typeof data.time !== 'undefined'
            ? `<a class="font-mono text-[12px]">${data.time}</a>`
            : ""}
        </div>
        
        
        <a class="px-0.5"> ${data.message} </a>

        </div>
      </li>
    `

    chatMessages.insertAdjacentHTML("beforeend", elementHTML)
    return
}

