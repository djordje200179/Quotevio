export interface Quote {
	id: string;

	text: string;
	author: string;

	created_at: Date;

	likes: number;
	dislikes: number;

	liked?: boolean;
	disliked?: boolean;
}