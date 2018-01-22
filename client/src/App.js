import React, { Component } from 'react';
import GraphiQL from 'graphiql'
import './App.css';
import '../node_modules/graphiql/graphiql.css'
import fetch from 'isomorphic-fetch'
import * as firebase from 'firebase'

var config = {
  apiKey: "AIzaSyD7tG11nYwNHUPmJ6jj6p-NFJC7R8_Rkj8",
  authDomain: "learn-with-images.firebaseapp.com",
  databaseURL: "https://learn-with-images.firebaseio.com",
  projectId: "learn-with-images",
  storageBucket: "learn-with-images.appspot.com",
  messagingSenderId: "977496925083"
};
firebase.initializeApp(config);

const fetcher = (params) => {
  const getUrl = window.location
  const baseUrl = `${getUrl.protocol}//${getUrl.host}/graphql`
  return fetch(baseUrl, {
      method: 'post',
      headers: { ContentType: 'application/json' },
      body: JSON.stringify(params)
  }).then(response => response.json())
  .catch(err => {
    console.log('error', err)
  })
}


class App extends Component {

  handleLogin = () => {
    try {
      // Using a popup.
      var provider = new firebase.auth.GoogleAuthProvider();
      provider.addScope('profile');
      provider.addScope('email');
      firebase.auth().signInWithPopup(provider).then(function(result) {
        // This gives you a Google Access Token.
        var token = result.credential.accessToken;
        // The signed-in user info.
        var user = result.user;
        console.log('firebase', user, token)
      });
      // Start a sign in process for an unauthenticated user.
      var provider = new firebase.auth.GoogleAuthProvider();
      provider.addScope('profile');
      provider.addScope('email');
      firebase.auth().signInWithRedirect(provider);        
    } catch (err) {
      console.error(err)
    }
  }

  render() {
    return (
      <div className="App">
        <div style={{ height: '100vh', width: '100vw' }}>
          <GraphiQL fetcher={fetcher} />
        </div>
        <div style={{ padding: 20 }}>
          <button onClick={this.handleLogin}>Login</button>
        </div>

      </div>
    );
  }
}

export default App;
