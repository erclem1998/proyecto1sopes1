import React from 'react';
import './Monitor.css';
import Header from '../Header/Header';
import {
    BrowserRouter as Router,
    Switch,
    Route
} from "react-router-dom";
import Homepage from '../Homepage/Homepage';
import CPU from '../CPU/CPU';
import RAM from '../RAM/RAM';
export default class Monitor extends React.Component {
    render() {
        return (
            <div>
                <Header />

                <Router>
                    <Switch>
                        <Route path="/CPU_Monitor">
                            <CPU />
                        </Route>
                        <Route path="/RAM_Monitor">
                            <RAM />
                        </Route>
                        <Route path="/">
                            <Homepage />
                        </Route>
                    </Switch>
                </Router>
            </div>
        );
    }
}

