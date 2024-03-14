interface IEvent {
    type: "ready_status" | "room_settings" | "update_gamestate" | "server_message" | "finish_game" | "update_players";
    payload: TReadyStatus | TRoomSettings | TGameState | TServerMessage | TFinishGame;
};

export type TReadyStatus = {
    clients: ReadyClient[];
};
export type TReadyClient = {
    name: string;
    isReady: boolean;
};
export type TRoomSettings = {
    roomName: string;
    category: string;
    difficulty: string;
    maxRounds: string;
};

export type TServerMessage = {
    message: string;
};

export type TGameState = {
    round: number;
    question: string;
    answers: string[];
    actions: RoundAction[];
    score: PlayerScore[];
    playersFinished: string[];
};
export type TRoundAction = {
    name: string;
    answer: number;
    round: number;
};
export type TPlayerScore = {
    name: string;
    points: number;
    roundsWon: number[];
};


export type TFinishGame = bool


// in the future we will use it
// type TSendMessageEvent = {
//     userName: string;
//     message: string;
// };
/////// D