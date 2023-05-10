import React, {createContext} from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import UserStore from "./store/UserStore";

export const StoreContext = createContext(null)

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
      <StoreContext.Provider value={{
          user: new UserStore()
      }}>
          <App />
      </StoreContext.Provider>
  </React.StrictMode>
);

