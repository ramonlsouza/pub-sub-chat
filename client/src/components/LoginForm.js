import React, { useState } from "react";

import TextInput from "./TextInput";
import Button from "./Button";

import "./LoginForm.css";

function LoginForm(props) {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  function login() {
    fetch(props.apiUrl + "auth", {
      method: "POST",
      headers: {
        "content-type": "application/x-www-form-urlencoded",
      },
      body: "username=" + username + "&password=" + password,
    })
      .then((resp) => resp.json())
      .then(function (data) {
        if (data.Error === false) {
          props.setCookie("token", data.Token);
        } else {
          alert(data.Message);
        }
      })
      .catch(function () {
        alert("error!");
      });
  }
  return (
    <div id="login-wrapper">
      <div id="login-main">
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

        <Button label="Login" handleClick={login} />
      </div>
    </div>
  );
}

export default LoginForm;
