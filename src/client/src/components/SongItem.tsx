import { useDispatch, useSelector } from "react-redux";
import { selectedSongId, selectedSongURL } from "../features/songsSlice";
import { AudioPlayer } from "./AudioPlayer";
import { MouseEvent } from "react";

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
  const songId = useSelector((state: { songs: { Id: string } }) => state.songs.Id);
  const songUrl = useSelector((state: { songs: { url: string } }) => state.songs.url);

  const handleSelectedSong = async (_: MouseEvent<HTMLLIElement>) => {
    dispatch(selectedSongId(song.Id))
    dispatch(selectedSongURL(`http://localhost:5000/api/song/${song.Id}`))
  }

  return (
    <li 
      className="w-full p-5 bg-green-500 my-2 rounded-3xl hover:bg-opacity-60" 
      onClick={handleSelectedSong}
    >
      <h1 className="text-2xl"> {song?.title} </h1>
      <div> By {song?.artist}, {song?.year}</div>
      {songId === song.Id && <AudioPlayer src={songUrl} />}
    </li>
  );
} 
