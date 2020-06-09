import React from 'react';
import './CPU.css';
import { Line } from 'react-chartjs-2';
import { Container, Row, Col } from 'react-bootstrap';
import axios from 'axios';


export default class CPU extends React.Component {
  constructor(props) {
    super(props);

    let data_percentage = {
      labels: [],
      datasets: [
        {
          label: '% de CPU utilizado',
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

    this.state = { labels: [], data: data_percentage.datasets[0].data, actual: 0, dataset: data_percentage };
  }
  componentDidMount() {
    this.interval = setInterval(() => {
      axios.get(`http://54.144.197.130:8080/cpu`)
        .then(res => {
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
                label: '% de CPU utilizado',
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

          this.setState({ actual: res.data, labels: labels1, data: data1, dataset: data_percentage });
          let lineChart = this.reference.chartInstance

          lineChart.update();


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
              <h1 > % de CPU utilizado</h1>
            </Col>
          </Row>
          <Row>
            <Col md={6}>
            </Col>
            <Col md={2}>
              <p > {this.state.actual}%</p>
            </Col>
          </Row>
          <Row>
            <Col md={2}>
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
