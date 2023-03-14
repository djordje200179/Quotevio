export interface Quote {
	id: string;

	text: string;
	author: string;

	created: Date;

	likes: number;
	dislikes: number;
}