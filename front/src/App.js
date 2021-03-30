import './App.css';
import ResultsDisplay from './results/ResultsDisplay';
import SearchBar from './searchbar/SearchBar';

function App() {
  return (
    <div className="App">
      <SearchBar />
      <ResultsDisplay />
    </div>
  );
}

export default App;
