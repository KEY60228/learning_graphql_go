import React from 'react';
import ReactDOM from 'react-dom/client';
import reportWebVitals from './reportWebVitals';
import { ApolloProvider, ApolloClient, HttpLink, InMemoryCache } from '@apollo/client'
import { persistCache, LocalStorageWrapper } from 'apollo3-cache-persist';

import App from './App';

const link = new HttpLink({
  uri: "http://localhost:8080/query",
  headers: {
    authorization: localStorage.getItem("token")
  }
})
const cache = new InMemoryCache()
persistCache({
  cache,
  storage: new LocalStorageWrapper(localStorage),
})
const client = new ApolloClient({link, cache})

if (localStorage['apollo-cache-persist']) {
  const cacheData = JSON.parse(localStorage['apollo-cache-persist'])
  cache.restore(cacheData)
}

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
