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

export const Quiz = () => {
  const { roomName } = useParams<{ roomName: string }>();
  const [isWaiting, setIsWaiting] = useState(true);
  const [users, setUsers] = useState<User[]>([]);
  const [currentUser, setCurrentUser] = useState<User | null>(null);

  const [category] = useState("Sample Category");
  const [round] = useState(1);
  const [question] = useState("Sample Question?");
  const [answers] = useState(["Answer A", "Answer B", "Answer C", "Answer D"]);

  const handleReady = () => {
    setIsWaiting(false);
  };

  const handleAnswer = (answer: number) => {
    console.log("Selected answer:", answer);
  };
  function routeEvent(event: Event) {
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
