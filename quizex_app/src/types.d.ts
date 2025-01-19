interface User {
  uuid: string; 
  name: string; 
  email: string; 
  image_url: string;
  is_draft: boolean;
  createdAt: Date; 
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

