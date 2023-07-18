export interface InputProps {
  placeholder?: string;
}

export const Input = (props: InputProps) => (
  <input
    {...props}
    type="text"
    class=":uno: click:w-250px inline-block min-h-8 min-w-0 px-3 inputlike"
  />
);
