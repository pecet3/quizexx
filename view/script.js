const entryForm = document.getElementById("entryForm")
const userName = document.getElementById("name")
const room = document.getElementById("room")
const generateBtn = document.getElementById("generateBtn")
const messageForm = document.getElementById("messageForm")
 
let namesArr = [""]
replaceInputRoom("room_1")

// LISTENERS

generateBtn.addEventListener("click",()=>{
    room.value = generateRoomName(8)

})

entryForm.addEventListener("submit",(e)=>{
    e.preventDefault();

    connectWs();
})

messageForm.addEventListener("submit",(e)=>{
    e.preventDefault();
    handleMessage(e)
})

messageForm.addEventListener("keydown", (e) => {
    if (e.key === 13) {
        e.preventDefault();
        handleMessage(e)    }
});

///// WS

function connectWs(){
   
    if (userName.value === "" || room.value === ""){
        return
    }

    if (window.WebSocket){
        conn = new WebSocket(`wss://czatex.pecet.it/ws?room=${room.value}&name=${userName.value}`)
        conn.onopen = (e)=>{
            showDashboardHiddeEntry()
            writeRoomTitle()
            addQuery("room",room.value)

            const greetingsMsg = {
                name: "klient",
                message: "Wpisz /users, aby zobaczyć nazwy użytkowników w pokoju."
            }
            writeMessage(greetingsMsg)
        }

        conn.onclose=(e)=>{
            alert("closed connection with ws server ")
        }

        conn.onmessage = (e)=>{
            const data = JSON.parse(e.data)
            writeMessage(data)
            writeClients(e)
            

        }
    }else{
        alert("your browser doesn't support websockets")
    }
}


//// DOM 
function writeMessage(data){
    const messagesList = document.getElementById("messagesList")
    console.log(userName.value, data.name)
    const elementHTML = `
      ${userName.value === data.name 
        ? `<li class="flex flex-row-reverse text-slate-200">`
        : `<li class="flex text-slate-200 p-0.5">` }


      ${userName.value === data.name 
        ? `<div class="p-1 flex flex-row-reverse bg-slate-700 rounded-md break-words max-w-64 sm:max-w-[38rem]">`
        : `<div class="p-1 flex bg-slate-700 rounded-md max-w-64 sm:max-w-[38rem] break-words">` }

        <div class="break-words flex flex-col justify-center items-center">

        ${data.name === "serwer" || data.name ==="klient" 
        ? `<a class="font-bold text-pink-500">[${data.name}] </a>` 
        : `<a class="font-bold text-pink-500">[${data.name}] </a>`}

        ${typeof data.date !== 'undefined' 
        ? `<a class="font-mono text-[12px]">${data.date}</a>` 
        : ""}
        </div>
        
        
        <a class="">${data.message}</a>

        </div>
      </li>
    `
  
    messagesList.insertAdjacentHTML("beforeend",elementHTML)
    return
}


function trackLastListElement(){

}

function showDashboardHiddeEntry(){
    const chatDashboard = document.getElementById("chatDashboard")
    entryForm.classList.add("hidden")
            
    chatDashboard.classList.remove("hidden")
    chatDashboard.classList.add("flex")
    return
}

function writeRoomTitle(){
    const roomDisplay = document.getElementById("roomDisplay")
    roomDisplay.textContent = room.value
    return
}

function writeClients(e){
    const clientsDisplay = document.getElementById("clientsDisplay")
    const data = JSON.parse(e.data)
    
    namesArr = data.clients

    clientsDisplay.textContent = data.clients.length
    return
}

//// HELPERS
function handleMessage(e){
    const messageElement = e.target.elements.message

    if (userName.value === "" || messageElement.value === ""){
        return
    }

    const trimmedMsg = messageElement.value.trim()

    if (trimmedMsg[0] === "/"){
        handleUserCmd(trimmedMsg)
        resetMessageInput(messageElement)
        return
    }

    const date = getCurrentDateTimeString()
    let data = {
        "name": userName.value,
        "message": messageElement.value,
        "date": date,
        "clients": [""]
    }
    

    conn.send(JSON.stringify(data));
    resetMessageInput(messageElement)
    return
}

function resetMessageInput(messageElement){
    messageElement.value = ""
    messageElement.focus()
    return
}

function handleUserCmd(cmd){
    const data = {
        name: "klient",
        message: "na serwerze są: " + namesArr.toString(),
    }

    if (cmd ==="/users"){
        writeMessage(data)
        return
    }
}

function replaceInputRoom(value){
    const url = new URL(window.location.href)
    const params = new URLSearchParams(url.search)
    const queryRoom = params.get("room")

    if (queryRoom ==="" || queryRoom === null){
        room.value = value
        return
    }
    room.value = queryRoom
    return
}


function addQuery(param,value){
    const url = new URL(window.location.href)
    url.searchParams.set(param,value)
    history.replaceState(null,null, url.href)
    return
}

function generateRoomName(length) {
    const characters = '0123456789ABCDE';
    let randomId = '0x';
  
    for (let i = 0; i < length; i++) {
      const randomIndex = Math.floor(Math.random() * characters.length);
      randomId += characters.charAt(randomIndex);
    }
  
    return randomId;
}


function getCurrentDateTimeString() {
    const currentDate = new Date();
  
    const year = currentDate.getFullYear();
    const month = (currentDate.getMonth() + 1).toString().padStart(2, '0');
    const day = currentDate.getDate().toString().padStart(2, '0');
    const hours = currentDate.getHours().toString().padStart(2, '0');
    const minutes = currentDate.getMinutes().toString().padStart(2, '0');
  
  
    const dateTimeString = `${year}-${month}-${day} ${hours}:${minutes}`;
  
    return dateTimeString;
  }
  