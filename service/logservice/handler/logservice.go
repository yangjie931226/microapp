package handler

import (
	"context"
	logservice "logservice/proto"
	"os"
	stlog "log"

)


var log *stlog.Logger

type fileLog string

func (fl fileLog) Write(data []byte) (int, error) {
	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0766)
	if err != nil {
		return 0, nil
	}
	defer f.Close()
	return f.Write(data)
}

func Run(name string) {
	log = stlog.New(fileLog(name), "[go]", stlog.LstdFlags)
}

func write(data string) {
	log.Printf("%v\n", data)
}

type Logservice struct{}

func (*Logservice) WriteLog(ctx context.Context,req *logservice.WriteLogRequest,resp *logservice.LogReply) error {
	write(req.Message)
	return nil
}



