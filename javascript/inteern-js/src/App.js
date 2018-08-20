import * as React from "react";
import PropTypes from "prop-types";
import Paper from "@material-ui";

export class Hello extends React.Component {
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
        <Counter increment={() => this.increment()} count={this.state.count} />
      </div>
    );
  }
}

const Counter = props => (
  <div>
    <Paper style={{margin: "1em", textAlign="center"}}>{props.count}</Paper>
    <button variant="contained" color="primary" onClick={props.increment}>
      increment
    </button>
  </div>
);
Counter.propTypes = { count: PropTypes.number, increment: PropTypes.func };
