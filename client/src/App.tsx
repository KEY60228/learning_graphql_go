import React from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';

import { User } from './User'
import { PostPhoto } from './PostPhoto';

const App: React.FC = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<User />} />
        <Route path="/newPhoto" element={<PostPhoto />} />
      </Routes>
    </BrowserRouter>
  )
}

export default App;
