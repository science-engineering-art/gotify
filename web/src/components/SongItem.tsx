import { useDispatch, useSelector } from "react-redux";
import { selectedSongId, selectedSongURL } from "../features/songsSlice";
import { AudioPlayer } from "./AudioPlayer";
import { MouseEvent } from "react";

export type SongItemPropsType = {
  song: {
    id: string;
    artist: string;
    title: number;
    year: string;
  };
};

export const SongItem = ({ song }: SongItemPropsType) =>{
  const dispatch = useDispatch();
  const songId = useSelector((state: { songs: { id: string } }) => state.songs.id);
  const songUrl = useSelector((state: { songs: { url: string } }) => state.songs.url);

  const handleSelectedSong = async (_: MouseEvent<HTMLLIElement>) => {
    console.log(song)
    dispatch(selectedSongId(song.id))
    dispatch(selectedSongURL(`http://api.gotify.com/song/${song.id}`))
  }

  return (
    <li 
      className="w-full p-5 bg-green-500 my-2 rounded-3xl hover:bg-opacity-60" 
      onClick={handleSelectedSong}
    >
      <h1 className="text-2xl"> {song?.title} </h1>
      <div> By {song?.artist}, {song?.year}</div>
      {songId === song.id && <AudioPlayer src={songUrl} />}
    </li>
  );
} 
