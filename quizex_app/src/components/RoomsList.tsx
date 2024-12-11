import { Link } from "react-router-dom";

// in map todo: fixed w-64 in name

export const RoomsList = () => {
  return (
    <div className="text-base flex flex-col py-8 px-6 max-w-2xl w-full">
      <Link to={"/test"} className="btn text-xs bg-red-300">
        Create a Room
      </Link>
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
          <tr className="border-b border-black w-full">
            <td className="px-4 py-2 w-64"></td>
            <td className="px-4 py-2">2/4</td>
            <td className="px-4 py-2">1</td>
            <td className="px-4 py-2 text-center">
              <button className="btn bg-teal-300 text-xs">Join</button>
            </td>
          </tr>
          <tr className="border-b border-black w-full">
            <td className="px-4 py-2">Room1</td>
            <td className="px-4 py-2">2/4</td>
            <td className="px-4 py-2">1</td>
            <td className="px-4 py-2 text-center">
              <button className="btn bg-teal-300 text-xs">Join</button>
            </td>
          </tr>
          <tr className="border-b border-black w-full">
            <td className="px-4 py-2">Room1</td>
            <td className="px-4 py-2">2/4</td>
            <td className="px-4 py-2">1</td>
            <td className="px-4 py-2 text-center">
              <button className="btn bg-teal-300 text-xs">Join</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  );
};
