require("bootstrap")
require("../css/main.scss")

var React = require('react')
var ReactDOM = require('react-dom')
var Header = require('./components/header.jsx')

function initApp() {
  var App = React.createClass({
    render: function() {
      return (
        <div>
          <Header/>
        </div>
      )
    }
  })

  ReactDOM.render(<App/>, document.getElementById('app'))

  var ws = new WebSocket("ws://localhost:3000/echo")

  ws.onopen = function(e) {
    console.log("ws opened")
  }

  ws.onmessage = function(e) {
    console.log(e)
  }

  ws.onclose = function(e) {
    console.log("ws closed")
  }

  window.ws = ws
}

document.addEventListener('DOMContentLoaded', initApp)
