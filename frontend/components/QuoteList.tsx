import {Quote} from "../models/quote";
import QuoteItem from "./QuoteItem";
import {FlatList, RefreshControl} from "react-native";

interface Props {
    quotes: Quote[];
    refreshing: boolean;
    onRefresh: () => void;
    onQuoteUpdate: (quote: Quote) => void;
}

export default function QuoteList({quotes, refreshing, onRefresh, onQuoteUpdate}: Props) {
    const refreshControl = <RefreshControl refreshing={refreshing} onRefresh={onRefresh}/>;

    return (
        <FlatList data={quotes}
                  renderItem={({item}) => <QuoteItem quote={item} onQuoteUpdate={onQuoteUpdate}/>}
                  keyExtractor={item => item.id}
                  refreshControl={refreshControl}
                  contentContainerStyle={{ flexGrow: 1 }}/>
    );
}