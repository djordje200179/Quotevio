import {SafeAreaView, StyleSheet, Text} from "react-native";
import React, {useCallback} from "react";
import {Button, TextInput} from "react-native-paper";
import {MD3LightTheme as DefaultTheme} from "react-native-paper/lib/typescript/styles/themes/v3/LightTheme";
import {useRouter} from "expo-router";

const styles = StyleSheet.create({
	container: {
		flex: 1,
		paddingTop: 5,
		paddingHorizontal: 5
	},
	authorEntry: {
		borderRadius: 5,
	},
	quoteEntry: {
		borderRadius: 5,
	}
});

export default function NewQuote() {
	const [author, setAuthor] = React.useState("");
	const [quote, setQuote] = React.useState("");

	const router = useRouter();

	const submit = useCallback(async () => {
		const url = `${process.env.EXPO_PUBLIC_API_URL}/quotes/`;

		const body = {
			author,
			text: quote
		};

		await fetch(url, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(body)
		});

		router.back();
	}, [author, quote]);

	return (
		<SafeAreaView style={styles.container}>
			<TextInput label="Author" value={author}
			           onChangeText={text => setAuthor(text)} />

			<TextInput label="Quote" value={quote} multiline={true}
			           onChangeText={text => setQuote(text)} />

			<Button icon="upload" mode="contained"
			        disabled={!author || !quote}
			        onPress={submit}>
				Submit
			</Button>
		</SafeAreaView>
	)
}