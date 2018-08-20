import { connect } from "react-redux";

import { TodoItemList } from "./components/TodoItemList";

import { NewToDoItem } from "./components/NewToDoItem";

import { addTask } from "./actions";

export const ConnectedTodoItemList = connect(state => state)(TodoItemList);

export const ConnectedNewToDoItem = connect(
  state => state,
  dispatch => ({ addTask: title => dispatch(addTask(title)) })
)(NewToDoItem);
