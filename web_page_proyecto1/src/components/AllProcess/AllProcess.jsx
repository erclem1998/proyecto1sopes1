import React from 'react';
import './AllProcess.css';
import { Table, Button } from 'react-bootstrap';
import axios from 'axios';
export default class AllProcess extends React.Component {
  constructor(props) {
    super(props);
    this.state = { Procesos: [] };
    axios.get(`http://18.204.15.140:8080/allprocess`)
      .then(res => {
        this.setState({ Procesos: res.data.Output });
      })
    //this.killprocess = this.killprocess.bind(this);
  }
  componentDidMount() {
    this.interval = setInterval(() => {
      axios.get(`http://18.204.15.140:8080/allprocess`)
        .then(res => {
          this.setState({ Procesos: res.data.Output });
        })
    }, 3000)
  }

  componentWillUnmount() {
    clearInterval(this.interval);
  }

  killprocess(e) {
    axios.post(`http://18.204.15.140:8080/killprocess`, e.target.value)
      .then(res => {
        console.log(res.data);
      })
  }
  render() {
    return (
      <Table bg="primary" hover striped>
        <thead bg="primary" >
          <tr>
            <th>PID</th>
            <th>Username</th>
            <th>Command</th>
            <th>State</th>
            <th>% RAM</th>
            <th> </th>
          </tr>
        </thead>
        <tbody>
          {
            this.state.Procesos.map((process, key) => {
              return (
                <tr key={process.PID}>
                  <td>{process.PID}</td>
                  <td>{process.Username}</td>
                  <td>{process.Command}</td>
                  <td>{process.State}</td>
                  <td>{process.Ram}</td>
                  <td><Button id={process.PID} value={process.PID} variant="danger" onClick={this.killprocess.bind(this)}>kill</Button></td>
                </tr>
              );
            })
          }
        </tbody>
      </Table>

    );
  }
}

