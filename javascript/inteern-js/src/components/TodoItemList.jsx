import * as React from "react";
import PropTypes from "prop-types";

import Paper from "@material-ui/core/Paper";
import CircularProgress from "@material-ui/core/CircularProgress";

const paperStyle = { margin: "0.5em", padding: "5px" };

export const TodoItemList = props => {
  if (props.isLoading === true) {
    return <CircularProgress />;
  }
  const tasks = props.tasks.map((v, i) => (
    <Paper style={paperStyle} key={i}>
      {v.title}
    </Paper>
  ));
  return <div>{tasks}</div>;
};
TodoItemList.propTypes = { isLoading: PropTypes.bool, tasks: PropTypes.array };
