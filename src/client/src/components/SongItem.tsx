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

    // get the selected song and create an URL object
    await fetch(`http://localhost:5000/api/song/${song.Id}`)
      .then(response => response.json())
      .then(data => {
        // Decodificar la cadena base64
        const decodedString = atob(data.data.rawSong)

        // Convertir la cadena decodificada en un arreglo de bytes
        const bytes = new Uint8Array(decodedString.length)
        for (let i = 0; i < decodedString.length; i++) {
          bytes[i] = decodedString.charCodeAt(i)
        }
        
        // Crear un objeto Blob a partir del contenido de la canción
        const blob = new Blob([bytes], { type: "audio/mp3" })
        
        // Crear un objeto URL a partir del objeto Blob
        dispatch(selectedSongURL(
          URL.createObjectURL(blob)
        ))

        // remove the last selected song url
        if (songUrl){
          URL.revokeObjectURL(songUrl)
        }
      })
      .catch(error => console.error(error))
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
