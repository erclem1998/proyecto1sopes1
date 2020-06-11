import React from 'react';
import './Homepage.css';
import { Container, Row, Col, Tab, Nav } from 'react-bootstrap';
import RunningProccess from '../RunningProccess/RunningProccess';
import SuspendProcess from '../SuspendProcess/SuspendProcess';
import StopProcess from '../StopProcess/StopProcess';
import ZombieProcess from '../ZombieProcess/ZombieProcess';
import TotalProcess from '../TotalProcess/TotalProcess';
import AllProcess from '../AllProcess/AllProcess';
import Arbol from '../Arbol/Arbol';
export default class Homepage extends React.Component {
  constructor(props) {
    super(props);
  }
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
          <Col md={12}>
            <Tab.Container id="Selector-Tabs" defaultActiveKey="first" >
              <Row>
                <Col md={3}>
                  <Nav variant="pills" className="flex-column " >
                    <Nav.Item className="pb-2">
                      <Nav.Link eventKey="first" className="btn btn-outline-dark">Todos los Procesos</Nav.Link>
                    </Nav.Item>

                    <Nav.Item >
                      <Nav.Link eventKey="second" className="btn btn-outline-dark">Arbol de procesos</Nav.Link>
                    </Nav.Item>
                  </Nav>
                </Col>
                <Col md={9}>
                  <Tab.Content>
                    <Tab.Pane eventKey="first">
                      <AllProcess />
                    </Tab.Pane>
                    <Tab.Pane eventKey="second">
                      <Arbol />
                    </Tab.Pane>
                  </Tab.Content>
                </Col>
              </Row>
            </Tab.Container>
          </Col>

        </Row >



      </Container >
    );
  }
}

