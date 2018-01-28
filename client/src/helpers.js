import { ref, firebaseAuth } from "./config";

export function auth(email, pw) {
  return firebaseAuth()
    .createUserWithEmailAndPassword(email, pw)
    .then(saveUser);
}

export function loginWithGoogle() {
  var provider = new firebaseAuth.GoogleAuthProvider();
  provider.addScope("profile");
  provider.addScope("email");
  firebaseAuth().signInWithRedirect(provider);
}

export function logout() {
  return firebaseAuth().signOut();
}

export function login(email, pw) {
  return firebaseAuth().signInWithEmailAndPassword(email, pw);
}

export const loginAnonymous = () => {
  return firebaseAuth().signInAnonymously()
}

export function resetPassword(email) {
  return firebaseAuth().sendPasswordResetEmail(email);
}

export function saveUser(user) {
  return ref
    .child(`users/${user.uid}/info`)
    .set({
      email: user.email,
      uid: user.uid
    })
    .then(() => user);
}
