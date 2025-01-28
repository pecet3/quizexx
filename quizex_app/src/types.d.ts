interface User {
  created_at: time;
  email: string;
  exp: number;
  image_url: string;
  is_draft: boolean;
  level: number;
  name: string;
  progress: number;
  uuid: string;
}

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


type FunFact = {
  topic: string;
  content: string;
}
