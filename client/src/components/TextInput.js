import React from "react";

import "./TextInput.css";

function TextInput(props) {
  return (
    <>
      <label className="text-input-label">{props.label}</label>

      <input
        className="text-input"
        type={props.type || "text"}
        placeholder={props.placeholder}
        onChange={props.handleChange}
      />
    </>
  );
}

export default TextInput;
