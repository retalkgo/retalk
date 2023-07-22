export interface InputProps {
  placeholder?: string;
}

export const Input = (props: InputProps) => (
  <input
    {...props}
    type="text"
    class=":uno: hover:w-250px focus:w-250px inline-block min-h-8 min-w-0 w-[calc(100%_/_3)] transition-all ease-out duration-100 px-3 inputlike"
  />
);
