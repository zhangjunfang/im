// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
	"net"
	"net/url"
	"os"
	"protocol"
	"strconv"
	"strings"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  void timStream(TimParam param)")
	fmt.Fprintln(os.Stderr, "  void timStarttls()")
	fmt.Fprintln(os.Stderr, "  void timLogin(Tid tid, string pwd)")
	fmt.Fprintln(os.Stderr, "  void timAck(TimAckBean ab)")
	fmt.Fprintln(os.Stderr, "  void timPresence(TimPBean pbean)")
	fmt.Fprintln(os.Stderr, "  void timMessage(TimMBean mbean)")
	fmt.Fprintln(os.Stderr, "  void timPing(string threadId)")
	fmt.Fprintln(os.Stderr, "  void timError(TimError e)")
	fmt.Fprintln(os.Stderr, "  void timLogout()")
	fmt.Fprintln(os.Stderr, "  void timRegist(Tid tid, string auth)")
	fmt.Fprintln(os.Stderr, "  void timRoser(TimRoster roster)")
	fmt.Fprintln(os.Stderr, "  void timMessageList(TimMBeanList mbeanList)")
	fmt.Fprintln(os.Stderr, "  void timPresenceList(TimPBeanList pbeanList)")
	fmt.Fprintln(os.Stderr, "  void timMessageIq(TimMessageIq timMsgIq, string iqType)")
	fmt.Fprintln(os.Stderr, "  void timMessageResult(TimMBean mbean)")
	fmt.Fprintln(os.Stderr, "  void timProperty(TimPropertyBean tpb)")
	fmt.Fprintln(os.Stderr, "  TimRemoteUserBean timRemoteUserAuth(Tid tid, string pwd, TimAuth auth)")
	fmt.Fprintln(os.Stderr, "  TimRemoteUserBean timRemoteUserGet(Tid tid, TimAuth auth)")
	fmt.Fprintln(os.Stderr, "  TimRemoteUserBean timRemoteUserEdit(Tid tid, TimUserBean ub, TimAuth auth)")
	fmt.Fprintln(os.Stderr, "  TimResponseBean timResponsePresence(TimPBean pbean, TimAuth auth)")
	fmt.Fprintln(os.Stderr, "  TimResponseBean timResponseMessage(TimMBean mbean, TimAuth auth)")
	fmt.Fprintln(os.Stderr, "  TimMBeanList timResponseMessageIq(TimMessageIq timMsgIq, string iqType, TimAuth auth)")
	fmt.Fprintln(os.Stderr, "  TimResponseBean timResponsePresenceList(TimPBeanList pbeanList, TimAuth auth)")
	fmt.Fprintln(os.Stderr, "  TimResponseBean timResponseMessageList(TimMBeanList mbeanList, TimAuth auth)")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := protocol.NewITimClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "timStream":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "TimStream requires 1 args")
			flag.Usage()
		}
		arg74 := flag.Arg(1)
		mbTrans75 := thrift.NewTMemoryBufferLen(len(arg74))
		defer mbTrans75.Close()
		_, err76 := mbTrans75.WriteString(arg74)
		if err76 != nil {
			Usage()
			return
		}
		factory77 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt78 := factory77.GetProtocol(mbTrans75)
		argvalue0 := protocol.NewTimParam()
		err79 := argvalue0.Read(jsProt78)
		if err79 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.TimStream(value0))
		fmt.Print("\n")
		break
	case "timStarttls":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "TimStarttls requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.TimStarttls())
		fmt.Print("\n")
		break
	case "timLogin":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "TimLogin requires 2 args")
			flag.Usage()
		}
		arg80 := flag.Arg(1)
		mbTrans81 := thrift.NewTMemoryBufferLen(len(arg80))
		defer mbTrans81.Close()
		_, err82 := mbTrans81.WriteString(arg80)
		if err82 != nil {
			Usage()
			return
		}
		factory83 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt84 := factory83.GetProtocol(mbTrans81)
		argvalue0 := protocol.NewTid()
		err85 := argvalue0.Read(jsProt84)
		if err85 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.TimLogin(value0, value1))
		fmt.Print("\n")
		break
	case "timAck":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "TimAck requires 1 args")
			flag.Usage()
		}
		arg87 := flag.Arg(1)
		mbTrans88 := thrift.NewTMemoryBufferLen(len(arg87))
		defer mbTrans88.Close()
		_, err89 := mbTrans88.WriteString(arg87)
		if err89 != nil {
			Usage()
			return
		}
		factory90 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt91 := factory90.GetProtocol(mbTrans88)
		argvalue0 := protocol.NewTimAckBean()
		err92 := argvalue0.Read(jsProt91)
		if err92 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.TimAck(value0))
		fmt.Print("\n")
		break
	case "timPresence":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "TimPresence requires 1 args")
			flag.Usage()
		}
		arg93 := flag.Arg(1)
		mbTrans94 := thrift.NewTMemoryBufferLen(len(arg93))
		defer mbTrans94.Close()
		_, err95 := mbTrans94.WriteString(arg93)
		if err95 != nil {
			Usage()
			return
		}
		factory96 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt97 := factory96.GetProtocol(mbTrans94)
		argvalue0 := protocol.NewTimPBean()
		err98 := argvalue0.Read(jsProt97)
		if err98 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.TimPresence(value0))
		fmt.Print("\n")
		break
	case "timMessage":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "TimMessage requires 1 args")
			flag.Usage()
		}
		arg99 := flag.Arg(1)
		mbTrans100 := thrift.NewTMemoryBufferLen(len(arg99))
		defer mbTrans100.Close()
		_, err101 := mbTrans100.WriteString(arg99)
		if err101 != nil {
			Usage()
			return
		}
		factory102 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt103 := factory102.GetProtocol(mbTrans100)
		argvalue0 := protocol.NewTimMBean()
		err104 := argvalue0.Read(jsProt103)
		if err104 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.TimMessage(value0))
		fmt.Print("\n")
		break
	case "timPing":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "TimPing requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.TimPing(value0))
		fmt.Print("\n")
		break
	case "timError":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "TimError requires 1 args")
			flag.Usage()
		}
		arg106 := flag.Arg(1)
		mbTrans107 := thrift.NewTMemoryBufferLen(len(arg106))
		defer mbTrans107.Close()
		_, err108 := mbTrans107.WriteString(arg106)
		if err108 != nil {
			Usage()
			return
		}
		factory109 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt110 := factory109.GetProtocol(mbTrans107)
		argvalue0 := protocol.NewTimError()
		err111 := argvalue0.Read(jsProt110)
		if err111 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.TimError(value0))
		fmt.Print("\n")
		break
	case "timLogout":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "TimLogout requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.TimLogout())
		fmt.Print("\n")
		break
	case "timRegist":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "TimRegist requires 2 args")
			flag.Usage()
		}
		arg112 := flag.Arg(1)
		mbTrans113 := thrift.NewTMemoryBufferLen(len(arg112))
		defer mbTrans113.Close()
		_, err114 := mbTrans113.WriteString(arg112)
		if err114 != nil {
			Usage()
			return
		}
		factory115 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt116 := factory115.GetProtocol(mbTrans113)
		argvalue0 := protocol.NewTid()
		err117 := argvalue0.Read(jsProt116)
		if err117 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.TimRegist(value0, value1))
		fmt.Print("\n")
		break
	case "timRoser":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "TimRoser requires 1 args")
			flag.Usage()
		}
		arg119 := flag.Arg(1)
		mbTrans120 := thrift.NewTMemoryBufferLen(len(arg119))
		defer mbTrans120.Close()
		_, err121 := mbTrans120.WriteString(arg119)
		if err121 != nil {
			Usage()
			return
		}
		factory122 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt123 := factory122.GetProtocol(mbTrans120)
		argvalue0 := protocol.NewTimRoster()
		err124 := argvalue0.Read(jsProt123)
		if err124 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.TimRoser(value0))
		fmt.Print("\n")
		break
	case "timMessageList":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "TimMessageList requires 1 args")
			flag.Usage()
		}
		arg125 := flag.Arg(1)
		mbTrans126 := thrift.NewTMemoryBufferLen(len(arg125))
		defer mbTrans126.Close()
		_, err127 := mbTrans126.WriteString(arg125)
		if err127 != nil {
			Usage()
			return
		}
		factory128 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt129 := factory128.GetProtocol(mbTrans126)
		argvalue0 := protocol.NewTimMBeanList()
		err130 := argvalue0.Read(jsProt129)
		if err130 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.TimMessageList(value0))
		fmt.Print("\n")
		break
	case "timPresenceList":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "TimPresenceList requires 1 args")
			flag.Usage()
		}
		arg131 := flag.Arg(1)
		mbTrans132 := thrift.NewTMemoryBufferLen(len(arg131))
		defer mbTrans132.Close()
		_, err133 := mbTrans132.WriteString(arg131)
		if err133 != nil {
			Usage()
			return
		}
		factory134 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt135 := factory134.GetProtocol(mbTrans132)
		argvalue0 := protocol.NewTimPBeanList()
		err136 := argvalue0.Read(jsProt135)
		if err136 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.TimPresenceList(value0))
		fmt.Print("\n")
		break
	case "timMessageIq":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "TimMessageIq requires 2 args")
			flag.Usage()
		}
		arg137 := flag.Arg(1)
		mbTrans138 := thrift.NewTMemoryBufferLen(len(arg137))
		defer mbTrans138.Close()
		_, err139 := mbTrans138.WriteString(arg137)
		if err139 != nil {
			Usage()
			return
		}
		factory140 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt141 := factory140.GetProtocol(mbTrans138)
		argvalue0 := protocol.NewTimMessageIq()
		err142 := argvalue0.Read(jsProt141)
		if err142 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.TimMessageIq(value0, value1))
		fmt.Print("\n")
		break
	case "timMessageResult":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "TimMessageResult_ requires 1 args")
			flag.Usage()
		}
		arg144 := flag.Arg(1)
		mbTrans145 := thrift.NewTMemoryBufferLen(len(arg144))
		defer mbTrans145.Close()
		_, err146 := mbTrans145.WriteString(arg144)
		if err146 != nil {
			Usage()
			return
		}
		factory147 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt148 := factory147.GetProtocol(mbTrans145)
		argvalue0 := protocol.NewTimMBean()
		err149 := argvalue0.Read(jsProt148)
		if err149 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.TimMessageResult_(value0))
		fmt.Print("\n")
		break
	case "timProperty":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "TimProperty requires 1 args")
			flag.Usage()
		}
		arg150 := flag.Arg(1)
		mbTrans151 := thrift.NewTMemoryBufferLen(len(arg150))
		defer mbTrans151.Close()
		_, err152 := mbTrans151.WriteString(arg150)
		if err152 != nil {
			Usage()
			return
		}
		factory153 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt154 := factory153.GetProtocol(mbTrans151)
		argvalue0 := protocol.NewTimPropertyBean()
		err155 := argvalue0.Read(jsProt154)
		if err155 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.TimProperty(value0))
		fmt.Print("\n")
		break
	case "timRemoteUserAuth":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "TimRemoteUserAuth requires 3 args")
			flag.Usage()
		}
		arg156 := flag.Arg(1)
		mbTrans157 := thrift.NewTMemoryBufferLen(len(arg156))
		defer mbTrans157.Close()
		_, err158 := mbTrans157.WriteString(arg156)
		if err158 != nil {
			Usage()
			return
		}
		factory159 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt160 := factory159.GetProtocol(mbTrans157)
		argvalue0 := protocol.NewTid()
		err161 := argvalue0.Read(jsProt160)
		if err161 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		arg163 := flag.Arg(3)
		mbTrans164 := thrift.NewTMemoryBufferLen(len(arg163))
		defer mbTrans164.Close()
		_, err165 := mbTrans164.WriteString(arg163)
		if err165 != nil {
			Usage()
			return
		}
		factory166 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt167 := factory166.GetProtocol(mbTrans164)
		argvalue2 := protocol.NewTimAuth()
		err168 := argvalue2.Read(jsProt167)
		if err168 != nil {
			Usage()
			return
		}
		value2 := argvalue2
		fmt.Print(client.TimRemoteUserAuth(value0, value1, value2))
		fmt.Print("\n")
		break
	case "timRemoteUserGet":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "TimRemoteUserGet requires 2 args")
			flag.Usage()
		}
		arg169 := flag.Arg(1)
		mbTrans170 := thrift.NewTMemoryBufferLen(len(arg169))
		defer mbTrans170.Close()
		_, err171 := mbTrans170.WriteString(arg169)
		if err171 != nil {
			Usage()
			return
		}
		factory172 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt173 := factory172.GetProtocol(mbTrans170)
		argvalue0 := protocol.NewTid()
		err174 := argvalue0.Read(jsProt173)
		if err174 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		arg175 := flag.Arg(2)
		mbTrans176 := thrift.NewTMemoryBufferLen(len(arg175))
		defer mbTrans176.Close()
		_, err177 := mbTrans176.WriteString(arg175)
		if err177 != nil {
			Usage()
			return
		}
		factory178 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt179 := factory178.GetProtocol(mbTrans176)
		argvalue1 := protocol.NewTimAuth()
		err180 := argvalue1.Read(jsProt179)
		if err180 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.TimRemoteUserGet(value0, value1))
		fmt.Print("\n")
		break
	case "timRemoteUserEdit":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "TimRemoteUserEdit requires 3 args")
			flag.Usage()
		}
		arg181 := flag.Arg(1)
		mbTrans182 := thrift.NewTMemoryBufferLen(len(arg181))
		defer mbTrans182.Close()
		_, err183 := mbTrans182.WriteString(arg181)
		if err183 != nil {
			Usage()
			return
		}
		factory184 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt185 := factory184.GetProtocol(mbTrans182)
		argvalue0 := protocol.NewTid()
		err186 := argvalue0.Read(jsProt185)
		if err186 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		arg187 := flag.Arg(2)
		mbTrans188 := thrift.NewTMemoryBufferLen(len(arg187))
		defer mbTrans188.Close()
		_, err189 := mbTrans188.WriteString(arg187)
		if err189 != nil {
			Usage()
			return
		}
		factory190 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt191 := factory190.GetProtocol(mbTrans188)
		argvalue1 := protocol.NewTimUserBean()
		err192 := argvalue1.Read(jsProt191)
		if err192 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		arg193 := flag.Arg(3)
		mbTrans194 := thrift.NewTMemoryBufferLen(len(arg193))
		defer mbTrans194.Close()
		_, err195 := mbTrans194.WriteString(arg193)
		if err195 != nil {
			Usage()
			return
		}
		factory196 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt197 := factory196.GetProtocol(mbTrans194)
		argvalue2 := protocol.NewTimAuth()
		err198 := argvalue2.Read(jsProt197)
		if err198 != nil {
			Usage()
			return
		}
		value2 := argvalue2
		fmt.Print(client.TimRemoteUserEdit(value0, value1, value2))
		fmt.Print("\n")
		break
	case "timResponsePresence":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "TimResponsePresence requires 2 args")
			flag.Usage()
		}
		arg199 := flag.Arg(1)
		mbTrans200 := thrift.NewTMemoryBufferLen(len(arg199))
		defer mbTrans200.Close()
		_, err201 := mbTrans200.WriteString(arg199)
		if err201 != nil {
			Usage()
			return
		}
		factory202 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt203 := factory202.GetProtocol(mbTrans200)
		argvalue0 := protocol.NewTimPBean()
		err204 := argvalue0.Read(jsProt203)
		if err204 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		arg205 := flag.Arg(2)
		mbTrans206 := thrift.NewTMemoryBufferLen(len(arg205))
		defer mbTrans206.Close()
		_, err207 := mbTrans206.WriteString(arg205)
		if err207 != nil {
			Usage()
			return
		}
		factory208 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt209 := factory208.GetProtocol(mbTrans206)
		argvalue1 := protocol.NewTimAuth()
		err210 := argvalue1.Read(jsProt209)
		if err210 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.TimResponsePresence(value0, value1))
		fmt.Print("\n")
		break
	case "timResponseMessage":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "TimResponseMessage requires 2 args")
			flag.Usage()
		}
		arg211 := flag.Arg(1)
		mbTrans212 := thrift.NewTMemoryBufferLen(len(arg211))
		defer mbTrans212.Close()
		_, err213 := mbTrans212.WriteString(arg211)
		if err213 != nil {
			Usage()
			return
		}
		factory214 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt215 := factory214.GetProtocol(mbTrans212)
		argvalue0 := protocol.NewTimMBean()
		err216 := argvalue0.Read(jsProt215)
		if err216 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		arg217 := flag.Arg(2)
		mbTrans218 := thrift.NewTMemoryBufferLen(len(arg217))
		defer mbTrans218.Close()
		_, err219 := mbTrans218.WriteString(arg217)
		if err219 != nil {
			Usage()
			return
		}
		factory220 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt221 := factory220.GetProtocol(mbTrans218)
		argvalue1 := protocol.NewTimAuth()
		err222 := argvalue1.Read(jsProt221)
		if err222 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.TimResponseMessage(value0, value1))
		fmt.Print("\n")
		break
	case "timResponseMessageIq":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "TimResponseMessageIq requires 3 args")
			flag.Usage()
		}
		arg223 := flag.Arg(1)
		mbTrans224 := thrift.NewTMemoryBufferLen(len(arg223))
		defer mbTrans224.Close()
		_, err225 := mbTrans224.WriteString(arg223)
		if err225 != nil {
			Usage()
			return
		}
		factory226 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt227 := factory226.GetProtocol(mbTrans224)
		argvalue0 := protocol.NewTimMessageIq()
		err228 := argvalue0.Read(jsProt227)
		if err228 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		arg230 := flag.Arg(3)
		mbTrans231 := thrift.NewTMemoryBufferLen(len(arg230))
		defer mbTrans231.Close()
		_, err232 := mbTrans231.WriteString(arg230)
		if err232 != nil {
			Usage()
			return
		}
		factory233 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt234 := factory233.GetProtocol(mbTrans231)
		argvalue2 := protocol.NewTimAuth()
		err235 := argvalue2.Read(jsProt234)
		if err235 != nil {
			Usage()
			return
		}
		value2 := argvalue2
		fmt.Print(client.TimResponseMessageIq(value0, value1, value2))
		fmt.Print("\n")
		break
	case "timResponsePresenceList":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "TimResponsePresenceList requires 2 args")
			flag.Usage()
		}
		arg236 := flag.Arg(1)
		mbTrans237 := thrift.NewTMemoryBufferLen(len(arg236))
		defer mbTrans237.Close()
		_, err238 := mbTrans237.WriteString(arg236)
		if err238 != nil {
			Usage()
			return
		}
		factory239 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt240 := factory239.GetProtocol(mbTrans237)
		argvalue0 := protocol.NewTimPBeanList()
		err241 := argvalue0.Read(jsProt240)
		if err241 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		arg242 := flag.Arg(2)
		mbTrans243 := thrift.NewTMemoryBufferLen(len(arg242))
		defer mbTrans243.Close()
		_, err244 := mbTrans243.WriteString(arg242)
		if err244 != nil {
			Usage()
			return
		}
		factory245 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt246 := factory245.GetProtocol(mbTrans243)
		argvalue1 := protocol.NewTimAuth()
		err247 := argvalue1.Read(jsProt246)
		if err247 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.TimResponsePresenceList(value0, value1))
		fmt.Print("\n")
		break
	case "timResponseMessageList":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "TimResponseMessageList requires 2 args")
			flag.Usage()
		}
		arg248 := flag.Arg(1)
		mbTrans249 := thrift.NewTMemoryBufferLen(len(arg248))
		defer mbTrans249.Close()
		_, err250 := mbTrans249.WriteString(arg248)
		if err250 != nil {
			Usage()
			return
		}
		factory251 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt252 := factory251.GetProtocol(mbTrans249)
		argvalue0 := protocol.NewTimMBeanList()
		err253 := argvalue0.Read(jsProt252)
		if err253 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		arg254 := flag.Arg(2)
		mbTrans255 := thrift.NewTMemoryBufferLen(len(arg254))
		defer mbTrans255.Close()
		_, err256 := mbTrans255.WriteString(arg254)
		if err256 != nil {
			Usage()
			return
		}
		factory257 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt258 := factory257.GetProtocol(mbTrans255)
		argvalue1 := protocol.NewTimAuth()
		err259 := argvalue1.Read(jsProt258)
		if err259 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.TimResponseMessageList(value0, value1))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
