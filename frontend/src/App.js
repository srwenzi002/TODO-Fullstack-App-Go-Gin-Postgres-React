import React from "react";

import Header from "./components/header";
import AddBar from "./components/addbar";
import TodoList from "./components/todolist";

import "./App.css";

class App extends React.Component {
    constructor(props) {
    super(props);
    this.state = {
      items: []
    };
  }
    componentDidMount() {
    fetch("http://localhost:8081/items")
      .then(res => res.json())
      .then(json => this.setState({ items: json.data }));
  }
  addItem = (item) => {
    return fetch('http://localhost:8081/item/create', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ item: item }),
      mode: 'cors'
    })
    .then(res => res.json())
    .then(json => {
      this.setState({
        items: [...this.state.items, json.data]
      });
      
    });
  }
    removeItem = (id) => {
    let item = this.state.items.find(item => item.id === id);
    return fetch(`http://localhost:8081/item/delete`, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json'
      }, 
      body: JSON.stringify({ id: id, item: item.item, done: item.done })
    })
      .then(() => {
        this.setState({
          items: this.state.items.filter(item => item.id !== id)
        });
      });
  }
    toggleDone = (id) => {
    let items = [...this.state.items];
    let item = items.find(item => item.id === id);
    item.done = !item.done;

    return fetch(`http://localhost:8081/item/update`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ id: id, done: item.done }),
    })
    .then(() => {
      this.setState({ items: items });
    });
  }


  render() {
    return (
      <div className="App">
        <Header />
        <AddBar addItem={this.addItem} />
        <TodoList 
          items={this.state.items} 
          removeItem={this.removeItem} 
          toggleDone={this.toggleDone} 
        />
      </div>
    );
  }

}

export default App;
