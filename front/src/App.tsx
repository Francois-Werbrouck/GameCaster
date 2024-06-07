import { Component, lazy } from 'solid-js';


import { Route, Router } from '@solidjs/router'
import Nav from './components/Nav';
import path from 'path';
import { AuthProvider } from './context/AuthContext';


const MainPage = lazy(() => import("./routes/MainPage"));
const AuthPage = lazy(() => import("./routes/AuthPage"));
const PasswordResetPage = lazy(() => import("./routes/auth/PasswordReset"));

const App: Component = () => {
  return (
    <>
      <AuthProvider>
        <Nav>
          <Router >
            <Route path="/" component={MainPage} />
            <Route path="/auth" component={AuthPage} />
            <Route path="/auth/reset" component={PasswordResetPage} />
          </Router>
        </Nav>
      </AuthProvider>
    </>
  );
};

export default App;
