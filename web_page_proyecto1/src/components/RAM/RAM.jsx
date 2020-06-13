import React from 'react';
import './RAM.css';
import { Line } from 'react-chartjs-2';
import { Container, Row, Col } from 'react-bootstrap';
import axios from 'axios';


export default class RAM extends React.Component {
  constructor(props) {
    super(props);

    let data_percentage = {
      labels: [],
      datasets: [
        {
          label: 'Total de RAM utilizada',
          fill: false,
          lineTension: 0.1,
          backgroundColor: 'rgba(75,192,192,0.4)',
          borderColor: 'rgba(75,192,192,1)',
          borderCapStyle: 'butt',
          borderDash: [],
          borderDashOffset: 0.0,
          borderJoinStyle: 'miter',
          pointBorderColor: 'rgba(75,192,192,1)',
          pointBackgroundColor: 'rgba(75,192,192,1)',
          pointBorderWidth: 1,
          pointHoverRadius: 5,
          pointHoverBackgroundColor: 'rgba(75,192,192,1)',
          pointHoverBorderColor: 'rgba(220,220,220,1)',
          pointHoverBorderWidth: 2,
          pointRadius: 5,
          pointHitRadius: 10,
          data: []
        }
      ]
    };
    //TOTAL DE RAM

    axios.get(`http://18.204.15.140:8080/total`)
      .then(res => {
        console.log(res.data)

        this.setState({ TotalRam: res.data });
      })
    this.state = { labels: [], data: data_percentage.datasets[0].data, actual: 0, dataset: data_percentage };
  }
  componentDidMount() {
    this.interval = setInterval(() => {
      axios.get(`http://18.204.15.140:8080/actualram`)
        .then(res => {
          let resp = res.data;

          let labels1 = this.state.labels;
          let dt = new Date();
          labels1.push(dt.toLocaleTimeString());
          let data1 = this.state.data;

          data1.push(res.data);
          if (labels1.length > 10) {
            labels1.shift();
            data1.shift();
          }

          this.setState({ data: data1 })
          //dataset
          let data_percentage = {
            labels: labels1,
            datasets: [
              {
                label: 'Total de RAM utilizada (MB)',
                fill: false,
                lineTension: 0.1,
                backgroundColor: 'rgba(75,192,192,0.4)',
                borderColor: 'rgba(75,192,192,1)',
                borderCapStyle: 'butt',
                borderDash: [],
                borderDashOffset: 0.0,
                borderJoinStyle: 'miter',
                pointBorderColor: 'rgba(75,192,192,1)',
                pointBackgroundColor: 'rgba(75,192,192,1)',
                pointBorderWidth: 1,
                pointHoverRadius: 5,
                pointHoverBackgroundColor: 'rgba(75,192,192,1)',
                pointHoverBorderColor: 'rgba(220,220,220,1)',
                pointHoverBorderWidth: 2,
                pointRadius: 5,
                pointHitRadius: 10,
                data: data1
              }
            ]
          };

          this.setState({ actual: resp, labels: labels1, data: data1, dataset: data_percentage });
          let lineChart = this.reference.chartInstance
          lineChart.update();

          axios.get(`http://18.204.15.140:8080/actualram`)
            .then(res => {
              console.log(res.data)

              this.setState({ RAMConsumida: res.data });
            })
          axios.get(`http://18.204.15.140:8080/ram`)
            .then(res => {
              console.log(res.data)

              this.setState({ percentage_ram: res.data });
            })
          axios.get(`http://18.204.15.140:8080/total`)
            .then(res => {
              console.log(res.data)

              this.setState({ TotalRam: res.data });
            })
        })
    }, 1000)
  }

  componentWillUnmount() {
    clearInterval(this.interval);
  }

  render() {
    return (
      <div>
        <Container fluid>
          <Row>
            <Col md={5}>
            </Col>
            <Col md={7}>
              <h1 > Monitor RAM</h1>
            </Col>
          </Row>
          <Row>
            <Col md={4}>
              <Row>
                <Col md={2}>
                </Col>
                <Col>
                  Cantidad Total de RAM
                </Col>

              </Row>
              <Row>
                <Col md={3}>
                </Col>
                <Col>
                  {this.state.TotalRam}MB
                </Col>
              </Row>
              <Row>
                <Col md={2}>
                </Col>
                <Col>
                  Cantidad consumida de RAM
                </Col>
              </Row>
              <Row>
                <Col md={3}>
                </Col>
                <Col>
                  {this.state.RAMConsumida}MB
                </Col>
              </Row>
              <Row>
                <Col md={2}>
                </Col>
                <Col>
                  % de Consumo de RAM
                </Col>
              </Row>
              <Row>
                <Col md={3}>
                </Col>
                <Col>
                  %{this.state.percentage_ram}
                </Col>
              </Row>
            </Col>
            <Col md={8}>
              <Line data={this.state.dataset} ref={(reference) => this.reference = reference} />
            </Col>
          </Row>

        </Container>

      </div>
    );
  }
}
