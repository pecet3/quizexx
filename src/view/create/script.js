const settingsForm = document.getElementById('settingsForm');
const categoryInput = document.getElementById('categoryInput');
const nameInput = document.getElementById('nameInput');
const maxRoundsInput = document.getElementById('maxRounds');
const difficultyInput = document.getElementById('difficulty');
const langInput = document.getElementById("lang");


const categories = [
    "General Knowledge",
    "Geography",
    "Music",
    "Movies",
    "Sports",
    "Literature",
    "Art and Artists",
    "Food and Cuisine",
    "Famous Landmarks",
    "Technology",
    "Science Fiction",
    "Mythology",
    "Animals",
    "World History",
    "Current Events",
    "Fashion",
    "Languages",
    "Famous Quotes",
    "Celebrities",
    "Math Puzzles",
    "Astronomy",
    "Pop Culture",
    "Environmental Science",
    "Business and Economics",
    "Psychology",
    "Inventions and Discoveries",
    "Architecture",
    "Cartoons and Animation",
    "Trivia",
    "Health and Wellness",
    "Cryptography",
    "Philosophy",
    "Superheroes",
    "Board Games",
    "Photography",
    "Automobiles",
    "World Religions",
    "Dance Styles",
    "Military History",
    "Gardening and Botany"
];

settingsForm.addEventListener('submit', function (event) {
    event.preventDefault();

    const categoryValue = categoryInput.value;
    const maxRoundsValue = maxRoundsInput.value;
    const difficultyValue = difficultyInput.value;
    const langValue = langInput.value;
    const nameValue = nameInput.value;

    const queryParams = new URLSearchParams();
    queryParams.set('roomName', nameValue);
    queryParams.set('newGame', true);

    queryParams.set('difficulty', difficultyValue);
    queryParams.set('maxRounds', maxRoundsValue);
    queryParams.set('category', categoryValue);
    queryParams.set('lang', langValue)

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