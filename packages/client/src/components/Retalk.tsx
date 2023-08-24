import { deepToCamelCase } from "@retalkgo/utils";
import { For, Show, createResource, createSignal } from "solid-js";
import { createStore } from "solid-js/store";

import { useApi } from "../contexts/api";
import { useI18n } from "../i18n";
import { Button } from "./Button";
import { Comment } from "./Comment";
import { Input } from "./Input";
import { Textarea } from "./Textarea";

export function Retalk() {
	const [user, setUser] = createStore<{
		name: string;
		email: string;
		website: string;
	}>({
		name: "",
		email: "",
		website: "",
	});
	const [content, setContent] = createSignal("");
	const i18n = useI18n();
	const api = useApi();
	const [data] = createResource(() => api.getComments());

	return (
		<div class="uno: flex flex-col gap-4.5">
			<div class=":uno: flex gap-3">
				<Input
					placeholder={i18n.name}
					value={user.name}
					onInput={(s) => setUser("name", s)}
				/>
				<Input
					placeholder={i18n.email}
					value={user.email}
					onInput={(s) => setUser("email", s)}
				/>
				<Input
					placeholder={i18n.link}
					value={user.website}
					onInput={(s) => setUser("website", s)}
				/>
			</div>
			<div class=":uno: flex">
				<Textarea
					placeholder={i18n.welcome}
					value={content()}
					onInput={setContent}
				/>
			</div>
			<div class=":uno: flex justify-end">
				<Button>
					<span>{i18n.send}</span>
				</Button>
			</div>
			<Show when={!!data()}>
				<For
					each={
						// TODO
						data()!.data.data
					}
				>
					{(comment) => (
						<Comment
							body={comment.body}
							author={deepToCamelCase(comment.author)}
							createdAt={comment.created_at}
							id={comment.id}
							path={comment.path}
						/>
					)}
				</For>
			</Show>
		</div>
	);
}
