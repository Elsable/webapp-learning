import * as React from "react";
import * as ReactDOM from "react-dom";
import { Hello } from "./view";

const view = React.createElement(Hello);
ReactDOM.render(view, document.getElementById("root"));
