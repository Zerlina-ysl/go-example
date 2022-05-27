package main

import (
	"bufio"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)
const socks5Ver = 0x05
const cmdBind = 0x01
const atypeIPV4 = 0x01
const atypeHOST = 0x03
const atypeIPV6 = 0x04
//tcp echo server
//返回输入信息的tcp server
func main(){
	//Listen监听一个网络端口上的连接
	server,err:=net.Listen("tcp","127.0.0.1:1080")
	if err!=nil{
		panic(err)
	}
	for{
		//Accept()会阻塞 直到新的连接被创建 并返回一个net.Conn来表示该连接
		client,err:=server.Accept()
		if err!=nil{
			log.Printf("Accept Failed %v",err)
			continue
		}
		//加入go关键字 每一次handleConn的调用都进入一个独立的goroutine 使其支持并发
		go process2(client)
	}
}
func process2(conn net.Conn){
	//defer语句用于延迟一个函数或方法的执行，会在外围函数或者方法返回之前但是返回值计算之后执行，可以在一个延迟执行的函数内部修改函数的命名返回值
	//如果一个函数有多个defer语句，以LIFO的顺序执行
	//该模式创建了一个值，在垃圾收集之前延迟执行一些关闭函数来清理该值
	defer conn.Close()
	//创建缓冲流
	reader:=bufio.NewReader(conn)
	err:=auth(reader,conn)
	if err!=nil{
		log.Printf("client %v auth failed:%v",conn.RemoteAddr(),err)
		return
	}
	log.Println("auth success")
	err=connect(reader,conn)
	if err!=nil{
		log.Printf("client %v auth failed:%v",conn.RemoteAddr(),err)
		return
	}

}
/**
实现认证功能
*/
func auth(reader *bufio.Reader,conn net.Conn)(err error){
	// VER NMETHODS METHODS
	// 1       1     1~255
	//VER:协议版本，socks5为0x05
	//NMTHODS 支持认证的方法数量
	//METHODS  对于NMTHODS,NMTHODS的值为多少，METHODS就有多少个字节
	// X'00'NO AUTHENTICATION REQUIRED
	// X'02' USERNAME/PASSWORD
	//鉴传
	ver,err:= reader.ReadByte()
	if err!=nil{
		//扩展了 fmt.Errorf 函数，加了一个 %w 来生成一个可以Wrapping Error,
		return fmt.Errorf("read ver failed:%w",err)
	}
	if ver!=socks5Ver{
		return fmt.Errorf("not supported ver:%v",ver)
	}
	methodSize ,err:= reader.ReadByte()
	if err!=nil{
		return fmt.Errorf("read methodSize failed:%v",err)
	}
	//使用methodsSize创建一个缓冲区slice
	method:=make([]byte,methodSize)
	_,err=io.ReadFull(reader,method)
	if err!=nil{
		return fmt.Errorf("read method failed:%v",err)
	}
	log.Println("ver",ver,"method",method)
	_,err=conn.Write([]byte{socks5Ver,0x00})
	if err!=nil{
		return fmt.Errorf("write failed:%v",err)
	}
	return nil
}
/**
请求阶段
 */
func connect(reader *bufio.Reader,conn net.Conn)(err error){
	// VER   CMD   RSV    ATYP     DST.ADDR    DST.PORT
	//  1     1    X'00    1         Variable    2
	//VER 版本号 socks5 0x05
	//CMD 请求的类型 0x01表示connect请求
	//RSV 保留字段 0x00
	//ATYP   目标地址类型 DST.ADDR的数据对应这个字段的类型
	// 0x01表示IPV4 DST.ADDR是4个字节
	// 0x03表示域名  DST.ADDR是一个可变长的域名
	//DST.ADDR
	//DST.PORT 目标端口 默认两个字节

	//向buf中读四个字节
	buf:=make([]byte,4)
	//浏览器向代理服务器发送报文
	_,err=io.ReadFull(reader,buf)
	if err!=nil{
		return fmt.Errorf("read header failed:%w",err)
	}
	ver,cmd,atyp:=buf[0],buf[1],buf[3]
	if ver!=socks5Ver{
		return fmt.Errorf("not supported ver:%v",ver)

	}
	if cmd!=cmdBind{
		return fmt.Errorf("not supported cmd:%v",cmd)
	}
	addr:=""
	switch atyp {
	case atypeIPV4:
		//四个字节
		_, err := io.ReadFull(reader, buf)
		if err != nil {
			return fmt.Errorf("read atyp failed:%w", err)
		}
		addr = fmt.Sprintf("%d.%d.%d.%d", buf[0], buf[1], buf[2], buf[3])
	case atypeHOST:
		hostSize,err:=reader.ReadByte()
		if err!=nil{
			return fmt.Errorf("read hostsize failed:%w",err)
		}
		host:=make([]byte,hostSize)
		_,err=io.ReadFull(reader,host)
		if err!=nil{
			return fmt.Errorf("read host failed:%w",err)
		}
		addr=string(host)
	case atypeIPV6:
		return errors.New("IPV6:NOT SUPPORTED NOW")
	default:
		return errors.New("invalid atype")

	}
	_,err=io.ReadFull(reader,buf[:2])
	if err!=nil{
		return fmt.Errorf("read port failed:%w",err)
	}
	port:=binary.BigEndian.Uint16(buf[:2])

	//与真正的服务器建立tcp连接
	dest,err:= net.Dial("tcp",fmt.Sprintf("%v:%v",addr,port))
	if err!=nil{
		return fmt.Errorf("dial dst failed:%w",err)
	}
	defer dest.Close()


	log.Println("dial",addr,port)

	//response报文
	_,err=conn.Write([]byte{0x05,0x00,0x00,0x01,0,0,0,0,0,0})
	if err!=nil{
		return fmt.Errorf("write failed:%w",err)
	}
	//建立浏览器和下游服务器的双向数据转发

	//WithCancel 返回具有新 Done 通道的 parent 副本。返回的上下文的完成通道在调用返回的取消函数或父上下文的完成通道关闭时关闭，以先发生者为准。
	//取消此上下文会释放与其关联的资源，因此代码应在此上下文中运行的操作完成后立即调用取消
	//goroutine速度很快 该函数会立即返回 返回时连接就关闭了
	//但需要copy函数出错时返回 因此使用context机制
	ctx,cancel:=context.WithCancel(context.Background())
	defer cancel()
	go func(){
		//从浏览器拷贝到服务器
		_,_=io.Copy(dest,reader)
		//调用cancel()通知后台Goroutine退出，避免泄露
		cancel()
	}()
	go func(){
		//从服务器拷贝到浏览器
		_,_=io.Copy(conn,dest)
		cancel()
	}()
	//代表退出协程 执行完成后cancel()被调用
	<-ctx.Done()





	return nil

}

