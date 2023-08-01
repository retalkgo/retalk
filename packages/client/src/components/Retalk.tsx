import { Button } from "./Button";
import { Comment } from "./Comment";
import { Input } from "./Input";
import { Textarea } from "./Textarea";

export const Retalk = () => (
	<div class="uno: flex flex-col gap-4.5">
		<div class=":uno: flex gap-3">
			<Input placeholder="昵称" />
			<Input placeholder="邮箱" />
			<Input placeholder="网站" />
		</div>
		<div class=":uno: flex">
			<Textarea placeholder="说点什么？" />
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
