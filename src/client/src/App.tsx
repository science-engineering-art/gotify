import { useState } from "react";
import { Modal } from "./layouts/Modal";
import { SplitScreen } from "./layouts/SplitScreen";
import { ListRenderer } from "./layouts/ListRenderer";
import { SongItem } from "./components/SongItem";

const LeftScreen = () => (
  <div className="bg-orange-500 text-white h-full">Right Screen</div>
);
const RightScreen = () => (
  <div className="bg-blue-700 text-white h-full">Left Screen</div>
);

type PersonItemPropsType = {
  person: {
    fullName: string;
    age: number;
    occupation: string;
  };
};
type ProductItemPropsType = {
  product: {
    name: string;
    price: number;
    quantity: number;
  };
};

const ProductItem = ({ product }: ProductItemPropsType) => (
  <li className="w-full p-5 bg-orange-200 my-2">
    <h1>{product.name}</h1>
    <div>Price: ${product?.price}</div>
    <div> Quantity: {product?.quantity}</div>
  </li>
);

const PersonItem = ({ person }: PersonItemPropsType) => (
  <li className="w-full p-5 bg-blue-300 my-2">
    <h1>{person.fullName}</h1>
    <div>age: {person?.age}</div>
    <div> occupation: {person?.occupation}</div>
  </li>
);

const people = [
  {
    fullName: "John Doe",
    age: 35,
    occupation: "AI Engineer",
  },
  {
    fullName: "Baqir Nekfar",
    age: 23,
    occupation: "Frontend Engineer",
  },
  {
    fullName: "Sara Smith",
    age: 27,
    occupation: "UI/UX designer",
  },
];
const products = [
  {
    name: "Dinning Table",
    price: 100,
    quantity: 10,
  },
  {
    name: "Programming Chair",
    price: 200,
    quantity: 20,
  },
  {
    name: "Sofa",
    price: 1000,
    quantity: 100,
  },
];

const songs = [
  {
    title: "Yellow",
    artist: "Cold Play",
    year: 2009
  },
  {
    title: "The Scientist",
    artist: "Cold Play",
    year: 2001
  },
  {
    title: "Paradise",
    artist: "Cold Play",
    year: 2001
  },
];

export default function App() {
  const [modalVisible, setModalVisible] = useState<boolean>(false);
  return (
    <div className="w-full h-full">
      <SplitScreen className="min-h-[600px]" leftWidth={4} rightWidth={8}>
        
        <div className="w-full h-full p-10">
          <h1>List of people shown in a Modal:</h1>
          <button
            className="text-orange-500 font-bold"
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

        <div className="w-full h-full p-10">
          <h1>List of people:</h1>
          <ListRenderer
            ItemComponent={PersonItem}
            resourceName="person"
            items={people}
          />
          <h1>List of products:</h1>
          <ListRenderer
            ItemComponent={ProductItem}
            resourceName="product"
            items={products}
          />
        </div>
      </SplitScreen>
    </div>
  );
}
