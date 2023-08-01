interface TextareaProps {
	placeholder?: string;
	onInput?: (value: string) => void;
	value?: string;
}

export const Textarea = (props: TextareaProps) => (
	<textarea
		{...props}
		value={props.value}
		onInput={(e) => {
			props.onInput?.((e.target as HTMLTextAreaElement).value);
		}}
		class=":uno: block min-h-20 w-full p-3 font-inherit inputlike"
	/>
);
