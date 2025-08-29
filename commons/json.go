// Need to parse json in post requests and convert to grpc code
package commons

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter,status int, data any){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)
	 json.NewEncoder(w).Encode(data) //serialize/marshal the data
}

func ReadJSON(r *http.Request,data any) error{
	return json.NewDecoder(r.Body).Decode(data)
}

func WriteError(w http.ResponseWriter,status int,message string){
	 writeJSON(w,status,map[string]string{"error":message}) //direct initialization and declaration
}