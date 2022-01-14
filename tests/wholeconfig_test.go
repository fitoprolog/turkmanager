package turkmanager 
import("testing"
  "github.com/fitoprolog/turkmanager/config"
)

func TestReadAConfig(t *testing.T){
  conf,err :=config.Parse("../config_json/test_config.json")
  if err != nil {
    t.Fatalf(err.Error())
  }
  if conf.Dbusername != "memeuser"{
    t.Fatalf("the field value is not the expected one (memeuser)")
    return
  }
  t.Logf("%+v\n",conf)
}
