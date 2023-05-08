package main

import (
	"./pb"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	var i int32 = 0
	test(i)
}
func test(t int32) {
	d := &pb.FrameD6{
		Datetimehour: "2020033014",
		Recordcnt:    t,
	}
	str, _ := json.Marshal(d)
	s := TransProtoToJson(d)
	fmt.Printf("@@@--incorrect JSON---> %+v \n", string(str))
	fmt.Printf("@@@--correct JSON---> %+v \n", s)
}

func TransProtoToJson(pb proto.Message) string {
	var pbMarshaler jsonpb.Marshaler
	pbMarshaler = jsonpb.Marshaler{
		EmitDefaults: true,
		OrigName:     true,
		EnumsAsInts:  true,
	}
	_buffer := new(bytes.Buffer)
	_ = pbMarshaler.Marshal(_buffer, pb)
	return string(_buffer.Bytes())
}
