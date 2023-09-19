package response

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Response struct {
	RespWriter gin.ResponseWriter
}

func (resp *Response) SetStatusOk() {
	resp.RespWriter.WriteHeader(http.StatusOK)
}

func (resp *Response) SetStatusBadRequest() {
	resp.RespWriter.WriteHeader(http.StatusBadRequest)
}

func (resp *Response) SetStatusNotFound() {
	resp.RespWriter.WriteHeader(http.StatusNotFound)
}

func (resp *Response) SetStatusInternalServerError() {
	resp.RespWriter.WriteHeader(http.StatusInternalServerError)
}

func (resp *Response) SetStatusConflict() {
	resp.RespWriter.WriteHeader(http.StatusConflict)
}

func (resp *Response) SetData(data []byte) {
	_, err := resp.RespWriter.Write(data)
	if err != nil {
		log.Println("No data has been sent!")
	}
}
