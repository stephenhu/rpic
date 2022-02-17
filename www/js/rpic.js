/* rpic.js */

const API                 = "/api";
const API_AUTH            = "/auth";
const API_NETWORKS        = "/networks";
const API_SERVICES        = "/services";
const API_SYSTEM          = "/systems";
const API_USERS           = "/users";

const ID_USER             = "user";
const ID_PASS             = "pass";

const HTTP_DELETE         = "delete";
const HTTP_GET            = "get";
const HTTP_PUT            = "put";

const HTTP_200            = 200;
const HTTP_401            = 401;

const ERR_LOGIN_FAILED    =
  "Login credentials are invalid";

const ATTR_ID             = "id";

const DBLOCK              = "d-block";
const DFLEX               = "d-flex";
const DNONE               = "d-none";

const VIEW_APP            = "app";
const VIEW_HOME           = "home";
const VIEW_LOGIN          = "login";
const VIEW_MAIN           = "main";
const VIEW_SIDEBAR        = "sidebar";

const VIEW_DASHBOARD      = "dashboard";
const VIEW_NETWORK        = "network";
const VIEW_STORAGE        = "storage";
const VIEW_SERVICES       = "services";
const VIEW_LOGS           = "logs";
const VIEW_WIREGUARD      = "wireguard";

const NULL_STRING         = "";


function parseIds(r) {

  var root = document.getElementById(r);

  if(root === undefined || root === null) {
    console.log("Not found");
  } else {

    var ids = new Array();

    for(var i = 0; i < root.children.length; i++) {
      ids.push(root.children[i].getAttribute(ATTR_ID));
    }

    return ids;
  
  }

} // parseIds


function getTagType(t) {

  if(t === DNONE || t === DBLOCK || t === DFLEX) {
    return t;
  } else {

    console.log("Invalid display type, defaulting to DBLOCK: " + t);
    return DBLOCK;

  }

} // getTagType


function showFromList(l, v, t) {

  let ids = parseIds(l);

  for(var i = 0; i < ids.length; i++) {

    if(ids[i] === v) {
      showView(ids[i], t);
    } else {
      hideView(ids[i], t);
    }

  }

} // showFromList


function hideList(v, t) {

  let ids = parseIds(VIEW_MAIN);

  for(var i = 0; i < ids.length; i++) {
    hideView(ids[i], t);
  }

} // hideList


function showView(v, t) {

  var e = document.getElementById(v);
  
  e.classList.remove(DNONE);

  let tag = getTagType(t);

  e.classList.add(tag);

} // showView


function hideView(v, t) {

  var e = document.getElementById(v);
  

  let tag = getTagType(t);


  e.classList.remove(tag);
  e.classList.add(DNONE);

} // hideView


function checkAuth() {

  window.addEventListener("hashchange", hashtagListener);

  fetch(`${API}${API_USERS}`, {
    method: HTTP_GET,
  })
  .then((response) => {

    if(response.status === HTTP_200) {

      hashtagListener();

    } else if(response.status === HTTP_401) {
      
      hideView(VIEW_APP, DFLEX);
      showView(VIEW_LOGIN, DBLOCK);

    }

  })
  .then((data) => {

  })
  .catch((error) => {
    console.log(error);
  })

} // checkAuth


function signin() {
  
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
      toggleView(VIEW_MAIN, VIEW_LOGIN);
    }

  })
  .then((data) => {
    console.log(data);
  })
  .catch((error) => {
    alert(ERR_LOGIN_FAILED);
  })

} // signin


function signout() {

  fetch(`${API_AUTH}`, {
    method: HTTP_DELETE,
  })
  .then((response) => {
    console.log(response);
    if(response.ok) {
      
      if(response.status === HTTP_200) {
        toggleView(VIEW_LOGIN, VIEW_MAIN);
      }

    }
  })
  .then((data) => {

  })
  .catch((error) => {
    console.log(error);
  })

} // signout


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


function hashtagListener() {

  var hash = new URL(document.URL).hash.substring(1);

  if(hash === NULL_STRING) {
    
    hideView(VIEW_LOGIN, DBLOCK);
    showView(VIEW_APP, DFLEX);
    
    showFromList(VIEW_MAIN, VIEW_DASHBOARD, DBLOCK);

  } else {

    showView(VIEW_APP, DFLEX);

    showFromList(VIEW_MAIN, hash, DBLOCK);

  }

} // hashtagListener


/*
function init(pages) {

  window.onload = function () {

    var main    = document.getElementById(VIEW_MAIN);
    var login   = document.getElementById(VIEW_LOGIN);

    var hash = new URL(document.URL).hash;

    if(hash === EMPTY) {

      hideView(VIEW_LOGIN, DBLOCK);
      
      showFromList(VIEW_APP, VIEW_DASHBOARD, DFLEX);

    } else {
      hashtagListener();
    }

  }

  window.addEventListener("hashchange", hashtagListener);

} // init
*/
