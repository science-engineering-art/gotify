import { useEffect, useState } from "react";
import { ListRenderer } from "../layouts/ListRenderer";
import { SongItem } from "./SongItem";
import React from "react";
import { songFilter } from '../features/songsSlice';
import { useDispatch, useSelector } from "react-redux";
import { AppDispatch } from "../app/store";
import { Metadata } from "../app/metadata";

type SongDTO = {
  id: string;
  title: string;
  artist: string;
  year: number;
}

export const SongsList: React.FC = () => {
  const playlist = useSelector((state: { songs: { playlist: SongDTO[] } }) => state.songs.playlist);
  // const dispatch = useDispatch<AppDispatch>()
  
  // 
  
  // dispatch(songFilter(filter));


  // const [songs, setSongs ]= useState<SongDTO[]>(playlist);

  // useEffect(() => {
  //   setSongs(playlist);
  // }, [])

  return (
    <div className="w-full h-full p-10 bg-gray-500">
      {playlist && <ListRenderer
        ItemComponent={SongItem}
        resourceName="song"
        items={playlist}
      />}
    </div>
  );
}