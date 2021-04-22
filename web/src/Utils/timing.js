// utility functions for time

export const toDateFormat = (unix) => {
  var months = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sept", "Oct", "Nov", "Dec"]
  var d = new Date(unix * 1000);
  var month = months[d.getMonth()];
  var f = month + " " + d.getDay().toString()
  return f
}

export const isExpiring = (exp_date) => {
  var seconds_remaining = parseInt(exp_date) - Date.now()/ 1000
  if(seconds_remaining < 0) {
    return false
  }
  // if event is between 0 and 2 days 
  if(seconds_remaining < 86400*2) {
    return true
  }
  return false
}

export const getRemainingTime = (start_date) => {
  var format = ""
  var delta = parseInt(start_date) - Date.now()/ 1000
  console.log("delta: ", delta)

  if(delta < 0) {
    return "ended"
  }
  if(delta < 3600) {
    return "starts in < 1 hour"
  }

  // calculate and subtract days
  var days = Math.floor(delta / 86400); 
  delta -= days * 86400;
  if(days == 1) {
    format = "starts in " + days + " day"
  } else {
    format = "starts in " + days + " days"
  }

  // calculate and subtract hours 
  var hours = Math.floor(delta / 3600) % 24;
  delta -= hours * 3600;
  if(hours == 1) {
    format += " and " + hours + " hour"
  } else {
    format += " and " + hours + " hours"
  }

  return format
}