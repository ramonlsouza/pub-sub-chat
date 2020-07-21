import React from "react";

function TextInput(props) {
  return (
    <>
      <label>{props.label}</label>

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
