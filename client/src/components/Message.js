import React from "react";

import "./Message.css";

function messageSide(userId, senderId) {
  if (parseInt(senderId) === parseInt(userId)) {
    return "message align-right";
  } else {
    return "message";
  }
}

function Message(props) {
  return (
    <span className={messageSide(props.userId, props.data.SenderId)}>
      <span className="message-inner">
        <b>{props.data.SenderName}: </b>
        {props.data.MessageText}{" "}
        <span className="message-date-time" title={props.data.Date}>
          ({props.data.Time})
        </span>
      </span>
    </span>
  );
}

export default Message;
