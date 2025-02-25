package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

/*
	book-controller will require unmarshalled data
	the request from user will be in JSON thus we unmarshall this data
	to use in controllers
*/

func ParseBody(r *http.Request, x interface{}){
	body, err:= io.ReadAll(r.Body)
	if err == nil{
		err = json.Unmarshal([]byte(body), x)
		if err != nil{
			return
		}
	}
}
