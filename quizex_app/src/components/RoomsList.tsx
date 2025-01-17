import { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import { HiOutlineRefresh } from "react-icons/hi";

export const RoomsList = () => {
  const [rooms, setRooms] = useState<Room[]>([]);
  const [loading, setLoading] = useState(false);

  const fetchRooms = async () => {
    setLoading(true);
    try {
      const response = await fetch("/api/quiz/rooms");
      if (!response.status) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data: Rooms = await response.json();
      if (!data.rooms) return;
      if (data.rooms.length > 0) {
        setRooms(data.rooms);
      }
    } catch (error) {
      console.error("Error fetching rooms:", error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchRooms();
  }, []);

  return (
    <div className="text-base flex flex-col py-8 px-6 max-w-2xl w-full">
      <div className="flex justify-between mb-6">
        <button
          onClick={fetchRooms}
          className="btn text-xs bg-blue-300 my-4 self-start"
          disabled={loading}
        >
          <HiOutlineRefresh size={18} />
        </button>
        <Link to="/create-room" className="btn text-xs bg-red-300">
          Create a Room
        </Link>
      </div>

      <table className="table-auto border-collapse w-auto">
        <thead>
          <tr>
            <th className="px-4 py-2">Name</th>
            <th className="px-4 py-2">Players</th>
            <th className="px-4 py-2">Round</th>
            <th className="px-4 py-2">&nbsp;</th>
          </tr>
        </thead>
        <tbody>
          {rooms.length > 0 ? (
            rooms.map((room: Room) => (
              <tr key={room.uuid} className="border-b border-black w-full">
                <td className="px-4 py-2 w-64">{room.name}</td>
                <td className="px-4 py-2">
                  {room.players}/{room.max_players}
                </td>
                <td className="px-4 py-2">
                  {room.round}/{room.max_rounds}
                </td>
                <td className="px-4 py-2 text-center">
                  <Link
                    to={`/quiz/${room.uuid}`}
                    className="btn bg-teal-300 text-xs"
                  >
                    Join
                  </Link>
                </td>
              </tr>
            ))
          ) : (
            <tr>
              <td colSpan={4} className="px-4 py-2 text-center">
                {loading ? "Loading rooms..." : "No rooms available"}
              </td>
            </tr>
          )}
        </tbody>
      </table>
    </div>
  );
};
