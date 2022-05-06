import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import { ApolloProvider, ApolloClient, HttpLink, InMemoryCache } from '@apollo/client'

const link = new HttpLink({
  uri: "http://localhost:8080/query",
  headers: {
    authorization: "e753ebecac81ade72470c507d0de84484dc84ca9"
  }
})
const cache = new InMemoryCache()
const client = new ApolloClient({link, cache})

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <React.StrictMode>
    <ApolloProvider client={client}>
      <App />
    </ApolloProvider>
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();