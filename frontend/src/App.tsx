import "./App.css";
import { CreateRoom } from "./components/CreateRoom/CreateRoom";
import { MainView } from "./components/MainView/MainView";
import { Room } from "./components/Room/Room";

function App() {
  return (
    <>
      <MainView />
      <CreateRoom />
      <Room />
    </>
  );
}

export default App;
