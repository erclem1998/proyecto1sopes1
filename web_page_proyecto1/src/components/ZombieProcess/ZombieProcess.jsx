import React from 'react';
import './ZombieProcess.css';
import { Table } from 'react-bootstrap';
export default class ZombieProcess extends React.Component {
  constructor(props) {
    super(props);

    this.state = { Procesos: [] };
  }

  render() {
    return (
      <div>
        <Table responsive>
          <thead>
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

