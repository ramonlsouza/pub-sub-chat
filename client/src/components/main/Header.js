import React from "react";
import Button from "./../Button";

import "./Header.css";

function Header(props) {
  function logout() {
    props.removeCookie("token");
    props.removeCookie("id");
    props.removeCookie("username");
  }

  return (
    <div id="header">
      <Button
        label="Logout"
        classlist="button-small logout-button align-flex-end"
        handleClick={logout}
      />
      <h1>Fluency chat</h1>
    </div>
  );
}

export default Header;
