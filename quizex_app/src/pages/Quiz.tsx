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

export const Quiz = () => {
  const { roomName } = useParams<{ roomName: string }>();
  const [ws, setWs] = useState<null | WebSocket>(null);

  const [isWaiting, setIsWaiting] = useState(true);
  const [users, setUsers] = useState<User[]>([]);
  const [currentUser, setCurrentUser] = useState<User | null>(null);

  const [category] = useState("Sample Category");
  const [round] = useState(1);
  const [question] = useState("Sample Question?");
  const [answers] = useState(["Answer A", "Answer B", "Answer C", "Answer D"]);

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
    console.log(event.type);

    console.log(event.payload);
    if (event.type === undefined) {
      alert("no type field in the event");
    }
    switch (event.type) {
      case "update_gamestate":
        break;
      case "update_players":
        break;
      case "server_message":
        break;
      case "waiting_state":
        break;
      case "finish_game":
        break;
      case "room_settings":
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
  useEffect(() => {
    // Tworzenie połączenia WebSocket
    const ws = new WebSocket(`ws://localhost:9090/api/quiz/${roomName}`);
    setWs(ws);
    ws.onopen = () => {
      console.log("connected with websockets!");
    };

    ws.onmessage = (event) => {
      ws.send(event as any);
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
        <WaitingRoom readyUsers={users} onReady={handleReady} />
      ) : (
        <>
          <GameDashboard
            category={category}
            round={round}
            question={question}
            answers={answers}
            users={users}
            onAnswer={handleAnswer}
          />
          <Chat />
        </>
      )}
    </div>
  );
};
