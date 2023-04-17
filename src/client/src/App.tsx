import { SplitScreen } from "./layouts/SplitScreen";
import { NavBar } from "./components/NavBar";
import { SongsList } from "./components/SongsList";

export default function App() {
  return (
    <div className="w-full h-full">
      <NavBar />
      <SplitScreen className="min-h-[600px]" leftWidth={4} rightWidth={8}>
        <div className="w-full h-full p-10 bg-orange-500">
          <h1 className="text-3xl text-center">Player</h1>
        </div>
        <SongsList />
      </SplitScreen>
    </div>
  );
}
