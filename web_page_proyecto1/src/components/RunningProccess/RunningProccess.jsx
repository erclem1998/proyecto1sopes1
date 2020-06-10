import React from 'react';
import './RunningProccess.css';
import { Badge } from 'react-bootstrap';
import axios from 'axios';
export default class RunningProccess extends React.Component {
  constructor(props) {
    super(props);
    this.state = { Procesos: 0 };
    axios.get(`http://54.144.197.130:8080/runningprocess`)
      .then(res => {
        this.setState({ Procesos: res.data });
      })

  }
  componentDidMount() {
    this.interval = setInterval(() => {
      axios.get(`http://54.144.197.130:8080/runningprocess`)
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

