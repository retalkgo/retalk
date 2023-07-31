import { Button } from "./Button";
import { Comment } from "./Comment";
import { Input } from "./Input";
import { Textarea } from "./Textarea";

export const Retalk = () => (
	<div class="uno: w-125 flex flex-col gap-4.5">
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
			author={
				{
					name: "Redish101",
					// TODO
				} as any
			}
			createdAt={Date.now().toString()}
			id={1}
			path="/"
		/>
	</div>
);
