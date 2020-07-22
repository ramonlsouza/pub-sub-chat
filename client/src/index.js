import React, { useState, useEffect } from "react";
import ReactDOM from "react-dom";
import { Cookies } from "react-cookie";

import TextInput from "./components/TextInput";
import Button from "./components/Button";

import "./index.css";

const App = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const cookies = new Cookies();
  const cookie_token = cookies.get("token");

  const [token, setToken] = useState(false);

  const api = "http://localhost:8000/";

  function logout() {
    cookies.remove("token");
    setToken(false);
  }

  function login() {
    fetch(api + "auth", {
      method: "POST",
      headers: {
        "content-type": "application/x-www-form-urlencoded",
      },
      body: "username=" + username + "&password=" + password,
    })
      .then((resp) => resp.json())
      .then(function (data) {
        if (data.Error === false) {
          alert("login ok!");
          cookies.set("token", data.Token);
          setToken(data.token);
        } else {
          alert(data.Message);
        }
      })
      .catch(function () {
        alert("error!");
      });
  }

  useEffect(() => {
    if (cookie_token !== undefined) {
      setToken(cookie_token);
    }
  }, [token, cookie_token]);

  return (
    <div id="chat-wrapper">
      <div id="main">
        <h1>Fluency chat</h1>
        {token === false && (
          <>
            <TextInput
              label="Username"
              placeholder="username"
              value={username}
              handleChange={(e) => setUsername(e.target.value)}
            />

            <TextInput
              label="Password"
              placeholder="password"
              type="password"
              value={password}
              handleChange={(e) => setPassword(e.target.value)}
            />

            <Button label="Login" handleClick={login} />
          </>
        )}
        {token !== false && (
          <>
            <p>Token: {token}</p>
            <Button label="Logout" handleClick={logout} />
          </>
        )}{" "}
      </div>
    </div>
  );
};

ReactDOM.render(<App />, document.getElementById("root"));
