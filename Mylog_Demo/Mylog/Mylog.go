package Mylog
//自定义一个日志库
import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

//造一个日志级别类型
type LogLevel uint16		//unit16表示无符号整型，这里定义了一个LogLevel的类型
const(
	UNKNOWN LogLevel=iota //给不同类型的日志记录分配不同的级别
	DEBUG//1
	TRACE//2
	INFO//3
	WARNING
	ERROR
	FATAL
)
//日志结构体
type Logger struct{
	Level LogLevel //日志结构体中只有一个表示级别的变量
}
//写一个方法将用户输入的字符类型转换成系统能够识别的LogLevel级别
//返回日志级别以及错误类型
func parseLogLevel(s string)(LogLevel,error){
	s= strings.ToLower(s)//把用户输入的字符串都转换成小写，是为了将用户不规范的输入转换成规范的输入
	switch s{
	case "debug":
		return DEBUG,nil
	case "trace":
		return TRACE,nil
	case "info":
		return INFO,nil
	case "warning":
		return WARNING,nil
	case "error":
		return ERROR,nil
	case "fatal":
		return FATAL,nil
	default:
		err:=errors.New("无效的日志级别")
		return UNKNOWN,err
	}
}
//将系统中的不同日志类型转换成字符，因为后面输出日志的时候需要标记类型
//返回一个字符串类型的数据
func getLogString(lv LogLevel)string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}
	return "BUG"
}
//输出日志时间，文件名，以及行数
func getInfo(n int)(funcName,fileName string,lineNo int){
	pc,file,lineNo,ok:=runtime.Caller(n)//如果可以取到，那么OK就是true；pc是调用函数相关的信息；file得到的是文件名；line是行数
	if !ok{
		fmt.Println("runtime.caller() failed!")
		return
	}
	funcName=runtime.FuncForPC(pc).Name()
	fileName=path.Base(file)
	funcName=strings.Split(funcName,".")[1]
	return
}