import { useEffect, useState } from "react";
import { ListRenderer } from "../layouts/ListRenderer";
import { SongItem } from "./SongItem";
import React from "react";

type SongDTO = {
  id: string;
  title: string;
  artist: string;
  year: number;
}

export const SongsList: React.FC = () => {
  const [songs, setSongs ]= useState<SongDTO[]>();

  const getSongs = async () => {
    const resp = await fetch(`http://api.gotify.com/songs`, {
      method: 'POST',
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({})
    }).then(res => res.json())
      .catch(e => console.log(e))
    setSongs(resp.data.songs)
  }
  
  useEffect(() => {
    getSongs();
  }, [])

  return (
    <div className="w-full h-full p-10 bg-gray-500">
      {songs && <ListRenderer
        ItemComponent={SongItem}
        resourceName="song"
        items={songs}
      />}
    </div>
  );
}