import React, { useState, MouseEvent } from 'react';
import { useDispatch } from 'react-redux';
import { AppDispatch } from '../app/store';
import { songFilter } from '../features/songsSlice';


export const SongFilter: React.FC = () => {
  const [title, setTitle]   = useState<string>("");
  const [artist, setArtist] = useState<string>("");
  const [album, setAlbum]   = useState<string>("");
  const [genre, setGenre]   = useState<string>("");
  
  const dispatch = useDispatch<AppDispatch>();
  
  const handleSubmit = async (e: MouseEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (title=="" && artist == "" && album == "" && genre == "") 
      return;
    
    dispatch(songFilter({
      title: title,
      artist: artist,
      album: album,
      genre: genre
    }));
  }

  return (
    <form
      onSubmit={handleSubmit} 
      className="bg-white shadow-lg rounded-md p-5 md:p-10 flex flex-col w-11/12 max-w-lg"
    >
      <label htmlFor="text" className="block text-sm font-medium text-gray-700">
        Title
      </label>
      <input 
        type="text" 
        name="title" 
        id="title" 
        placeholder='Title' 
        className="mt-1 block w-full py-2 px-3 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
        onChange={e => setTitle(e.target.value)}
      />

      <label htmlFor="text" className="block text-sm font-medium text-gray-700">
        Artist
      </label>
      <input 
        type="text" 
        name="artist" 
        id="artist" 
        placeholder='Artist' 
        className="mt-1 block w-full py-2 px-3 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
        onChange={e => setArtist(e.target.value)}
      />

      <label htmlFor="text" className="block text-sm font-medium text-gray-700">
        Album
      </label>
      <input 
        type="text" 
        name="album" 
        id="album" 
        placeholder='Album' 
        className="mt-1 block w-full py-2 px-3 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
        onChange={e => setAlbum(e.target.value)}
      />

      <label htmlFor="text" className="block text-sm font-medium text-gray-700">
        Genre
      </label>
      <input 
        type="text" 
        name="genre" 
        id="genre" 
        placeholder='Genre' 
        className="mt-1 block w-full py-2 px-3 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
        onChange={e => setGenre(e.target.value)}
      />

      <button type="submit" className="mt-4 bg-indigo-600 text-white py-2 px-4 rounded-md hover:bg-indigo-500">
        Filter
      </button>
    </form>
  );    
};  
