import React, { useState } from "react";

import TextInput from "./TextInput";
import Button from "./Button";

import "./MainPage.css";

function MainPage(props) {
  const [message, setMessage] = useState("");
  const [messageList, setMessageList] = useState([]);

  function logout() {
    props.removeCookie("token");
  }
  function getMessages() {
    let headers = {};
    headers["Content-Type"] = "application/x-www-form-urlencoded";
    headers["Token"] = `${props.token}`;

    fetch(props.apiUrl + "get-messages", {
      method: "GET",
      headers,
    })
      .then((resp) => resp.json())
      .then(function (data) {
        if (data.Error === false && data.UserMessages != null) {
          setMessageList(data.UserMessages);
        } else {
          if (data.Error === true && data.Message != "") {
            alert(data.Message);
          }
        }
      })
      .catch(function () {
        alert("error!");
      });
  }
  function sendMessage() {
    let headers = {};
    headers["Content-Type"] = "application/x-www-form-urlencoded";
    headers["Token"] = `${props.token}`;

    fetch(props.apiUrl + "send-message", {
      method: "POST",
      headers,
      body: "message=" + message,
    })
      .then((resp) => resp.json())
      .then(function (data) {
        alert(data.Message);
        setMessage("");
      })
      .catch(function () {
        alert("error!");
      });
  }

  return (
    <div id="chat-wrapper">
      <div id="header">
        <Button
          label="Logout"
          classlist="button-small align-flex-end"
          handleClick={logout}
        />
        <h1>Fluency chat</h1>
      </div>
      <div id="main">
        {messageList.map((message) => (
          <p>
            <b>{message.SenderName}:</b> {message.MessageText}
          </p>
        ))}
      </div>
      <Button
        label="Get new messages"
        classlist="button get-messages-button"
        handleClick={getMessages}
      />
      <div id="new-message">
        <TextInput
          placeholder="type something here to send a message"
          classlist="new-message-text"
          value={message}
          handleChange={(e) => setMessage(e.target.value)}
        />
        <Button
          label="Send message"
          classlist="button new-message-button"
          handleClick={sendMessage}
        />
      </div>
    </div>
  );
}

export default MainPage;
