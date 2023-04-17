import { SplitScreen } from "./layouts/SplitScreen";
import { NavBar } from "./components/NavBar";
import { SongsList } from "./components/SongsList";
import { AudioPlayer } from "./components/AudioPlayer";
import { useSelector } from "react-redux";

export default function App() {
  const songId = useSelector((state: { songs: { Id: string } }) => state.songs.Id);

  return (
    <div className="w-full h-full">
      <NavBar />
      <SplitScreen className="min-h-[600px]" leftWidth={4} rightWidth={8}>
        <div className="w-full h-full p-10 bg-orange-500">
          {/* {songId && <AudioPlayer src={`http://localhost:5000/api/song/${songId}`} />} */}
          <h1>Player</h1>
        </div>
        <SongsList />
      </SplitScreen>
    </div>
  );
}
