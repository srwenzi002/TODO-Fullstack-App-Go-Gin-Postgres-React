import React from "react";
import "./styles/Addbar.css";

class AddBar extends React.Component {
  addItem = event => {
    if (event.key === "Enter") {
      const value = event.target.value;
      if (value.trim() !== "") {
        this.props.addItem(value)
          .then(() => {
            this.setState({ value: "" });
          });
      }
    }
  };
  state = {
    value: "" // 受控输入框的值
  };
    // 输入框内容变化
  handleChange = (event) => {
    this.setState({ value: event.target.value });
  };

  render() {
    return (
      <div className="AddBar">
        <input
          className="AddBar-Text"
          type="text"
          value={this.state.value}
          onChange={this.handleChange}
          placeholder="Enter TODO Item"
          onKeyDown={this.addItem}
        />
      </div>
    );
  }
}

export default AddBar;
