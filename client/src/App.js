import React, { Component } from "react";
import GraphiQL from "graphiql";
import "./App.css";
import "../node_modules/graphiql/graphiql.css";
import fetch from "isomorphic-fetch";
import { loginWithGoogle, login, logout } from "./helpers";
import { firebaseAuth } from './config'

class App extends Component {
  state = {
    loggedIn: false
  };

  state = {
    loggedIn: false,
    loading: true,
    firebaseIdToken: null,
    displayName: null
  };
  componentDidMount() {
    this.removeListener = firebaseAuth().onAuthStateChanged(async user => {
      console.log('user', user)
      if (user) {
        this.setState({
          loggedIn: true,
          loading: false,
          firebaseIdToken: await user.getIdToken(),
          displayName: user.displayName
        });
      } else {
        this.setState({
          loggedIn: false,
          loading: false
        });
      }
    });

    firebaseAuth().getRedirectResult().then(function(result) {
      if (result.credential) {
        // This gives you a Google Access Token. You can use it to access the Google API.
        var token = result.credential.accessToken;
        // ...
      }
      // The signed-in user info.
      var user = result.user;
      console.log('redirect result', user)
    }).catch(function(error) {
      console.log('error catching redirect', error)
      // Handle Errors here.
      var errorCode = error.code;
      var errorMessage = error.message;
      // The email of the user's account used.
      var email = error.email;
      // The firebase.auth.AuthCredential type that was used.
      var credential = error.credential;
      // ...
    });
  }

  componentWillUnmount() {
    this.removeListener();
  }

  handleLogin = () => {
    loginWithGoogle()
  };

  fetcher = params => {
    const getUrl = window.location;
    const baseUrl = `${getUrl.protocol}//${getUrl.hostname}:8080/graphql`;
    return fetch(baseUrl, {
      method: "post",
      headers: {
        ContentType: "application/json",
        Authorization: `Bearer ${this.state.firebaseIdToken}`
      },
      body: JSON.stringify(params)
    })
      .then(response => response.json())
      .catch(err => {
        console.log("error", err);
      });
  };

  render() {
    return (
      <div className="App">
        {this.state.loggedIn ? (
          <div style={{ height: "100vh", width: "100vw" }}>
            <p>{this.state.displayName}</p>
            <button onClick={() => logout()}>Logout</button>
            <GraphiQL fetcher={this.fetcher} />

          </div>
        ) : (
          <div style={{ padding: 20 }}>
            <button onClick={this.handleLogin}>Login</button>
          </div>
        )}
      </div>
    );
  }
}

export default App;
