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