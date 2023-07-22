export interface InputProps {
	placeholder?: string;
}

export const Input = (props: InputProps) => (
	<input
		{...props}
		type="text"
		class=":uno: inline-block min-h-8 min-w-0 w-[calc(100%_/_3)] px-3 font-inherit transition-all duration-100 ease-out focus:w-250px hover:w-250px inputlike"
	/>
);
