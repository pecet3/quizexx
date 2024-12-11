export const RoomsList = () => {
  return (
    <div className="text-base flex py-8 px-6">
      <table className="table-auto border-collapse w-full">
        <thead>
          <tr>
            <th className="px-4 py-2">Name</th>
            <th className="px-4 py-2">Players</th>
            <th className="px-4 py-2">Round</th>
            <th className="px-4 py-2">&nbsp;</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td className="px-4 py-2">Room1</td>
            <td className="px-4 py-2">2/4</td>
            <td className="px-4 py-2">1</td>
            <td className="px-4 py-2 text-center">
              <button className="btn bg-teal-300">Join</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  );
};
