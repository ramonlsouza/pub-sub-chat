import React, { useState } from "react";
import Button from "./../Button";
import TextInput from "./../TextInput";

import "./Actions.css";

function Actions(props) {
  const [message, setMessage] = useState("");

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
          props.setMessageList(data.UserMessages);
          var element = document.getElementById("main");
          element.scrollTop = element.scrollHeight;
        } else {
          if (data.Error === true && data.Message !== "") {
            alert(data.Message);
          }
        }
      })
      .catch(function () {
        alert("Could not get messages!");
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
        setMessage("");
        getMessages();
      })
      .catch(function () {
        alert("Could not send message!");
      });
  }

  return (
    <>
      <Button
        label="Click here to check for new messages"
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
    </>
  );
}

export default Actions;
