import { FC, ReactNode } from "react";

interface ModalPropsType {
  children: ReactNode;
  visible: boolean;
  requestToClose: () => void;
}

export const Modal: FC<ModalPropsType> = ({
  children,
  visible,
  requestToClose,
}) => {
  if (!visible) return null;
  return (
    <div
      className="fixed inset-0  bg-black/70 flex justify-center items-center"
      onClick={requestToClose}
    >
      <div
        className="relative min-w-[900px] p-5 min-h-[400px] bg-white"
        onClick={(e) => e.stopPropagation()}
      >
        <button onClick={requestToClose} className="text-red-500">Hide Modal</button>
        {children}
      </div>
    </div>
  );
};