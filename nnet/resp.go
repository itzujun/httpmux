package nnet

import (
	"net/http"
)

type ResponseWrap struct {
	Writer http.ResponseWriter
	Status int
	Committed bool
}

func newResponseWrap(w http.ResponseWriter)*ResponseWrap{
	return &ResponseWrap{Writer:w}
}


func(r*ResponseWrap)Header()http.Header{
	return r.Writer.Header()
}

func (r *ResponseWrap)Flush(){
	r.Writer.(http.Flusher).Flush()
}

func(r*ResponseWrap)WriteHeader(code int){

	if r.Committed{
		return
	}

	r.Status=code
	r.Writer.WriteHeader(code)
	r.Committed=true
}

func (r*ResponseWrap)Write(b[]byte){
	if !r.Committed{
		r.WriteHeader(http.StatusOK)
	}
}




