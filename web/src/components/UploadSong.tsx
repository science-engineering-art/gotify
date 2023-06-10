import { ChangeEvent, MouseEvent, useState } from 'react';
import { Modal } from '../layouts/Modal';
import { useDispatch, useSelector } from 'react-redux';
import { songFilter, uploadSong } from '../features/songsSlice';
import { AppDispatch } from '../app/store';
import { Metadata } from '../app/metadata';

export const UploadSong = () => {  
  const dispatch = useDispatch<AppDispatch>()

  const filter = useSelector((state: { songs: { filter: Metadata } }) => state.songs.filter);

  const [song, setSong] = useState<File>();
  const [modalVisible, setModalVisible] = useState<boolean>(false);

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    if (e.target.files === null) 
      return;
    setSong(e.target.files[0]);
  }

  const handleSubmit = async (e: MouseEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (!song) return;
    dispatch(uploadSong(song));
    setModalVisible(false);
    
    setTimeout(()=>{
      dispatch(songFilter(filter));
    }, 500)
  }

  return (
    <div className="w-full h-full p-10">
      <button
        className="bg-green-500 hover:bg-opacity-80 text-black font-bold py-1 px-2 rounded-lg"
        onClick={() => setModalVisible(true)}
      >
        Add Song
      </button>
      <Modal
        visible={modalVisible}
        requestToClose={() => setModalVisible(false)}
      >
        <div>
          <form onSubmit={handleSubmit}>
            <input 
              title='File'
              type='file' 
              onChange={handleChange}
              />
            <input
              className='bg-green-500 hover:bg-opacity-80 text-black font-bold py-1 px-2 rounded-lg' 
              type='submit' 
              value={'Submit'}
            />
          </form>
        </div>
      </Modal>
    </div>
  );
}
