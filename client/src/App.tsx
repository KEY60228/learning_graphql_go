import React from 'react';
import { BrowserRouter } from 'react-router-dom';

import { User } from './User'
import { AuthorizedUser } from './AuthorizedUser'

const App: React.FC = () => {
  return (
    <BrowserRouter>
      <AuthorizedUser />
      <User />
    </BrowserRouter>
  )
}

export default App;
