import React from 'react';
import './Homepage.css';
import { Container, Row, Col } from 'react-bootstrap';
import RunningProccess from '../RunningProccess/RunningProccess';
import SuspendProcess from '../SuspendProcess/SuspendProcess';
import StopProcess from '../StopProcess/StopProcess';
import ZombieProcess from '../ZombieProcess/ZombieProcess';
import TotalProcess from '../TotalProcess/TotalProcess';
import AllProcess from '../AllProcess/AllProcess';
export default class Homepage extends React.Component {
  render() {
    return (
      <Container fluid>
        <Row>
          <Col md={5}>
          </Col>
          <Col>

            <h1>
              Procesos
            </h1>
          </Col>
        </Row>
        <Row>
          <Col md={4}>
          </Col>
          <Col>
            <h4>
              Total de Procesos
            </h4>
          </Col>
          <Col>
            <h4><TotalProcess /></h4>
          </Col>
        </Row>
        <Row>
          <Col md={4}>
          </Col>
          <Col>
            <h4>
              Procesos en Ejecucion
            </h4>
          </Col>
          <Col>
            <h4>
              <RunningProccess />
            </h4>
          </Col>
        </Row>
        <Row>
          <Col md={4}>
          </Col>
          <Col>
            <h4>
              Procesos Suspendidos
            </h4>
          </Col>
          <Col>
            <h4>
              <SuspendProcess />
            </h4>
          </Col>
        </Row>
        <Row>
          <Col md={4}>
          </Col>
          <Col>
            <h4>
              Procesos Detenidos
            </h4>
          </Col>
          <Col>
            <h4>
              <StopProcess />
            </h4>
          </Col>
        </Row>
        <Row>
          <Col md={4}>
          </Col>
          <Col>
            <h4>
              Procesos Zombie
            </h4>
          </Col>
          <Col>
            <h4>
              <ZombieProcess />
            </h4>
          </Col>
        </Row>
        <Row>
          <AllProcess />
        </Row>



      </Container>
    );
  }
}

