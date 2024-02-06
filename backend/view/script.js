let conn;
let userName;

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
