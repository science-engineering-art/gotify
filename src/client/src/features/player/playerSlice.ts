import { createSlice } from '@reduxjs/toolkit'

export const playerSlice = createSlice({
  name: 'player',
  initialState: {
    Id: ""
  },
  reducers: {
    playSong: (state, action) => {
      state.Id = action.payload;
    }
  }
})

export const playSongAsync = (Id: string) => (dispatch: (action: any) => void) => {
  setTimeout(() => {
    dispatch(playSong(Id))
  }, 1000)
}

// const fetchSongById = (songId: string) => {
//   // the inside "thunk function"
//   return async (
//         dispatch: (action: any) => void, 
//         getState: any
//     ) => {
//       try {
//         // make an async call in the thunk
//         const user = await userAPI.fetchById(userId)
//         // dispatch an action when we get the response back
//         dispatch(userLoaded(user))
//       } catch (err) {
//         // If something went wrong, handle it here
//       }
//     }
// }

export const { playSong } = playerSlice.actions

export default playerSlice.reducer