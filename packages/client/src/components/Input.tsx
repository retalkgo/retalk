export interface InputProps {
	placeholder?: string;
	onInput?: (value: string) => void;
	value?: string;
	disabled?: boolean;
}

export const Input = (props: InputProps) => (
	<input
		{...props}
		type="text"
		value={props.value}
		disabled={props.disabled}
		onInput={(e) => {
			props.onInput?.((e.target as HTMLInputElement).value);
		}}
		class="inline-block min-h-8 min-w-0 w-[calc(100%_/_3)] inputlike px-3 font-inherit transition-all duration-100 ease-out"
		classList={{
			"opacity-50": props.disabled,
			"focus:w-250px hover:w-250px inputlike-active": !props.disabled,
		}}
	/>
);
