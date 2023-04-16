import { useState } from "react";
import { Modal } from "./layouts/Modal";
import { SplitScreen } from "./layouts/SplitScreen";
import { ListRenderer } from "./layouts/ListRenderer";
import { SongItem } from "./components/SongItem";
import { NavBar } from "./components/NavBar";

const songs = [
  {
    Id: "643bee980a62fea680ada970",
    title: "Yellow",
    artist: "Cold Play",
    year: 2009
  },
  {
    Id: "643bef5769adb538379887ad",
    title: "The Scientist",
    artist: "Cold Play",
    year: 2001
  },
  {
    Id: "643bef5e69adb538379887ae",
    title: "Paradise",
    artist: "Cold Play",
    year: 2001
  },
];

export default function App() {
  const [modalVisible, setModalVisible] = useState<boolean>(false);
  return (
    <div className="w-full h-full">
      <NavBar />
      <SplitScreen className="min-h-[600px]" leftWidth={4} rightWidth={8}>
        
        <div className="w-full h-full p-10 bg-orange-500">
          <h1>List of people shown in a Modal:</h1>
          <button
            className="text-white font-bold"
            onClick={() => setModalVisible(true)}
          >
            See List
          </button>
          <Modal
            visible={modalVisible}
            requestToClose={() => setModalVisible(false)}
          >
            <ListRenderer
              ItemComponent={SongItem}
              resourceName="song"
              items={songs}
            />
          </Modal>
        </div>

        <div className="w-full h-full p-10 bg-gray-800">
          <ListRenderer
            ItemComponent={SongItem}
            resourceName="song"
            items={songs}
          />
        </div>
      </SplitScreen>
    </div>
  );
}
