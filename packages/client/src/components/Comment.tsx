export interface CommentProps {
	avatar: string;
	nickname: string;
	isAdmin: boolean;
	time: string;
	content: string;
}

export const Comment = (props: CommentProps) => (
	<div class=":uno: flex gap-3">
		<div class="h-10 w-10 rounded-3 pr-3 shadow-avatar">
			<img src={props.avatar} />
		</div>
		<div class="w-full rounded-3 px-8 pb-3 pt-4 text-xs font-600 shadow-comment">
			<div class="text-primary">{props.nickname}</div>
		</div>
	</div>
);
