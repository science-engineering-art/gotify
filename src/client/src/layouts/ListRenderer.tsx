import { FC } from "react";

interface ListRendererPropsType {
  items: unknown[];
  resourceName: string;
  ItemComponent: any;
}

export const ListRenderer: FC<ListRendererPropsType> = ({
  items,
  resourceName,
  ItemComponent,
}) => {
  return (
    <ul>
      {items.map((item, i) => (
        <ItemComponent key={i} {...{ [resourceName]: item }} />
      ))}
    </ul>
  );
};