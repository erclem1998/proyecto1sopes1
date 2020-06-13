import React from 'react';
import './Homepage.css';
import { Container, Row, Col, Tab, Nav } from 'react-bootstrap';
import Statistics from '../Statistics/Statistics';
import AllProcess from '../AllProcess/AllProcess';
import Arbol from '../Arbol/Arbol';
export default class Homepage extends React.Component {
  constructor(props) {
    super(props);
  }
  render() {
    return (
      <Container fluid>
        <Statistics />
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

