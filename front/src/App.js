import { useState } from 'react';
import './App.css';
import ResultsDisplay from './results/ResultsDisplay';
import SearchBar from './searchbar/SearchBar';

function App() {
  const [state, setState] = useState({
      account: "",
      min: 0,
      max: 0,
      year: 0,
      list: {}
    });

  return (
    <div className="App">
      <SearchBar  searchParams={state} setState={setState}/>
      <ResultsDisplay />
    </div>
  );
}

export default App;
