import { ChangeEvent, MouseEvent, useState } from 'react';
import { Modal } from '../layouts/Modal';

export const UploadSong = () => {  
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

    var form = new FormData();
    form.append("file", song);

    await fetch('http://localhost:5000/song', {
      method: 'POST',
      body: form
    }).then(res => res.json())
      .catch(e => console.log(e))
  }

  return (
    <div className="w-full h-full p-10">
      <button
        className="bg-green-500 hover:bg-gray-300 text-black font-bold py-1 px-2 rounded-lg"
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
            <input type='submit' value={'Submit'}/>
          </form>
        </div>
      </Modal>
    </div>
  );
}
