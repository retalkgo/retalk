import { deepToCamelCase } from "@retalkgo/utils";
import { For, Show, createResource, createSignal } from "solid-js";
import { createStore } from "solid-js/store";

import { useApi } from "../contexts/api";
import { usePath } from "../hooks/use-path";
import { useI18n } from "../i18n";
import { Button } from "./Button";
import { Comment } from "./Comment";
import { Input } from "./Input";
import { Textarea } from "./Textarea";

export function Retalk() {
	const i18n = useI18n();
	const api = useApi();
	const path = usePath();
	const [user, setUser] = createStore<{
		name: string;
		email: string;
		link: string;
	}>({
		name: "",
		email: "",
		link: "",
	});
	const [content, setContent] = createSignal("");
	const [loading, setLoading] = createSignal(false);
	const [data, { refetch }] = createResource(path, () =>
		api.getComments(path()),
	);
	async function handleSubmit() {
		setLoading(true);
		await api.createComment({
			path: path(),
			body: content(),
			...user,
		});
		setLoading(false);
		await refetch();
	}

	return (
		<div class="uno: flex flex-col gap-4.5">
			<div class="flex gap-3">
				<Input
					disabled={loading()}
					placeholder={i18n.name}
					value={user.name}
					onInput={(s) => setUser("name", s)}
				/>
				<Input
					disabled={loading()}
					placeholder={i18n.email}
					value={user.email}
					onInput={(s) => setUser("email", s)}
				/>
				<Input
					disabled={loading()}
					placeholder={i18n.link}
					value={user.link}
					onInput={(s) => setUser("link", s)}
				/>
			</div>
			<div class="flex">
				<Textarea
					disabled={loading()}
					placeholder={i18n.welcome}
					value={content()}
					onInput={setContent}
				/>
			</div>
			<div class="flex justify-end">
				<Button disabled={loading()} onClick={handleSubmit}>
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
					{(comment) => <Comment {...deepToCamelCase(comment)} />}
				</For>
			</Show>
		</div>
	);
}
