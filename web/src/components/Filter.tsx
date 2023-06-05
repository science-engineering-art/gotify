import React, { useState } from 'react';

interface Music {
  title: string;
  artist: string;
}

const musicList: Music[] = [
  { title: 'Song 1', artist: 'Artist 1' },
  { title: 'Song 1', artist: 'Artist 2' },
  { title: 'Song 2', artist: 'Artist 2' },
  { title: 'Song 3', artist: 'Artist 3' },
  { title: 'Song 4', artist: 'Artist 4' },
];

const MusicFilter: React.FC = () => {
  const [music, setMusic] = useState<Music[]>(musicList);
  const [artist, setGen] = useState<string>("");
  const [title, setTitle] = useState<string>("");

  const handleSearchChange_title= (event: React.ChangeEvent<HTMLInputElement>) => {
    setTitle(event.target.value);
  };

  const handleSearchChange_art= (event: React.ChangeEvent<HTMLInputElement>) => {
    setGen(event.target.value);
  };

  const handleSearchClick = () => {
    const filteredMusic = musicList.filter((m) => {
      const name_match = m.title.toLowerCase().includes(title.toLowerCase());
      const artist_match = m.artist.toLowerCase().includes(artist.toLowerCase())
      
      return name_match && artist_match
    }
    );
    setMusic(filteredMusic);
  };

  return (
    <div>
      <div>
        <label htmlFor="title">Name:</label>
        <input type="text" value={title} onChange={handleSearchChange_title} />
      </div>

      <div>
        <label htmlFor="title"> Artist:</label>
        <input type="text" value={artist} onChange={handleSearchChange_art} />
      </div>

      <div>
       <button onClick={handleSearchClick}>Search</button>
      </div>

      <ul>
        {music.map((m) => (
          <li key={m.title}>
            {m.title} - {m.artist}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default MusicFilter;