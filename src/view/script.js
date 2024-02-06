const connectButton = document.getElementById("connectButton")

connectButton.addEventListener("click", () => {
    const roomName = document.getElementById('joinRoomInput')

    const room = roomName.value

    const link = '/room/?roomName=' + encodeURIComponent(room)

    window.location.href = link;
})