package ghost

import (
  "os"
  "fmt"
  "io/ioutil"
  "encoding/json"
  "strings"
)

type Dbconfig struct {
  Encrypted string
  Visible string
}
var conf Dbconfig
func Init() {
  if _,
  err := os.Stat("./.DB"); os.IsNotExist(err) {
    err := os.Mkdir("./.DB", 0700)
    if err != nil {
      fmt.Println(err)
    }
  }
  content, errd := ioutil.ReadFile("dbconfig.json")
  if errd != nil {}
  json.Unmarshal([]byte(string(content)), &conf)
}
func Set(coll string, key string, val string) {
  chk := "./.DB/."+coll
  khk := chk+"/"+key+".[]"
  khkb := []byte(val)
  if _,error := os.Stat(chk); os.IsNotExist(error) {
    os.Mkdir(chk,0700)
    ioutil.WriteFile(khk,khkb,0) 
  }else{
     ioutil.WriteFile(khk,khkb,0) 
  }
}
func Get(coll string, key string) string{
  chk := "./.DB/."+coll
  khk := chk+"/"+key+".[]"
  out,err:=ioutil.ReadFile(khk)
  if err!= nil{
    fmt.Println(err)
  }
  return string(out)
}
func Del(coll string, key string){
  chk := "./.DB/."+coll
  khk := chk+"/"+key+".[]"
  os.Remove(khk)
}
func Delcol(coll string){
  chk := "./.DB/."+coll
  os.Remove(chk)
}
func Update(coll string,key string,val string) {
  chk := "./.DB/."+coll
  khk := chk+"/"+key+".[]"
  khkb := []byte(val)
  err:=ioutil.WriteFile(khk,khkb,0)
  if err!=nil{
    fmt.Println(err)
  }
}
func Getall(coll string) []string{
  chk := "./.DB/."+coll
  files, _ := ioutil.ReadDir(chk)
  var k []string
  for _, f := range files {
    if strings.HasSuffix(f.Name(), ".[]") {
        k=append(k,f.Name())
    }
  }
  i:=len(k)
  for j:=0;j<i;j++ {
      k[j]=strings.TrimRight(k[j], ".[]")
  }
  for m:=0;m<i;m++{
      k[m]=Get(coll,k[m])
  }
  return k
}