import {useEffect, useState} from "react";
import {Quote} from "../models/quote";

export function useQuotes() {
	const [quotes, setQuotes] = useState<Quote[]>();
	const [loading, setLoading] = useState(true);
	const [refresh, forceRefresh] = useState({});

	useEffect(() => {
		async function fetchQuotes() {
			const url = `${process.env.EXPO_PUBLIC_API_URL}/quotes/`;

			const response = await fetch(url);
			const quotes = await response.json() as Quote[];
			setQuotes(quotes);
			setLoading(false);
		}

		fetchQuotes();
	}, [refresh]);

	function updateQuote(quote: Quote) {
		setQuotes(quotes => quotes?.map(q => q.id === quote.id ? quote : q));
	}

	function refreshQuotes() {
		setLoading(true);
		forceRefresh({});
	}

	return [quotes, loading, updateQuote, refreshQuotes] as const;
}