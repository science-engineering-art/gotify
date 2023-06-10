import { createSlice, createAsyncThunk } from '@reduxjs/toolkit'
import { client } from '../api/client'
import { Metadata, SongDTO } from '../app/metadata';

const initialState: {
  Id: string,
  url: string,
  filter: Metadata,
  playlist: SongDTO[],
  status: 'idle' | 'loading' | 'succeeded' | 'failed',
  error: string | null
} = {
  Id: '',
  url: '',
  filter: { 
    title: "", 
    artist: "", 
    album: "", 
    genre: "" 
  },
  playlist: [],
  status: 'idle',
  error: null
};

export const songsSlice = createSlice({
  name: 'songs',
  initialState,
  reducers: {
    selectedSongId(state, action) {
      state.Id = action.payload
    },
    selectedSongURL(state, action) {
      state.url = action.payload
    },
    selectedPlaylist(state, action) {
      state.playlist = action.payload
    },
    filterByArtist(state, action) {
      state.filter.artist = action.payload
    },
    filterByAlbum(state, action) {
      state.filter.album = action.payload
    },
    filterByGenre(state, action) {
      state.filter.genre = action.payload
    },
    filterByTitle(state, action) {
      state.filter.title = action.payload
    },
  },
  extraReducers(builder) {
      builder
        .addCase(songFilter.pending, (state, _) => {
          state.status = 'loading';
        })
        .addCase(songFilter.fulfilled, (state, action) => {
          state.playlist = action.payload;
          state.status = 'idle';
        })
  },
})

export const uploadSong = createAsyncThunk('songs/uploadSong', async (song: File) => {
  var form = new FormData();
  form.append('file', song);  
  client.post('song', form);
});

export const songFilter = createAsyncThunk('songs/songFilter', async (filter: Metadata) => {

  const data = {
    'artist': filter.artist,
    'album': filter.album,
    'genre': filter.genre,
    'title': filter.title,
  }
  // client.post('songs', JSON.stringify(data));

  const resp = await fetch('http://api.gotify.com/songs', {
    headers: {
      'Content-Type': 'application/json'
    },
    method: 'POST',
    body: JSON.stringify(data)
  }).then(res => res.json())
    .catch(e => console.log(e))

  return resp.data.songs
});

export const { 
  selectedSongId, 
  selectedSongURL, 
  selectedPlaylist,
  filterByAlbum, 
  filterByArtist, 
  filterByGenre, 
  filterByTitle 
} = songsSlice.actions

export default songsSlice.reducer