const settingsForm = document.getElementById('settingsForm');
const categoryInput = document.getElementById('categoryInput');
const nameInput = document.getElementById('nameInput');
const maxRoundsInput = document.getElementById('maxRounds');
const difficultyInput = document.getElementById('difficulty');

settingsForm.addEventListener('submit', function (event) {
    event.preventDefault();

    const categoryValue = categoryInput.value;
    const maxRoundsValue = maxRoundsInput.value;
    const difficultyValue = difficultyInput.value;
    const nameValue = nameInput.value;

    const queryParams = new URLSearchParams();
    queryParams.set('roomName', nameValue);
    queryParams.set('newGame', true);

    queryParams.set('difficulty', difficultyValue);
    queryParams.set('maxRounds', maxRoundsValue);
    queryParams.set('category', categoryValue);

    const redirectURL = `/room?${queryParams.toString()}`;

    window.location.href = redirectURL;
});

function getRoomNames() {
    try {
        const res = fetch("")
    } catch {
        console.error("error fetch room names")
    }
}