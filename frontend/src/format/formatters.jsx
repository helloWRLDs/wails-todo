import { GoDotFill } from "react-icons/go";

export const wrapPrior = (num) => {
    switch(num) {
        case 1: return <GoDotFill style={{color: "red"}}/>
        case 2: return <GoDotFill style={{color: "orange"}}/>
        case 3: return <GoDotFill style={{color: "yellow"}}/>
        default: return <GoDotFill style={{color: "green"}}/>
    }
}

export const formatDate = (dateString) => {
    const options = { day: 'numeric', month: 'short', hour: '2-digit', minute: '2-digit' };
    return new Date(dateString).toLocaleString('en-US', options).replace(',', '');
}