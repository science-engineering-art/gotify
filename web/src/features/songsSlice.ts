import { createSlice, createAsyncThunk } from '@reduxjs/toolkit'
import { client } from '../api/client'
import { Metadata } from '../app/metadata';

const initialState: {
  Id: string,
  url: string,
  status: 'idle' | 'loading' | 'succeeded' | 'failed',
  error: string | null
} = {
  Id: '',
  url: '',
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
  }
})

export const uploadSong = createAsyncThunk('songs/uploadSong', async (song: File) => {
  var form = new FormData();
  form.append('file', song);  
  client.post('song', form);
});

export const songFilter = createAsyncThunk('songs/songFilter', async (filter: Metadata) => {
  client.post('songs', {
    'title': filter.title,
    'artist': filter.artist,
    'album': filter.album,
    'genre': filter.genre
  });
});

export const { selectedSongId, selectedSongURL } = songsSlice.actions

export default songsSlice.reducer