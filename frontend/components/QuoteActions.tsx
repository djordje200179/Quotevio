import { Quote } from "../models";
import { StyleSheet, View } from "react-native";
import { MD3LightTheme as DefaultTheme, Chip } from "react-native-paper";
import { useState } from "react";

interface Props {
	quote: Quote;
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

export default function QuoteActions({ quote }: Props) {
	const [liked, setLiked]       = useState(false);
	const [disliked, setDisliked] = useState(false);

	function like() {
		const params = new URLSearchParams({
			action: "like"
		});

		const url = `http://localhost:3000/quotes/${quote.id}?${params}`;

		fetch(url, { method: "PUT" })
			.then(() => setLiked(true))
			.catch(console.log);
	}

	function dislike() {
		const params = new URLSearchParams({
			action: "dislike"
		});

		const url = `http://localhost:3000/quotes/${quote.id}?${params}`;

		fetch(url, { method: "PUT" })
			.then(() => setDisliked(true))
			.catch(console.log);
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