import React from 'react';
import './SuspendProcess.css';
import { Badge } from 'react-bootstrap';
import axios from 'axios';
export default class SuspendProcess extends React.Component {
  constructor(props) {
    super(props);
    this.state = { Procesos: 0 };
    axios.get(`http://18.204.15.140:8080/suspendprocess`)
      .then(res => {

        this.setState({ Procesos: res.data });

      })

  }
  componentDidMount() {
    this.interval = setInterval(() => {
      axios.get(`http://18.204.15.140:8080/suspendprocess`)
        .then(res => {

          this.setState({ Procesos: res.data });

        })
    }, 15000)
  }

  componentWillUnmount() {
    clearInterval(this.interval);
  }

  render() {
    return (
      <Badge variant="secondary">{this.state.Procesos}</Badge>

    );
  }
}

