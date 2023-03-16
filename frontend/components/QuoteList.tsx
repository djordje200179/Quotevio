import {Quote} from "../models";
import QuoteItem from "./QuoteItem";
import {FlatList, RefreshControl} from "react-native";

interface Props {
    quotes: Quote[];
    refreshing: boolean;
    onRefresh: () => void;
}

export default function QuoteList({quotes, refreshing, onRefresh}: Props) {
    const refreshControl = <RefreshControl refreshing={refreshing} onRefresh={onRefresh}/>;

    return (
        <FlatList data={quotes}
                  renderItem={({item}) => <QuoteItem quote={item}/>}
                  keyExtractor={item => item.id}
                  refreshControl={refreshControl}/>
    );
}