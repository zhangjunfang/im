package serializer

import (
	"fmt"
	"testing"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/zhangjunfang/im/base64Util"
	"github.com/zhangjunfang/im/protocol"
)

func Test_ser(t *testing.T) {
	mbean := protocol.NewTimMBean()
	body := "wuxiaodong"
	mbean.Body = &body
	b, _ := thrift.NewTSerializer().Write(mbean)
	base64str := string(base64Util.Base64Encode(b))
	fmt.Println(">>>>>>>>", base64str)
	var mbean2 *protocol.TimMBean = protocol.NewTimMBean()
	bb, _ := base64Util.Base64Decode([]byte(base64str))
	thrift.NewTDeserializer().Read(mbean2, bb)
	fmt.Println(mbean2)
	fmt.Println(*mbean2.Body)
}
