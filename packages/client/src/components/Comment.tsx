export interface CommentProps {
	avatar: string;
	nickname: string;
	isAdmin: boolean;
	time: string;
	content: string;
}

export const Comment = (_props: CommentProps) => (
	<div class=":uno: inline-block min-h-9 cursor-pointer rounded-4.5 border-none bg-primary px-7.5 py-2.5 text-3.6 font-semibold text-white transition duration-animation hover:shadow-active" />
);
