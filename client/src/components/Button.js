import React from "react";

import "./Button.css";

function Button(props) {
  return (
    <>
      <button
        className={props.classlist || "button"}
        onClick={() => props.handleClick()}
      >
        {props.label}
      </button>
    </>
  );
}

export default Button;
