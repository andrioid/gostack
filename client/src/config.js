import firebase from 'firebase'

const config = {
    apiKey: "AIzaSyD7tG11nYwNHUPmJ6jj6p-NFJC7R8_Rkj8",
    authDomain: "learn-with-images.firebaseapp.com",
    databaseURL: "https://learn-with-images.firebaseio.com",
    projectId: "learn-with-images",
    storageBucket: "learn-with-images.appspot.com",
    messagingSenderId: "977496925083"
  }

firebase.initializeApp(config)

export const ref = firebase.database().ref()
export const firebaseAuth = firebase.auth