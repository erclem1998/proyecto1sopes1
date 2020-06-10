import React from 'react';
import './Homepage.css';
import { Container, Row, Col } from 'react-bootstrap';
import RunningProccess from '../RunningProccess/RunningProccess';
import SuspendProcess from '../SuspendProcess/SuspendProcess';
import StopProcess from '../StopProcess/StopProcess';
import ZombieProcess from '../ZombieProcess/ZombieProcess';
import TotalProcess from '../TotalProcess/TotalProcess';
export default class Homepage extends React.Component {
  render() {
    return (
      <div>
        <Container fluid>
          <Row>
            <Col md={4}>
              <RunningProccess />
            </Col>
            <Col md={4}>
              <SuspendProcess />
            </Col>
            <Col md={4}>
              <StopProcess />
            </Col>
          </Row>
        </Container>
      </div>
    );
  }
}

