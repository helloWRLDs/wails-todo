import { GetHostName } from '../../../wailsjs/go/main/App';
import './header.css'
import { useEffect, useState } from "react";

const Header = () => {
    const [greeting, setGreeting] = useState("")
    
    const getGreetings = () => {
        const time = new Date()
        if (time.getHours() < 12) {
            return "Good Morning"
        } else if (time.getHours() < 17) {
            return "Good Afternoon"
        } else {
            return "Good Evening"
        }
    }
    
    useEffect(() => {
        GetHostName()
            .then(resp => setGreeting(`${getGreetings()}, ${resp}`))
            .catch(e => console.error(e))
    }, [])

    return (
        <header>
            {greeting}
        </header>
    )
}

export default Header;