import * as React from "react";
import PropTypes from "prop-types";

export function Hello() {
  return (
    <div>
      <Counter />
    </div>
  );
}

class Clock extends React.Component {
  constructor(props) {
    super(props);
    this.state = { date: new Date() };
  }
  componentDidMount() {
    this.timerID = setInterval(() => this.setState({ date: new Date() }), 1000);
  }
  componentWillUnMount() {
    clearInterval(this.timerID);
  }
  render() {
    return <StatelessClock date={new Date()} />;
  }
}

class Counter extends React.Component {
  constructor(props) {
    super(props);
    this.state = { count: 0 };
  }
  increment() {
    this.setState({ count: this.state.count + 1 });
  }
  render() {
    return (
      <div>
        <div>{this.state.count}</div>
        <button onClick={this.increment()}>increment</button>
      </div>
    );
  }
}

const StatelessClock = props => <div>{props.date.toLocaleTimeString()}</div>;
StatelessClock.propTypes = { date: PropTypes.object };

const StatelessCounter = props => <div>{props.count}</div>;
StatelessCounter.propTypes = { count: PropTypes.number };
