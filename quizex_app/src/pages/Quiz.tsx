import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { WaitingRoom } from "../components/quiz/WaitingRoom";
import { Dashboard } from "../components/quiz/Dashboard";
import { Chat } from "../components/quiz/Chat";
import { Error } from "../components/Error";
import { useProtectedContext } from "../context/protectedContext";

export type Player = {
  name: string;
  points: number;
  is_answered: boolean;
  user: User;
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
  sec_for_answer: number;
};

export type GameState = {
  round: number;
  question: string;
  answers: string[];
  actions: RoundAction[];
  score: PlayerScore[];
  players_answered: string[];
  roundWinners?: string[];
};

type PlayerScore = {
  user: User;
  points: number;
  roundsWon: number[];
  isAnswered: boolean;
};

export type RoundAction = {
  uuid: string;
  answer: number;
  round: number;
};
export type SendMessageEvent = {
  userName: string;
  message: string;
};

export type WaitingPlayer = {
  name: string;
  is_ready: boolean;
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
  date: Date;
};
const defaultGameState: GameState = {
  round: 0,
  question: "",
  answers: [""],
  actions: [],
  score: [
    {
      isAnswered: false,
      points: 0,
      roundsWon: [0],
      user: {
        name: "",
        createdAt: new Date(),
        email: "",
        image_url: "",
        is_draft: false,
        uuid: "",
      },
    },
  ],
  players_answered: [],
  roundWinners: [],
};
const defaultSettings: Settings = {
  difficulty: "",
  gen_content: "",
  language: "",
  max_rounds: 0,
  sec_for_answer: 0,
  name: "",
};
const defaultWaitingState: WaitingState = {
  players: [],
};
export const Quiz = () => {
  const { user } = useProtectedContext();
  const { roomName } = useParams<{ roomName: string }>();
  const [ws, setWs] = useState<null | WebSocket>(null);
  const [isWaiting, setIsWaiting] = useState(true);
  const [messages, setMessages] = useState<ChatMessage[]>([]);
  const [timer, setTimer] = useState(0);

  const [err, setErr] = useState("");
  const [serverMessage, setServerMessage] = useState("");
  const [gameState, setGameState] = useState<GameState>(defaultGameState);
  const [settings, setSettings] = useState<Settings>(defaultSettings);

  const [waitingState, setWaitingState] =
    useState<WaitingState>(defaultWaitingState);

  const handleReady = () => {
    ws?.send(
      JSON.stringify({
        type: "ready_player",
        payload: "",
      })
    );
  };
  // checking if is game
  useEffect(() => {
    if (gameState.round !== 0) setIsWaiting(false);
  }, [gameState]);

  const handleAnswer = (answer: number) => {
    if (!user) return;
    console.log("Selected answer:", answer);
    const payload: RoundAction = {
      answer,

      round: gameState.round,
      uuid: user.uuid,
    };
    ws?.send(
      JSON.stringify({
        type: "send_answer",
        payload,
      })
    );
  };
  const handleMessage = (message: string) => {
    if (!user) return;
    const payload: ChatMessage = {
      message,
      name: user.name,
      date: new Date(),
    };
    ws?.send(
      JSON.stringify({
        type: "chat_message",
        payload,
      })
    );
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
        console.log();
        break;
      case "server_message":
        setServerMessage(event.payload.message);
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
        setGameState((prev) => ({
          ...prev,
          players_answered: event.payload,
        }));

        break;
      case "chat_message":
        setMessages((prev) => [...prev, event.payload]);
        break;
      case "set_timer":
        setTimer(event.payload);
        break;
      case "update_players":
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
    const ws = new WebSocket(`ws://localhost:9090/api/quiz/rooms/${roomName}`);
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

    ws.onerror = () => {
      setErr("Something went wrong...");
    };

    ws.onclose = () => {};

    return () => {
      setErr("Something went wrong...");

      ws.close();
    };
  }, []);

  return (
    <div className="p-2 bg-opacity-70 text-center m-auto">
      {err != "" ? (
        <Error err={err} />
      ) : (
        <>
          {isWaiting ? (
            <WaitingRoom
              serverMessage={serverMessage}
              waitingState={waitingState}
              onReady={handleReady}
            />
          ) : (
            <>
              <Dashboard
                settings={settings!}
                gameState={gameState!}
                serverMessage={serverMessage}
                onAnswer={handleAnswer}
                timer={timer}
              />
              <Chat onMessage={handleMessage} messages={messages} />
            </>
          )}
        </>
      )}
    </div>
  );
};
