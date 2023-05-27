import { configureStore } from '@reduxjs/toolkit'
import songsReducer from '../features/songsSlice'

export const store = configureStore({
  reducer: {
    songs: songsReducer
  }
})

export type AppDispatch = typeof store.dispatch;