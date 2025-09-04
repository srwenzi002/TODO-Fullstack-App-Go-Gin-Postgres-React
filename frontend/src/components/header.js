import React from "react";
import "./styles/Header.css";

const system = { title: "ToDo List Insert1" };

class Header extends React.Component {
  render() {
    return <div className="Header"> { system.title } </div>;
  }
}

export default Header;
