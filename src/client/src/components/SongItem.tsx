
export type SongItemPropsType = {
  song: {
    artist: string;
    title: number;
    year: string;
  };
};

export const SongItem = ({ song }: SongItemPropsType) => (
  <li className="w-full p-5 bg-blue-300 my-2">
    <h1 className="text-2xl"> {song.title} </h1>
    <div> By {song.artist}, {song.year}</div>
  </li>
);
