import React, { Component } from 'react';
import GraphiQL from 'graphiql'
import logo from './logo.svg';
import './App.css';
import '../node_modules/graphiql/graphiql.css'
import fetch from 'isomorphic-fetch'

const fetcher = (params) => {
  const baseUrl = 'http://localhost:8080/graphql'
  return fetch(baseUrl, {
      method: 'post',
      headers: { ContentType: 'application/json' },
      body: JSON.stringify(params)
  }).then(response => response.json())
}


class App extends Component {
  render() {
    return (
      <div className="App">
        <div style={{ height: '100vh', width: '100vw' }}>
          <GraphiQL fetcher={fetcher} />
        </div>
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h1 className="App-title">Welcome to React</h1>
        </header>
        
      </div>
    );
  }
}

export default App;
