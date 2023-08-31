interface TextareaProps {
	placeholder?: string;
	onInput?: (value: string) => void;
	value?: string;
	disabled?: boolean;
}

export const Textarea = (props: TextareaProps) => (
	<textarea
		{...props}
		value={props.value}
		disabled={props.disabled}
		onInput={(e) => {
			props.onInput?.((e.target as HTMLTextAreaElement).value);
		}}
		class=":uno: block min-h-20 w-full inputlike p-3 font-inherit"
		classList={{
			":uno: opacity-50": props.disabled,
			":uno: inputlike-active": !props.disabled,
		}}
	/>
);
