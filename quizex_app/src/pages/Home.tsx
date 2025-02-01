import { RoomsList } from "../components/RoomsList";
import { PaperWrapper, PatternWrapper } from "../components/Wrappers";
import { useProtectedContext } from "../context/protectedContext";
import UserProfileCard from "../components/UserProfileCarrd";
import { TopUsersList } from "../components/TopUsersList";
import { RoomJoiner } from "../components/RoomJoiner";

export const Home = () => {
  const { funFact, user, topUsers } = useProtectedContext();
  console.log(funFact);
  return (
    <main className="grid grid-cols-1 lg:grid-cols-2 p-1 pt-10 px-2 xl:px-40 justify-items-center gap-16 lg:gap-4  w-full m-auto justify-center">
      <PaperWrapper>
        <RoomsList />
      </PaperWrapper>
      <div className="flex flex-col gap-8 items-center self-center">
        <RoomJoiner />
        <UserProfileCard user={user!} />
      </div>
      <div className="self-center ">
        <p className="italic text-xl sm:text-2xl font-handwritten">
          {funFact?.content}
        </p>{" "}
        <span className="text-sm sm:text-base flex justify-end gap-2">
          Fun fact about:
          <p className="text-right font-bold"> {funFact?.topic}</p>
        </span>
      </div>
      <PatternWrapper>
        <TopUsersList users={topUsers!} />
      </PatternWrapper>
    </main>
  );
};
