import React from "react";

import "./Button.css";

function Button(props) {
  return (
    <>
      <button className="button" onClick={() => props.handleClick()}>
        {props.label}
      </button>
    </>
  );
}

export default Button;
