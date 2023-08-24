import type { ParentComponent } from "solid-js";
import { children } from "solid-js";

export interface ButtonProps {
	disabled?: boolean;
	onClick?: () => void;
}

export const Button: ParentComponent<ButtonProps> = (props) => {
	const resolvedChildren = children(() => props.children);

	return (
		<button
			disabled={props.disabled}
			onClick={() => {
				props.onClick?.();
			}}
			class="inline-block min-h-9 rounded-4.5 border-none bg-primary px-7.5 py-2.5 text-3.6 font-semibold text-white transition duration-animation"
			classList={{
				"opacity-50": props.disabled,
				"hover:shadow-active cursor-pointer": !props.disabled,
			}}
		>
			{resolvedChildren()}
		</button>
	);
};
