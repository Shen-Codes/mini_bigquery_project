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
    const {account, min, max, year} = params;

    const url = `https://srshif0waf.execute-api.us-east-1.amazonaws.com/testing/companyname?measure_tag=${account}&min=${min}&max=${max}&fiscal_year=${year}`;


    await fetch(url)
      .then(response => response.json())
      .then(data => setList(data))
      .catch(error => console.log(error))

  }

  return (
    <div className="App">
      {console.log(list)}
      <SearchBar  searchParams={params} setState={setParams} fetch={fetchData}/>
      <ResultsDisplay list={list} />
    </div>
  );
}

export default App;

//fiscal_year=2017&min=0&max=100000000&measure_tag=Revenues