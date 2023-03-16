import { Quote } from "../models";
import { StyleSheet, View } from "react-native";
import { MD3LightTheme as DefaultTheme, Chip } from "react-native-paper";
import { useState } from "react";

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
	const [liked, setLiked]       = useState(false);
	const [disliked, setDisliked] = useState(false);

	function like() {
		const url = `http://192.168.1.26:8080/quotes/${quote.id}/like`;

		fetch(url, { method: "POST" })
			.then(response=>response.json())
			.then(newQuote => {
				setLiked(true);
				onQuoteUpdate(newQuote);
			})
			.catch(console.error);
	}

	function dislike() {
		const url = `http://192.168.1.26:8080/quotes/${quote.id}/dislike`;

		fetch(url, { method: "POST" })
			.then(response=>response.json())
			.then(newQuote => {
				setDisliked(true);
				onQuoteUpdate(newQuote);
			})
			.catch(console.error);
	}

	return (
		<View style={styles.container}>
			<Chip icon="thumb-up" onPress={like}
			      textStyle={styles.actionText} style={styles.likesContainer}>
				{quote.likes}
			</Chip>

			<Chip icon="thumb-down" onPress={dislike}
			      textStyle={styles.actionText} style={styles.dislikesContainer}>
				{quote.dislikes}
			</Chip>
		</View>
	);
}