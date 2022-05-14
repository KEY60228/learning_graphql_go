import React from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';

import { User } from './User'

const App: React.FC = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<User />} />
      </Routes>
    </BrowserRouter>
  )
}

export default App;
