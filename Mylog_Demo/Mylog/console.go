package Mylog
//往终端写日志相关内容
import (
	"fmt"
	"time"
)
//新建一条日志，输入级别（字符串类型）返回定义的级别常量
func Newlog(levelStr string) Logger {
	level,err:=parseLogLevel(levelStr)
	if err!=nil{
		panic(err)
	}
	return Logger{
		Level:level, //这里是为了实现级别，例如只有info之后的级别才能显示
	}
}
//写一个方法来实现开关，来控制日志的输出级别
func (l Logger)enable(logLevel LogLevel)bool{
	return l.Level<logLevel//如果输入的级别l比允许的logLevr小则不能输出？？？
}
//记录日志的函数
func (l Logger)log(lv LogLevel,format string,arg ...interface{}){//因为输出到文件和输出到终端的Print函数形式不同，所以需要将log函数转换成方法
	if l.enable(lv) {
		msg := fmt.Sprintf(format, arg...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(2)                                                                                   //为什么需要在这里调用
		fmt.Printf("[%s][%s][%s:%s:%d]%s\n", now.Format("2006-01-02 15：04：05"), getLogString(lv), funcName, fileName, lineNo, msg) //日期格式化msg
	}
}
//希望记录日志的时候不仅记录字符串还有值等类型，这里的format string,arg ...interface{}是参考fmt.Printf()函数，可以传入任何值
func (l Logger)Debug(format string,arg ...interface{}){//格式化文本，可变长的接口
		l.log(DEBUG,format,arg...)

}
func (l Logger)Trace(format string,arg ...interface{}){
	l.log(TRACE,format,arg...)
}
func (l Logger)Info(format string,arg ...interface{}){
		l.log(INFO,format,arg...)
}
func (l Logger)Warning(format string,arg ...interface{}){
		l.log(WARNING,format,arg...)
}
func (l Logger)Error(format string,arg ...interface{}){
		l.log(ERROR,format,arg...)
}
func (l Logger)Fatal(format string,arg ...interface{}){
		l.log(FATAL,format,arg...)
}