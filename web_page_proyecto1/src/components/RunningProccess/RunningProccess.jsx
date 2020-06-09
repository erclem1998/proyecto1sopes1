import React from 'react';
import './RunningProccess.css';
import { Table } from 'react-bootstrap';
export default class RunningProccess extends React.Component {
  constructor(props) {
    super(props);

    this.state = { Procesos: [] };
  }

  render() {
    return (
      <div>
        <Table responsive bg="primary" hover>
          <thead bg="primary" variant="dark">
            <tr>
              <th>PID</th>
              <th>Command</th>
              <th>Username</th>
              <th>State</th>
              <th>% RAM</th>

            </tr>
          </thead>
          <tbody>
            {
              this.state.Procesos.map((process) => {
                return (
                  <tr>
                    <td>{process.PID}</td>
                    <td>{process.Username}</td>
                    <td>{process.Command}</td>
                    <td>{process.State}</td>
                    <td>{process.Ram}</td>

                  </tr>
                );
              })
            }
          </tbody>
        </Table>
      </div>
    );
  }
}

