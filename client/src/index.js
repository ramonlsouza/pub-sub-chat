import React, { useState } from "react";
import ReactDOM from "react-dom";

import TextInput from "./components/TextInput";
import Button from "./components/Button";

import "./index.css";

const App = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const api = "http://localhost:8000/";

  function onClick() {
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
        } else {
          alert(data.Message);
        }
      })
      .catch(function () {
        alert("error!");
      });
  }

  return (
    <div id="chat-wrapper">
      <div id="main">
        <h1>Fluency chat</h1>

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

        <Button label="Login" handleClick={onClick} />
      </div>
    </div>
  );
};

ReactDOM.render(<App />, document.getElementById("root"));
