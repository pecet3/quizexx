import { useEffect, useState } from "react";
import { Chat, GameDashboard, WaitingRoom } from "../components/DashboardQuiz";
import { useParams } from "react-router-dom";

export type User = {
  name: string;
  points: number;
};
export type Event = {
  type: string;
  payload: any;
};

export type Settings = {
  name: string;
  gen_content: string;
  difficulty: string;
  max_rounds: number;
  language: string;
};

export type GameState = {
  round: number;
  question: string;
  answers: string[];
  actions: RoundAction[];
  score: PlayerScore[];
  playersAnswered: string[];
  roundWinners?: string[];
};

type PlayerScore = {
  user: User;
  points: number;
  roundsWon: number[];
  isAnswered: boolean;
};

type RoundQuestion = {
  question: string;
  answers: string[];
  correctAnswer: number;
};
export type RoundAction = {
  uuid: string;
  answer: string;
  round: number;
};
export type SendMessageEvent = {
  userName: string;
  message: string;
};

export type WaitingPlayer = {
  name: string;
  isReady: boolean;
};

export type WaitingState = {
  players: WaitingPlayer[];
};

export type ServerMessage = {
  message: string;
};

export type PlayersAnswered = {
  players: string[];
};

export type ChatMessage = {
  name: string;
  message: string;
  date: string;
};
const defaultGameState = {
  round: 0,
  question: "",
  answers: [],
  actions: [],
  score: [],
  playersAnswered: [],
  roundWinners: [],
};
const defaultSettings: Settings = {
  difficulty: "",
  gen_content: "",
  language: "",
  max_rounds: 0,
  name: "",
};
const defaultWaitingState: WaitingState = {
  players: [],
};
export const Quiz = () => {
  const { roomName } = useParams<{ roomName: string }>();
  const [ws, setWs] = useState<null | WebSocket>(null);
  const [isWaiting, setIsWaiting] = useState(true);

  const [gameState, setGameState] = useState<GameState>(defaultGameState);
  const [settings, setSettings] = useState<Settings>(defaultSettings);
  const [waitingState, setWaitingState] =
    useState<WaitingState>(defaultWaitingState);

  const [users, setUsers] = useState<User[]>([]);

  const handleReady = () => {
    setIsWaiting(false);
    ws?.send(
      JSON.stringify({
        type: "ready_player",
        payload: "",
      })
    );
  };

  const handleAnswer = (answer: number) => {
    console.log("Selected answer:", answer);
  };
  function routeEvent(event: Event) {
    console.log("event type", event.type);

    console.log(event.payload);
    if (event.type === undefined) {
      alert("no type field in the event");
    }
    switch (event.type) {
      case "update_gamestate":
        setGameState(event.payload);
        break;
      case "update_players":
        break;
      case "server_message":
        break;
      case "waiting_state":
        setWaitingState(event.payload);
        break;
      case "finish_game":
        break;
      case "room_settings":
        setSettings(event.payload);
        break;
      case "players_answered":
        break;
      case "chat_message":
        break;
      default:
        console.log(event.type);
        break;
    }
  }
  // debug
  useEffect(() => {
    console.log("gamestate", gameState);
    console.log("settings", settings);
  }, [gameState, settings]);
  useEffect(() => {
    const ws = new WebSocket(`ws://localhost:9090/api/quiz/${roomName}`);
    setWs(ws);
    ws.onopen = () => {
      console.log("connected with websockets!");
    };

    ws.onmessage = (event) => {
      try {
        const eventJSON = JSON.parse(event.data);
        routeEvent(eventJSON);
      } catch (error) {
        console.error(error);
      }
    };

    ws.onerror = (error) => {
      console.error(error);
    };

    ws.onclose = () => {
      console.log("closed connection with websockets");
    };

    return () => {
      ws.close();
    };
  }, []);

  return (
    <div className="p-2 bg-opacity-70 text-center m-auto">
      {isWaiting ? (
        <WaitingRoom waitingState={waitingState} onReady={handleReady} />
      ) : (
        <>
          <GameDashboard
            settings={settings!}
            gameState={gameState!}
            users={users}
            onAnswer={handleAnswer}
          />
          <Chat />
        </>
      )}
    </div>
  );
};
