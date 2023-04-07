import { ChangeEvent, MouseEvent, useState } from 'react';

export const UploadSong = () => {  
  const [song, setSong] = useState<File>();

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
  );
}
