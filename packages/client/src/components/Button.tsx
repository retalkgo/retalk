import type { ParentComponent } from "solid-js";
import { children } from "solid-js";

export interface ButtonProps {}

export const Button: ParentComponent<ButtonProps> = (props) => {
  const resolvedChildren = children(() => props.children);

  return (
    <button class=":uno: inline-block min-h-9 cursor-pointer rounded-4.5 border-none bg-primary px-7.5 py-2.5 text-3.6 font-semibold text-white transition duration-animation hover:shadow-active">
      {resolvedChildren()}
    </button>
  );
};
