import {StyleSheet, SafeAreaView, Platform, StatusBar, ActivityIndicator, View} from "react-native";
import { MD3LightTheme as DefaultTheme, FAB, Searchbar } from "react-native-paper";
import {useEffect, useMemo, useState} from "react";
import { Quote } from "./models";
import QuoteList from "./components/QuoteList";

const styles = StyleSheet.create({
	container: {
		flex: 1,
		paddingTop: 5 + (Platform.OS === "android" ? StatusBar.currentHeight : 0),
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
		const newQuotes = allQuotes.map(currQuote => currQuote.id === quote.id ? quote : currQuote);
		setAllQuotes(newQuotes);
	}

	function loadData() {
		fetch("http://192.168.1.26:8080/quotes")
			.then((response) => response.json())
			.then(quotes => {
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

			<View>
				{refreshing ? <ActivityIndicator/> : null}

				<QuoteList quotes={filteredQuotes}
						   refreshing={refreshing} onRefresh={loadData}
						   onQuoteUpdate={updateQuote}
				/>
			</View>

			<FAB icon="plus" color={DefaultTheme.colors.inversePrimary}
			     size="medium" style={styles.fab}
			     onPress={() => console.log("Pressed")}
			/>
		</SafeAreaView>
	);
}