package controllers

import (
	"net/http"
	"github.com/gorilla/sessions"
)


type EndPoint struct {
  JSONBody bool
  Secure   bool 
  JSONResponse bool 
  Callback func (http.ResponseWriter ,*http.Request)
  CookieStore sessions.CookieStore
}

func NewEndPoint() * EndPoint{
  var ret EndPoint 
  return &ret 
}

func (endpoint *EndPoint ) UseJSONBody() *EndPoint{
  endpoint.JSONBody=true
  return endpoint
}

func (endpoint *EndPoint ) IsSecure() *EndPoint{
  endpoint.Secure=true
  return endpoint
}

func (endpoint *EndPoint ) UseJSONResponse() *EndPoint{
  endpoint.JSONResponse=true
  return endpoint
}
func (endpoint *EndPoint) SetCallback (Callback func (http.ResponseWriter ,*http.Request)) *EndPoint{
  endpoint.Callback =Callback;
  return endpoint
}

func (endpoint *EndPoint ) Call( res http.ResponseWriter , req *http.Request){
 if endpoint.Callback == nil{
   panic("unseted callback")
   return
 }
 endpoint.Callback(res, req)  
}

