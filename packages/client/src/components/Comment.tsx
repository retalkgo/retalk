import type { Comment as CommentType } from "@retalkgo/schema";
import type { DeepToCamelCase } from "@retalkgo/utils";
import { Show, createMemo } from "solid-js";

import { useOptions } from "./Options";

export type CommentProps = DeepToCamelCase<CommentType>;

export function Comment(props: CommentProps) {
	const options = useOptions();

	const avatar = createMemo(
		() => `${options.gravatarProxy}${props.author.email}`,
	);

	return (
		<div class=":uno: flex gap-3">
			<img
				class=":uno: h-10 max-w-full w-10 rounded-3 shadow-avatar"
				src={avatar()}
			/>
			<div class=":uno: w-full flex-1 rounded-3 px-8 pb-3 pt-4 text-13px text-sm font-600 shadow-comment">
				<div class=":uno: flex items-center gap-1">
					<div class=":uno: text-primary">{props.author.name}</div>
					<Show when={props.author.isAdmin}>
						<div class=":uno: scale-80 rounded bg-primary/18 p-1.5 text-xs">
							<span class=":uno: text-primary">管理员</span>
						</div>
					</Show>
				</div>
			</div>
		</div>
	);
}
