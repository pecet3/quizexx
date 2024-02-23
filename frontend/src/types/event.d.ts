interface IEvent {
    type: "ready_status" | "room_settings" | "update_gamestate" | "server_message" | "finish_game";
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
    name: string;
    category: string;
    difficulty: string;
    maxRounds: string;
};

export export type TServerMessage = {
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