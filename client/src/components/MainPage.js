import React, { useState } from "react";

import Message from "./Message";
import Header from "./main/Header";
import Actions from "./main/Actions";

import "./MainPage.css";

function MainPage(props) {
  const [messageList, setMessageList] = useState([]);

  return (
    <div id="chat-wrapper">
      <Header removeCookie={props.removeCookie} />

      <div id="main">
        {messageList.map((data, index) => (
          <Message key={index} data={data} userId={props.cookies.id} />
        ))}
      </div>

      <Actions
        setMessageList={setMessageList}
        apiUrl={props.apiUrl}
        token={props.token}
      />
    </div>
  );
}

export default MainPage;
