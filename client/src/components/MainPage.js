import React from "react";

import Button from "./Button";

import "./MainPage.css";

function MainPage(props) {
  function logout() {
    props.removeCookie("token");
  }

  return (
    <div id="chat-wrapper">
      <div id="header">
        <Button
          label="Logout"
          type="button-small align-flex-end"
          handleClick={logout}
        />
        <h1>Fluency chat</h1>
      </div>
      <div id="main">
        <p>
          <b>user:</b> chat item
        </p>
        <p>
          <b>user:</b> chat item
        </p>
        <p>
          <b>user:</b> chat item
        </p>
        <p>
          <b>user:</b> chat item
        </p>
      </div>
      <div id="new-message">
        <input type="text" id="new-message-text" />
        <input type="button" id="new-message-button" value="Send message" />
      </div>
    </div>
  );
}

export default MainPage;
