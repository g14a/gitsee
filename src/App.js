import React, { Component } from "react";
import "./App.css";
import UserSearchFrom from "./components/SearchField";
import ReactNotification from "react-notifications-component";
import PageComponent from "./components/PageComponent";
import { Route, BrowserRouter } from "react-router-dom";
import {Helmet} from "react-helmet";

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      username: "",
      error: null,
    };
  }

  handleInputChange = (event) => {
    event.preventDefault()
    this.setState({
      [event.target.name]: event.target.value
    })
  }

  render() {
    return (
      <div className="App">
        <Helmet>
          <meta charSet="utf-8" />
          <title>Gitsee</title>
        </Helmet>
        <BrowserRouter>
          <Route exact path="/gitsee">
            <ReactNotification />
            <UserSearchFrom handleInputChange={this.handleInputChange} username={this.state.username} />
            <Footer />
          </Route>
          <Route exact path="/gitsee/user/:username" component={(userProps) => {
            return <PageComponent username={userProps.match.params.username} />
          }} >
          </Route>
        </BrowserRouter>
      </div>
    );
  }
}

function Footer() {
  return (
    <footer className="footer" style={{ color: "white", marginTop: "30%" }}>
      <p>
        Built with GitHub V4 and React JS. Source is available on{" "}
        <a href="https://github.com/g14a/gitsee" style={{ color: "white" }}>
          Github
        </a>{" "}
      </p>
    </footer>
  );
}

export default App;
