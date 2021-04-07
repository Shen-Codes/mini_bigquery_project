import { useState } from 'react';
import './App.css';
import ResultsDisplay from './results/ResultsDisplay';
import SearchBar from './searchbar/SearchBar';

function App() {
  const [params, setParams] = useState({
      account: "",
      min: 0,
      max: 0,
      year: 0
    });

  const [list, setList] = useState();

    
  const fetchData = async () => {
    // const {account, min, max, year} = params;
    // const url = `https://something.com/account?=${account}&min=${min}&max=${max}&year=${year}`;
    // await fetch(url)
    //   .then(response => response.json)
    //   .then(response => {
    //     setList({list: response})
    //   })
    console.log(params)
  }

  return (
    <div className="App">
      <SearchBar  searchParams={params} setState={setParams} fetch={fetchData}/>
      <ResultsDisplay list={list} />
    </div>
  );
}

export default App;
