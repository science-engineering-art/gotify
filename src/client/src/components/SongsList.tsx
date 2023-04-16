import { useEffect, useState } from "react";
import { ListRenderer } from "../layouts/ListRenderer";
import { SongItem } from "./SongItem";
import { client } from "../api/client";

type SongDTO = {
    Id: string;
    title: string;
    artist: string;
    year: number;
}

export const SongsList = () => {
  const [songs, setSongs ]= useState<SongDTO[]>([
    {
      Id: "643bee980a62fea680ada970",
      title: "Yellow",
      artist: "Cold Play",
      year: 2009
    },
    {
      Id: "643bef5769adb538379887ad",
      title: "The Scientist",
      artist: "Cold Play",
      year: 2001
    },
    {
      Id: "643bef5e69adb538379887ae",
      title: "Paradise",
      artist: "Cold Play",
      year: 2001
    },
  ]);

  const getSongs = async () => {
    const resp = await fetch(`http://localhost:5000/api/songs`, {
      method: 'GET'
    }).then(res => res.json())
      .catch(e => console.log(e))
    setSongs(resp.data.data)
  }

  useEffect(() => {
    getSongs();
  }, [])

  return (
    <div className="w-full h-full p-10 bg-gray-800">
      {songs && <ListRenderer
        ItemComponent={SongItem}
        resourceName="song"
        items={songs}
      />}
    </div>
  );
}