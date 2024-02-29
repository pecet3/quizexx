const gameForm = document.getElementById('gameForm');
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
// const displayPlayersInGame = document.getElementById('displayPlayersInGame')
const displayReadyCount = document.getElementById('displayReadyCount')
const displayServerMessageWaiting = document.getElementById('displayServerMessageWaiting')
const displayServerMessageDashboard = document.getElementById('displayServerMessageDashboard')


function updateVirtualDom(newVirtualDom) {
    virtualDom = newVirtualDom
    console.log(virtualDom)
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
    answerAElement.innerHTML = gameState.answers[0]
    answerBElement.innerHTML = gameState.answers[1]
    answerCElement.innerHTML = gameState.answers[2]
    answerDElement.innerHTML = gameState.answers[3]

    displayRoundElement.innerHTML = gameState.round
    displayQuestionElement.innerHTML = gameState.question
    displayCategoryElement.innerHTML = gameState.category
    // displayPlayersInGame.innerHTML = gameState.actions.length
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

function updateDomSettings(data) {
    const displayCategoryElement = document.getElementById('displayCategory');
    displayCategoryElement.innerHTML = data.category

}

