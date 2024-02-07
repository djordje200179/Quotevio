import {StyleSheet, SafeAreaView, Platform, StatusBar, ActivityIndicator, View} from "react-native";
import { MD3LightTheme as DefaultTheme, FAB, Searchbar } from "react-native-paper";
import {useEffect, useMemo, useState} from "react";
import { Quote } from "./models/quote";
import QuoteList from "./components/QuoteList";

const styles = StyleSheet.create({
	container: {
		flex: 1,
		paddingTop: 5 + (Platform.OS === "android" ? StatusBar.currentHeight! : 0),
		paddingHorizontal: 5
	},
	fab: {
		position: "absolute",
		right: 16,
		bottom: 16,
		backgroundColor: DefaultTheme.colors.primary
	}
});

export default function App() {
	const [refreshing, setRefreshing] = useState(true);
	const [searchQuery, setSearchQuery] = useState("");
	const [allQuotes, setAllQuotes] = useState<Quote[]>([]);

	function updateQuote(quote: Quote) {
		const newAllQuotes = allQuotes.map(q => q.id === quote.id ? quote : q);
		setAllQuotes(newAllQuotes);
	}

	function loadData() {
		const url = `${process.env.EXPO_PUBLIC_API_URL}/quotes/`;
		console.log(`Loading data from ${url}`);

		fetch(url)
			.then(response => response.json())
			.then((quotes:Quote[]) => {
				setRefreshing(false);
				setAllQuotes(quotes);
			})
			.catch(console.error);
	}

	useEffect(loadData,[]);

	const filteredQuotes = useMemo(
		() => allQuotes?.filter(quote => quote.text.includes(searchQuery)),
		[searchQuery, allQuotes]
	);

	return (
		<SafeAreaView style={styles.container}>
			<Searchbar placeholder="Search quotes"
			           value={searchQuery} onChangeText={setSearchQuery}/>

			{refreshing ? <ActivityIndicator/> : null}

			<QuoteList quotes={filteredQuotes}
					   refreshing={refreshing} onRefresh={loadData}
					   onQuoteUpdate={updateQuote}
			/>

			<FAB icon="plus" color={DefaultTheme.colors.inversePrimary}
			     size="medium" style={styles.fab}
			     onPress={() => console.log("Pressed")}
			/>
		</SafeAreaView>
	);
}