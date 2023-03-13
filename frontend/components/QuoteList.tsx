import { Quote } from "../models";
import QuoteItem from "./QuoteItem";
import { FlatList } from "react-native";

interface Props {
	quotes: Quote[];
}

export default function QuoteList({ quotes }: Props) {
	return (
		<FlatList data={quotes}
		          renderItem={({ item }) => <QuoteItem quote={item}/>}
		          keyExtractor={item => item.id}/>
	);
}