export interface User {
  uuid: string; 
  name: string; 
  email: string; 
  image_url: string; 
  createdAt: Date; 
}

export type QuizSettings = {
    name: string;
    gen_content: string;
    difficulty: string;
    max_rounds: string;
    language: string;
};

type Room = {
  uuid: string;
  name: string;
  players: number;
  max_players: number;
  round: number;
  max_rounds: number;
};

type Rooms = {
  rooms: Room[];
};

type GameState= {
  round: number;
  question: string;
  answers: string[];
  actions: RoundAction[];
  score: PlayerScore[];
  playersAnswered: string[];
  roundWinners?: string[];
}

type  RoundAction = {
  name: string;
  answer: number;
  round: number;
}

type  PlayerScore =  {
  user: User;
  points: number;
  roundsWon: number[];
  isAnswered: boolean;
}

type RoundQuestion = {
  question: string;
  answers: string[];
  correctAnswer: number;
}