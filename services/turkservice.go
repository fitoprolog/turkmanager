package services 
import(
   "net/http"
   "github.com/fitoprolog/turkmanager/config"
   //"github.com/fitoprolog/turkmanager/controllers"
   "github.com/gorilla/mux"  
   "time"
   "os"
   "log"
)

func dummyCallback( res http.ResponseWriter , req *http.Request){
  req.ParseForm()
  msg:=req.Form.Get("message")
  if msg == ""{
    res.Write([]byte("provide a message\n"))
    return
  }
  fd, err := os.OpenFile("messages.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
  if err != nil {
    res.Write([]byte("could not open file\n")) 
  }

  defer fd.Close()

  if _, err = fd.WriteString(msg+"\n"); err != nil {
    panic(err)
  }
  res.Write([]byte("werks\n"))
}

func Serve(configPath string)(err error){
  
  conf,err := config.Parse(configPath)
  
  if err != nil {
    return 
  }
  router := mux.NewRouter()
  //endp :=controllers.NewEndPoint().SetCallback(dummyCallback)
  router.HandleFunc("/werks/{message}/",dummyCallback).Methods("GET","PUT")
  srv := &http.Server{
		Handler: router,
		Addr:    conf.Serverbase,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
  log.Printf("ready to rock")
  err=srv.ListenAndServe()
  return 
}

