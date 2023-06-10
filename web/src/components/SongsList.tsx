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
  const dispatch = useDispatch<AppDispatch>()
  const filter = useSelector((state: { songs: { filter: Metadata } }) => state.songs.filter);
  
  const playlist = useSelector((state: { songs: { playlist: SongDTO[] } }) => state.songs.playlist);
  
  useEffect(() => {
    dispatch(songFilter(filter));
  }, [])

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