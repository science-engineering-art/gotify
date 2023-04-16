import { useDispatch } from "react-redux";
import { selectedSong } from "../features/songsSlice";

export type SongItemPropsType = {
  song: {
    Id: string;
    artist: string;
    title: number;
    year: string;
  };
};

export const SongItem = ({ song }: SongItemPropsType) =>{
  const dispatch = useDispatch();

  return (
      <li 
      className="w-full p-5 bg-cyan-500 my-2" 
      onClick={_ => {dispatch(selectedSong(song.Id))}}
      >
      <h1 className="text-2xl"> {song?.title} </h1>
      <div> By {song?.artist}, {song?.year}</div>
    </li>
  );
} 
