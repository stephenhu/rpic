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

const HTTP_401            = 401;

const ERR_LOGIN_FAILED    =
  "Login credentials are invalid";

const DBLOCK              = "d-block";
const DNONE               = "d-none";

const PAGE_HOME           = "home";
const PAGE_LOGIN          = "login";


function togglePage(s, h) {

  var show = document.getElementById(s);
  var hide = document.getElementById(h);

  show.classList.remove(DNONE);
  show.classList.add(DBLOCK);

  hide.classList.remove(DBLOCK);
  hide.classList.add(DNONE);  

} // togglePage


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
    
    console.log(response);
    if(!response.ok) {

      if(response.status === HTTP_401) {
        alert(ERR_LOGIN_FAILED);
      }
      
    } else {
      togglePage(PAGE_HOME, PAGE_LOGIN);
    }

  })
  .then((data) => {
    console.log(data);
  })
  .catch((error) => {
    alert(ERR_LOGIN_FAILED);
  })

} // login


function logout() {

  fetch(`${API_AUTH}`, {
    method: HTTP_DELETE,
  })
  .then((response) => {
    console.log(response);
    if(response.ok) {
      
    }
  })
  .then((data) => {

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
