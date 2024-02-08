import { Quote } from "../models/quote";
import { StyleSheet, View } from "react-native";
import { MD3LightTheme as DefaultTheme, Chip } from "react-native-paper";

interface Props {
	quote: Quote;
	onQuoteUpdate: (quote: Quote) => void;
}

const styles = StyleSheet.create({
	container: {
		flexDirection: "row"
	},
	actionText: {
		color: DefaultTheme.colors.onPrimary,
		fontSize: 14
	},
	likesContainer: {
		backgroundColor: "#6495ed",
		marginRight: 3,
		width: 75
	},
	dislikesContainer: {
		backgroundColor: "#ff7f50",
		marginLeft: 3,
		width: 75
	}
});

export default function QuoteActions({ quote, onQuoteUpdate }: Props) {
	async function like() {
		const url = `${process.env.EXPO_PUBLIC_API_URL}/quotes/${quote.id}/like`;

		const response = await fetch(url, { method: "PATCH" });
		const newQuote: Quote = await response.json();

		newQuote.liked = true;
		onQuoteUpdate(newQuote);
	}

	async function dislike() {
		const url = `${process.env.EXPO_PUBLIC_API_URL}/quotes/${quote.id}/dislike`;

		const response = await fetch(url, { method: "PATCH" });
		const newQuote: Quote = await response.json();

		newQuote.disliked = true;
		onQuoteUpdate(newQuote);
	}

	return (
		<View style={styles.container}>
			<Chip icon="thumb-up" onPress={like} disabled={quote.liked} selected={quote.liked}
				  textStyle={styles.actionText} style={styles.likesContainer}>
				{quote.likes}
			</Chip>

			<Chip icon="thumb-down" onPress={dislike} disabled={quote.disliked} selected={quote.disliked}
			      textStyle={styles.actionText} style={styles.dislikesContainer}>
				{quote.dislikes}
			</Chip>
		</View>
	);
}