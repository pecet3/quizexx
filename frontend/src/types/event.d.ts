interface IEvent {
    type: "ready_status" | "room_settings" | "update_gamestate" | "server_message" | "finish_game";
    payload: TReadyStatus | TRoomSettings | TGameState | TServerMessage | TFinishGame;
};

type TReadyStatus = {
    clients: ReadyClient[];
};
type TReadyClient = {
    name: string;
    isReady: boolean;
};
type TRoomSettings = {
    name: string;
    category: string;
    difficulty: string;
    maxRounds: string;
};

type TServerMessage = {
    message: string;
};

type TGameState = {
    round: number;
    question: string;
    answers: string[];
    actions: RoundAction[];
    score: PlayerScore[];
    playersFinished: string[];
};
type TRoundAction = {
    name: string;
    answer: number;
    round: number;
};
type TPlayerScore = {
    name: string;
    points: number;
    roundsWon: number[];
};


type TFinishGame = bool


// in the future we will use it
// type TSendMessageEvent = {
//     userName: string;
//     message: string;
// };
/////// D