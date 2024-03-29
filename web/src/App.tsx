import { SplitScreen } from "./layouts/SplitScreen";
import { NavBar } from "./components/NavBar";
import { SongsList } from "./components/SongsList";
import { SongFilter } from "./components/SongFilter";

export default function App() {

  return (
    <div className="w-full h-full">
      <NavBar />
      <SplitScreen className="min-h-[600px]" leftWidth={4} rightWidth={8}>
        <div className="w-full h-full p-10 bg-orange-500">
          <SongFilter />
        </div>
        <SongsList />
      </SplitScreen>
    </div>
  );
}
