import { StyleSheet, SafeAreaView, Platform, StatusBar, ActivityIndicator } from "react-native";
import { MD3LightTheme as DefaultTheme, FAB, Searchbar } from "react-native-paper";
import { useMemo, useState } from "react";
import { Quote } from "./models";
import QuoteList from "./components/QuoteList";
import useFetch from "react-fetch-hook";

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
	const [searchQuery, setSearchQuery] = useState("");

	const { isLoading: quotesLoading, data: allQuotes } = useFetch<Quote[]>("http://192.168.1.26:8080/quotes");

	const filteredQuotes = useMemo(
		() => allQuotes?.filter(quote => quote.text.includes(searchQuery)),
		[searchQuery, allQuotes]);

	return (
		<SafeAreaView style={styles.container}>
			<Searchbar placeholder="Search quotes"
			           value={searchQuery} onChangeText={setSearchQuery}/>

			{quotesLoading ? <ActivityIndicator/> : <QuoteList quotes={filteredQuotes}/>}

			<FAB icon="plus" color={DefaultTheme.colors.inversePrimary}
			     size="medium" style={styles.fab}
			     onPress={() => console.log("Pressed")}
			/>
		</SafeAreaView>
	);
}