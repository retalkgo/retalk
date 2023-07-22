interface TextareaProps {
  placeholder?: string;
}

export const Textarea = (props: TextareaProps) => (
  <textarea {...props} class=":uno: block min-h-20 p-3 w-full inputlike" />
);
