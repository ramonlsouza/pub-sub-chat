import React from "react";
import ReactDOM from "react-dom";
import { useCookies } from "react-cookie";

import LoginForm from "./components/LoginForm";
import MainPage from "./components/MainPage";

import "./index.css";

const App = () => {
  const [cookies, setCookie, removeCookie] = useCookies(["token"]);

  const apiUrl = "http://localhost:8000/";

  return (
    <>
      {cookies.token === undefined && (
        <LoginForm apiUrl={apiUrl} setCookie={setCookie} cookies={cookies} />
      )}
      {cookies.token !== undefined && (
        <MainPage apiUrl={apiUrl} removeCookie={removeCookie} />
      )}
    </>
  );
};

ReactDOM.render(<App />, document.getElementById("root"));
