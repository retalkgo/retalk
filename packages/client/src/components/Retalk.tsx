import { createSignal } from "solid-js";
import { createStore } from "solid-js/store";

import useI18n from "../i18n";
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
			<Comment
				body="Hiiiiiiiiiiiiiiiii"
				author={{
					name: "Ray",
					email: "d571021199a1b1d1962fd1f4a7879ffc",
					isAdmin: true,
					createdAt: "",
					id: 1,
					link: "",
				}}
				createdAt="2023年6月9日 17:41"
				id={1}
				path="/"
			/>
		</div>
	);
}
