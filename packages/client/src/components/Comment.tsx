import type { Comment as CommentType } from "@retalkgo/schema";
import type { DeepToCamelCase } from "@retalkgo/utils";
import { Show } from "solid-js";

export type CommentProps = DeepToCamelCase<CommentType>;

export const Comment = (props: CommentProps) => (
	<div class=":uno: flex gap-3">
		<div class=":uno: h-10 w-10 rounded-3 pr-3 shadow-avatar">
			{/* TODO */}
			{/* <img src={props.avatar} /> */}
		</div>
		<div class=":uno: w-full rounded-3 px-8 pb-3 pt-4 text-13px text-sm font-600 shadow-comment">
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
