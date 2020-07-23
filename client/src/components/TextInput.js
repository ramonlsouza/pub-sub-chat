import React from "react";

import "./TextInput.css";

function TextInput(props) {
  return (
    <>
      <label className="text-input-label">{props.label}</label>

      <input
        className={props.classlist || "text-input"}
        type={props.type || "text"}
        placeholder={props.placeholder}
        onChange={props.handleChange}
        value={props.value}
      />
    </>
  );
}

export default TextInput;
