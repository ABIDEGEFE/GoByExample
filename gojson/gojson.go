package gojson

import (
	"fmt"
	"encoding/json/v2"
   "strings"
)

func GoJson(){
	fmt.Println("Go json is running..")
}
type response1 struct{
	Page int
	Fruits []string
}

type response2 struct{
	Page int `json:"Page"`
	Fruits []string `json:"Fruits"`
}
func JsonCaller(){
   bolB, _ := json.Marshal(true) // returns ASCII value of seach character
   fmt.Println(string(bolB))

   sliceF, _ := json.Marshal([]string{"fruits", "mango", "orage"})
   fmt.Println(string(sliceF))

   resp1 := &response1{
	Page: 1, 
	Fruits: []string{"Orage", "Pinapple", "Mango"},
   }
   enrespo1, _ := json.Marshal(resp1)
   fmt.Println("Encoding to json string", string(enrespo1))

   resp2 := &response2{
	Page: 2,
	Fruits: []string{"mango", "orange", "grapes"},
   }
   enrespo2, _ := json.Marshal(resp2)
   fmt.Println("Encoding with tag", string(enrespo2))

   byt := []byte(`{"num":5.5, "str":["a", "b", "c"]}`)
   var dat map[string]any 
   if err := json.Unmarshal(byt, &dat); err != nil{
	panic(err)
   }
   fmt.Println("data", dat["num"].(float64))

   letters := dat["str"].([]any)
   fmt.Println(letters[0])

   mydata := `{"Page":1, "Fruits":["Orange", "Avocado"]}`
   res := response1{}
   if err := json.Unmarshal([]byte(mydata), &res); err != nil{
      panic(err)
   }
   fmt.Println("New data", res.Fruits[0])

   // mydata := `{"Page":1, "Fruits":["Orange", "Avocado"]}`
   res1 := response2{}
    _ = json.UnmarshalRead(strings.NewReader(mydata), &res1)
    fmt.Println(res1)
}