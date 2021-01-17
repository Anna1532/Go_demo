package Mylog

import (
	"fmt"
	"os"
	"path"
	"time"
)

//往文件中输出日志记录
type FileLogger struct{
	Level LogLevel
	FilePath string //日志文件保存路径
	FileName string //日志文件名
	FileObj *os.File//需要先打开一个文件
	errFileObj *os.File
	maxFileSize int64
}
//构造函数
func NewFileLog(levelStr,fp,fn string,maxSize int64)*FileLogger{
	logLevel,err:=parseLogLevel(levelStr)
	if err!=nil{
		panic(err)
	}
	fl:=&FileLogger{
		Level: logLevel,
		FilePath:fp,
		FileName:fn ,
		maxFileSize:maxSize,
	}
	err=fl.initFile()//打开文件并做初始化
	if err!=nil{
		panic(err)
	}
	return fl
}
//打开日志文件
func (f *FileLogger)initFile()(error){
	fullFileName:=path.Join(f.FilePath,f.FileName)//将文件名字和路径进行拼接
	fileObj,err:=os.OpenFile(fullFileName,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err!=nil{
		fmt.Println("Open log failed,err:%v",err)
		return err
	}
	errFileObj,err:=os.OpenFile(fullFileName+".err",os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err!=nil{
		fmt.Println("Open error log failed,err:%v",err)
		return err
	}
	f.FileObj=fileObj
	f.errFileObj=errFileObj
	return nil
}
//用来判断文件大小的方法,是否需要切割
func (f *FileLogger)CheckSize()bool{
	//获取文件对象详细信息
	fileInfo,err:=f.FileObj.Stat()
	if err!=nil{
		fmt.Printf("Get checksize fileInfo failed,err:%v\n",err)//报错
		return false
	}
	fmt.Printf("文件大小是:%dB\n",fileInfo.Size())
	//如果当前文件大小>=日志文件最大值就返回true
	return fileInfo.Size()>=f.maxFileSize
}
//需要切割日志文件切割方法
func (f *FileLogger)SplitFile()(*os.File,error){
	nowStr:=time.Now().Unix()
	fileInfo,err:=f.FileObj.Stat()
	if err!=nil{
		fmt.Printf("get splitfile file info failed,err:%v\n",err)
		return nil,err
	}
	//拼接一个日志文件备份的名字
	logName:=path.Join(f.FilePath,fileInfo.Name())
	newLogName:=fmt.Sprintf("%s/%s.bak%s",logName,nowStr)
	//1、关闭当前文件
		f.FileObj.Close()
		//2、备份rename
		//拿到当前的日志文件的完整路径
		os.Rename(logName,newLogName)
		//3、打开一个新的日志文件
		FileObj,err:=os.OpenFile(logName,os.O_CREATE|os.O_APPEND|os.O_WRONLY,0644)
		if err!=nil{
			fmt.Printf("open new log file failed,err:%v\n",err)
			return nil,err
		//4、将打开的日志文件对象赋值给f.FileObj
		f.FileObj=FileObj
		}
	return FileObj,nil
}
//记录日志的函数
func (f *FileLogger)log(lv LogLevel,format string,arg ...interface{}){
	if f.enable(lv) {
		msg := fmt.Sprintf(format, arg...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		if f.CheckSize() {
			newFile,err:=f.SplitFile()//日志文件
			if err!=nil{
				return
			}
			f.FileObj=newFile
		}
		fmt.Fprintf(f.FileObj,"[%s][%s][%s:%s:%d]%s\n", now.Format("2006-01-02 15：04：05"), getLogString(lv), funcName, fileName, lineNo, msg) //日期格式化msg
		if lv>=ERROR{
			//如果要记录的日志大于等于error级别，还需要在err日志文件中再记录一遍
			if f.CheckSize() {
				newFile,err:=f.SplitFile()//日志文件
				if err!=nil{
					return
				}
				f.errFileObj=newFile
			}
			fmt.Fprintf(f.errFileObj,"[%s][%s][%s:%s:%d]%s\n", now.Format("2006-01-02 15：04：05"), getLogString(lv), funcName, fileName, lineNo, msg)
		}
	}
}
//写一个方法来实现开关，来控制日志的输出级别
func (f FileLogger)enable(logLevel LogLevel)bool{
	return f.Level<=logLevel//如果输入的级别l比允许的logLevr小则不能输出？？？
}
//希望记录日志的时候不仅记录字符串还有值等类型，这里的format string,arg ...interface{}是参考fmt.Printf()函数，可以传入任何值
func (f FileLogger)Debug(format string,arg ...interface{}){//格式化文本，可变长的接口
		f.log(DEBUG,format,arg...)
}
func (f FileLogger)Trace(format string,arg ...interface{}){//格式化文本，可变长的接口
	f.log(TRACE,format,arg...)
}
func (f FileLogger)Info(format string,arg ...interface{}){
		f.log(INFO,format,arg...)
}
func (f FileLogger)Warning(format string,arg ...interface{}){
		f.log(WARNING,format,arg...)
}
func (f FileLogger)Error(format string,arg ...interface{}){
		f.log(ERROR,format,arg...)
}
func (f FileLogger)Fatal(format string,arg ...interface{}){
		f.log(FATAL,format,arg...)
}
//关闭文件
func (f *FileLogger)CloseFile(){
	f.FileObj.Close()
	f.errFileObj.Close()
}
