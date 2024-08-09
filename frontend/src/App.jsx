import './App.css';
import Footer from './components/footer/footer';
import MainPage from './components/main/mainPage';
import Header from './components/header/header';

function App() {
    // const [resultText, setResultText] = useState("Please enter your name below 👇");
    // const [name, setName] = useState('');
    // const updateName = (e) => setName(e.target.value);
    // const updateResultText = (result) => setResultText(result);

    // function greet() {
    //     Greet(name).then(updateResultText);
    // }

    return (
        <div id="App">
            <Header />
            <MainPage />
            <Footer />
        </div>
    )
}

export default App
