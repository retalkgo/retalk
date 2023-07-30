export interface Author {
	created_at: string;
	email: string;
	id: number;
	is_admin: boolean;
	link: string;
	name: string;
}

export interface Comment {
	author: Author;
	body: string;
	created_at: String;
	id: number;
	path: string;
}

export interface ApiResult<T = any> {
	data: T;
	msg: string;
	success: boolean;
}
