export interface IComment {
	id: string;
	comment: string;
	userId: number;
	snippet_id: string;
}
export const comments: Array<IComment> = [];
