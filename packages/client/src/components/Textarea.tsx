interface TextareaProps {
	placeholder?: string;
}

export const Textarea = (props: TextareaProps) => (
	<textarea
		{...props}
		class=":uno: block min-h-20 w-full p-3 font-inherit inputlike"
	/>
);
