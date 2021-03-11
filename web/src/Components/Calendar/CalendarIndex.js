import "./Calendar.css";

import React from "react";
import ReactDOM from "react-dom";

import CalendarApp from "./CalendarApp.js";

import themes from "devextreme/ui/themes";
themes.initialized(() =>
  ReactDOM.render(<CalendarApp />, document.getElementById("app"))
);
