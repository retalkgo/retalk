import { createSignal } from "solid-js";
import { createStore } from "solid-js/store";

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

	return (
		<div class="uno: flex flex-col gap-4.5">
			<div class=":uno: flex gap-3">
				<Input
					placeholder="昵称"
					value={user.name}
					onInput={(s) => setUser("name", s)}
				/>
				<Input
					placeholder="邮箱"
					value={user.email}
					onInput={(s) => setUser("email", s)}
				/>
				<Input
					placeholder="网站"
					value={user.website}
					onInput={(s) => setUser("website", s)}
				/>
			</div>
			<div class=":uno: flex">
				<Textarea
					placeholder="说点什么？"
					value={content()}
					onInput={setContent}
				/>
			</div>
			<div class=":uno: flex justify-end">
				<Button>
					<span>按钮</span>
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
