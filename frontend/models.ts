export interface Date {
	day: number;
	month: number;
	year: number;
}

export interface Quote {
	id: string;

	text: string;
	author: string;

	created: Date;

	likes: number;
	dislikes: number;
}