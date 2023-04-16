import { createSlice, createAsyncThunk } from '@reduxjs/toolkit'
import { client } from '../api/client'

const initialState: {
  Id: string,
  status: 'idle' | 'loading' | 'succeeded' | 'failed',
  error: string | null
} = {
  Id: '',
  status: 'idle',
  error: null
};

export const songsSlice = createSlice({
  name: 'songs',
  initialState,
  reducers: {
    selectedSong: (state, action) => {
      state.Id = action.payload;
    }
  }
})

export const uploadSong = createAsyncThunk('songs/uploadSong', async (song: File) => {
  var form = new FormData();
  form.append("file", song);  
  client.post('song', form);
});

export const { selectedSong } = songsSlice.actions

export default songsSlice.reducer