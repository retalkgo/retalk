export interface InputProps {
  placeholder?: string;
}

export function Input(props: InputProps) {

  return (
    <input
      {...props}
      type="text"
      class=":uno: inline-block min-h-8 min-w-0 px-3 click:w-250px inputlike"
    />
  );
}
