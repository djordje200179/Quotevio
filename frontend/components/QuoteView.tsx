import { Quote } from "../models/quote";
import { Card, Paragraph, Text } from "react-native-paper";
import { StyleSheet, View } from "react-native";
import QuoteActions from "./QuoteActions";

interface Props {
	quote: Quote;
	onQuoteUpdate: (quote: Quote) => void;
}

const styles = StyleSheet.create({
	card: {
		margin: 5
	},
	infoText: {
		fontSize: 12,
		fontStyle: "italic",
		textAlign: "right",
		textAlignVertical: "bottom"
	},
	bottom: {
		flexDirection: "row",
		justifyContent: "space-between",
		marginTop: 10
	},
});

export default function QuoteView({ quote, onQuoteUpdate }: Props) {
	return (
		<Card style={styles.card}>
			<Card.Content>
				<View>
					<Paragraph>{quote.text}</Paragraph>

					<View style={styles.bottom}>
						<QuoteActions quote={quote} onQuoteUpdate={onQuoteUpdate} />

						<View>
							<Text style={styles.infoText}>{quote.author}</Text>
							<Text style={styles.infoText}>22.08.2022.</Text>
						</View>

					</View>
				</View>
			</Card.Content>
		</Card>
	);
}