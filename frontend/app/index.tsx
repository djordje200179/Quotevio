import {ActivityIndicator, FlatList, RefreshControl, SafeAreaView, StyleSheet} from "react-native";
import {FAB, MD3LightTheme as DefaultTheme, Searchbar} from "react-native-paper";
import {useMemo, useState} from "react";
import {useQuotes} from "../hooks/useQuotes";
import QuoteView from "../components/QuoteView";

const styles = StyleSheet.create({
	container: {
		flex: 1,
		paddingTop: 5,
		paddingHorizontal: 5
	},
	fab: {
		position: "absolute",
		right: 16,
		bottom: 16,
		backgroundColor: DefaultTheme.colors.primary
	},
	searchBar: {
		borderRadius: 5,
	}
});

export default function Index() {
	const [allQuotes, loading, updateQuote, refreshQuotes] = useQuotes();
	const [searchQuery, setSearchQuery] = useState("");

	const filteredQuotes = useMemo(
		() => allQuotes?.filter(quote => quote.text.includes(searchQuery)),
		[searchQuery, allQuotes]
	);

	return (
		<SafeAreaView style={styles.container}>
			<Searchbar placeholder="Search quotes" style={styles.searchBar}
			           value={searchQuery} onChangeText={setSearchQuery}/>

			{loading ? <ActivityIndicator/> : null}

			<FlatList data={filteredQuotes}
			          renderItem={({item}) => <QuoteView quote={item} onQuoteUpdate={updateQuote}/>}
			          keyExtractor={item => item.id}
			          refreshControl={<RefreshControl refreshing={loading} onRefresh={refreshQuotes}/>}
			          contentContainerStyle={{ flexGrow: 1 }}/>

			<FAB icon="plus" color={DefaultTheme.colors.inversePrimary}
			     size="medium" style={styles.fab}
			     onPress={() => console.log("Pressed")}
			/>
		</SafeAreaView>
	);
}