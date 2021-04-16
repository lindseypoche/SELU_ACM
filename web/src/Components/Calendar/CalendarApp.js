import 'whatwg-fetch';
import React from 'react';
import Scheduler from 'devextreme-react/scheduler';
import CustomStore from 'devextreme/data/custom_store';
import './Calendar.css';

function getData(_, requestOptions) {
    const PUBLIC_KEY = 'AIzaSyDoBXwZDgJdowyJC0rsJ_RaEuKP6LaDqPQ',
        CALENDAR_ID = 'c_9optv6njchcc4qck9sjn7vghm8@group.calendar.google.com';
    const dataUrl = ['https://www.googleapis.com/calendar/v3/calendars/',
        CALENDAR_ID, '/events?key=', PUBLIC_KEY].join('');

    return fetch(dataUrl, requestOptions).then(
        (response) => response.json()
    ).then((data) => data.items);
}

// function getEvents() {
//     let that = this;
//     function start() {
//         gapi.client.init({
//             'apiKey': 'AIzaSyDoBXwZDgJdowyJC0rsJ_RaEuKP6LaDqPQ'
//         }).then(function () {
//             return gapi.client.request({
//                 'path': 'https://www.googleapis.com/calendar/v3/calendars/${c_9optv6njchcc4qck9sjn7vghm8@group.calendar.google.com}/events',
//         })
//       }).then( (response) => {
//         let events = response.result.items
//         that.setState({
//           events
//         }, ()=>{
//           console.log(that.state.events);
//         })
//       }, function(reason) {
//         console.log(reason);
//       });
//     }
//     gapi.load('client', start)
//   }

const dataSource = new CustomStore({
    load: (options) => getData(options, { showDeleted: false })
});

const currentDate = new Date();
const views = ['day', 'workWeek', 'month'];

class CalendarApp extends React.Component {
    render() {
        return (
            <div className = "background">
            <div className="paraCont">
                <head>
                    <title>DevExtreme Demo</title>
                    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
                    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
                    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0" />
                    <link rel="stylesheet" type="text/css" href="https://cdn3.devexpress.com/jslib/20.2.5/css/dx.common.css" />
                    <link rel="stylesheet" type="text/css" href="https://cdn3.devexpress.com/jslib/20.2.5/css/dx.light.css" />

                </head>

                <body className="dx-viewport">
                    <div className="demo-container">
                        <div id="app">
                            <React.Fragment>
                                <div className="long-title">
                                    <h3>ACM Event Calendar</h3>
                                </div>
                                <Scheduler
                                    dataSource={dataSource}
                                    views={views}
                                    defaultCurrentView="month"
                                    defaultCurrentDate={currentDate}
                                    startDayHour={7}
                                    editing={false}
                                    showAllDayPanel={false}
                                    startDateExpr="start.dateTime"
                                    endDateExpr="end.dateTime"
                                    textExpr="summary"
                                    timeZone="America/Chicago" />
                            </React.Fragment>
                        </div>
                    </div>
                </body>
            </div>
            </div>

        );
    }
}

export default CalendarApp;

//https://js.devexpress.com/Demos/WidgetsGallery/Demo/Scheduler/GoogleCalendarIntegration/React/Light/