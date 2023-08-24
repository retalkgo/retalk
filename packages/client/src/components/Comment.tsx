import type { Comment as CommentType } from "@retalkgo/api";
import type { DeepToCamelCase } from "@retalkgo/utils";
import { Match, Show, Switch, createMemo } from "solid-js";

import { useOptions } from "../contexts/options";
import { useI18n } from "../i18n";

export type CommentProps = DeepToCamelCase<CommentType>;

export function Comment(props: CommentProps) {
	const options = useOptions();
	const i18n = useI18n();

	const avatar = createMemo(
		() => `${options.gravatarProxy}${props.author.email}`,
	);

	return (
		<div class="flex gap-3">
			<img
				class="h-10 max-w-full w-10 rounded-3 shadow-avatar"
				src={avatar()}
			/>
			<div class="w-full flex-1 rounded-3 px-8 pb-3 pt-4 text-13px text-sm shadow-comment">
				<div class="flex items-center gap-1 font-600">
					<Switch
						fallback={
							<span class="text-lg text-primary">{props.author.name}</span>
						}
					>
						<Match when={props.author.link}>
							<a
								href={props.author.link}
								class="text-lg text-primary decoration-none"
							>
								{props.author.name}
							</a>
						</Match>
					</Switch>
					<Show when={props.author.isAdmin}>
						<div class="scale-80 rounded bg-primary/18 p-1 text-sm">
							<span class="text-primary">{i18n.admin}</span>
						</div>
					</Show>
				</div>
				<span class="text-xs font-light text-second/56">{props.createdAt}</span>
				<div class=":uno:">{props.body}</div>
				<div class="mt-1 flex justify-end">
					<span
						class="cursor-pointer text-primary"
						onClick={() => {
							// eslint-disable-next-line no-alert
							alert("Reply");
						}}
					>
						{i18n.reply}
					</span>
				</div>
			</div>
		</div>
	);
}
