import { useAppStateContext } from "../../custom-hooks/useAppContext";
import { useWebSocket } from "../../custom-hooks/useWebSocket";
import { GameDashboard } from "./GameDashboard";
import { WaitingRoom } from "./WaitingRoom";

export function Room() {
  const { appState } = useAppStateContext();

  const socket = useWebSocket()

  console.log(socket, "<-socket")
  return (
    <div className="text-center m-auto flex flex-col items-center">
      <WaitingRoom />

      <GameDashboard />
    </div>
  );
}
