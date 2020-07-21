import React from "react";

function Button(props) {
  return (
    <>
      <button className="Button-btn" onClick={() => props.handleClick()}>
        {props.label}
      </button>

      <style jsx="true">
        {`
          .Button-btn {
            padding: 10px 0;
          }
        `}
      </style>
    </>
  );
}

export default Button;
