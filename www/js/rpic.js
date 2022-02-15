/* rpic.js */

const API                 = "/api";
const API_NETWORKS        = "/networks";
const API_SERVICES        = "/services";
const API_SYSTEM          = "/systems";
const API_AUTH            = "/auth";

const ID_USER             = "user";
const ID_PASS             = "pass";

const HTTP_DELETE         = "delete";
const HTTP_GET            = "get";
const HTTP_PUT            = "put";


function login() {
  
  var user = document.getElementById(ID_USER);
  var pass = document.getElementById(ID_PASS);

  var data = new FormData();

  data.append(ID_USER, user.value);
  data.append(ID_PASS, pass.value);

  fetch(`${API_AUTH}`, {
    method: HTTP_PUT,
    body: data
  })
  .then((response) => {
    if(response.ok) return response.json();
  })
  .then((data) => {
    console.log(data.length);

  })
  .catch((error) => {
    console.log(error);
  })

} // login


function logout() {

  fetch(`${API_AUTH}`, {
    method: HTTP_DELETE,
  })
  .then((response) => {
    if(response.ok) return response.json();
  })
  .then((data) => {
    console.log(data.length);

  })
  .catch((error) => {
    console.log(error);
  })

} // logout


function reboot() {

  fetch(`${API}${API_SYSTEMS}`, {
    method: HTTP_PUT,
  })
  .then((response) => {
    if(response.ok) return response.json();
  })
  .then((data) => {
    console.log(data.length);

  })
  .catch((error) => {
    console.log(error);
  })

} // reboot


function shutdown() {

  fetch(`${API}${API_SYSTEMS}`, {
    method: HTTP_PUT,
  })
  .then((response) => {
    if(response.ok) return response.json();
  })
  .then((data) => {
    console.log(data.length);

  })
  .catch((error) => {
    console.log(error);
  })

} // shutdown


function serviceCommand(c) {

  fetch(`${API}${API_SERVICES}`, {
    method: HTTP_PUT,
  })
  .then((response) => {
    if(response.ok) return response.json();
  })
  .then((data) => {
    console.log(data.length);

  })
  .catch((error) => {
    console.log(error);
  })

} // serviceCommand
