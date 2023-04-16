import { Children, ReactNode } from "react";

export const SplitScreen = ({
  leftWidth,
  rightWidth,
  className,
  children,
}: {
  className?: string;
  leftWidth?: number;
  rightWidth?: number;
  children: ReactNode;
}) => {
  const [left, right] = Children.toArray(children);
  return (
    <div className={`flex w-full ${className}`}>
      <div style={{ flex: `${leftWidth}` }}>{left}</div>
      <div style={{ flex: `${rightWidth}` }}>{right}</div>
    </div>
  );
};

SplitScreen.defaultProps = {
  className: "",
  leftWidth: 6,
  rightWidth: 6,
};