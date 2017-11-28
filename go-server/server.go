package main

import (
	"fmt"
	"net"
	"sync"
)

const (
	ip = "127.0.0.1"
)

type conn_pro struct {
	seq      uint16
	msg_send chan []byte
}

// 连接的客户端
var map_lock = struct{
	lock sync.RWMutex
	cli_conn_map map[*net.TCPConn](*conn_pro)
}{cli_conn_map:make(map[*net.TCPConn](*conn_pro))}



// 全局静态 缓冲区声明
var msg_r = make(chan []byte, 100) //10000

//定义消息缓冲队列长度为100字节
// 包头内容为 Hello Go!
const (
	CYCLIBUFSIZE  = 100 //1024 * 2 * 4
	PACKHEAD = "Hello Go!"
)

// 消息包的结构是 PACKHEAD + 2字节消息内容长度 + 消息内容

func main() {

	go HandlConn()
	HandlData()
}

func HandlConn() {
	listen, err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP(ip), 3333, ""})
	if err != nil {
		fmt.Println("监听端口失败", err.Error())
		return
	}

	fmt.Println("已初始化链接，等待客户端链接。。。")

	seq := uint16(1)

	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println("接受客户端连接异常:", err.Error())
			continue
		}

		prop := conn_pro{seq,make(chan []byte, 1024 * 2)}
		map_lock.lock.Lock()
		map_lock.cli_conn_map[conn] = &prop
		map_lock.lock.Unlock()


		seq++

		fmt.Printf("客户端 %d 连接到服务器\n", seq - 1)
		fmt.Println(conn)

		defer conn.Close()

		go RecvMsg(conn)

		go SendMsg(conn)

	}
}

func RecvMsg(conn *net.TCPConn) {
	CycliBuffer := make([]byte, CYCLIBUFSIZE)
	cur_curcor, last_curcor := 0, 0

	for {
		data := make([]byte, 100) // 1024*4
		map_lock.lock.Lock()
		cli_seq:= map_lock.cli_conn_map[conn].seq
		map_lock.lock.Unlock()
		i, err := conn.Read(data)
		if err != nil {
			fmt.Printf("客户端 %d 断开连接\n", cli_seq)
			map_lock.lock.Lock()
			delete(map_lock.cli_conn_map, conn)
			map_lock.lock.Unlock()
			break
		}


		// 循环缓冲区接收read 数据
		j := 0
		for ; j < i; j++ {
			CycliBuffer[(j + last_curcor) % CYCLIBUFSIZE] = data[j]
		}
		last_curcor = (j + last_curcor) % CYCLIBUFSIZE

		Unpack(CycliBuffer, cli_seq, &cur_curcor)
	}

}

// 消息发送给客户端
func SendMsg(conn *net.TCPConn) {

	for {
		map_lock.lock.Lock()
		if map_lock.cli_conn_map[conn] == nil {
			map_lock.lock.Unlock()
			break
		}


		select {
		case msg := <- map_lock.cli_conn_map[conn].msg_send:
			conn.Write(msg)

			fmt.Printf("cli- %d says: %s \n", uint16(msg[0]) << 8 + uint16(msg[1]), string(msg[2:]))
		default:

		}

		map_lock.lock.Unlock()
	}

}

// 解包
func Unpack(data []byte, seq uint16, cur_curcor *int){

	PackHead := make([]byte,2)
	PackHead[0] = byte(seq >> 8)
	PackHead[1] = byte(seq % 256)

	head_len := len(PACKHEAD)


	for i := *cur_curcor; ; {
		pack_head := ""
		if i + head_len > CYCLIBUFSIZE && i < CYCLIBUFSIZE {
			pack_head = string(data[i:]) + string(data[:(i + head_len) % CYCLIBUFSIZE])
		} else if i + head_len <= CYCLIBUFSIZE {
			pack_head = string(data[i : i + head_len])
		}else if i >= CYCLIBUFSIZE {
			pack_head = string(data[i % CYCLIBUFSIZE : (i + head_len) % CYCLIBUFSIZE])
		}


		if PACKHEAD == pack_head {
			msg_size := int(data[(i + head_len) % CYCLIBUFSIZE]) << 8 + int(data[(i + head_len + 1) % CYCLIBUFSIZE])
			fmt.Println("消息字节大小：",msg_size)

			msg_content := ""
			if i + head_len + 2 >= CYCLIBUFSIZE {
				cur := (i + head_len + 2) % CYCLIBUFSIZE
				if cur + msg_size  > CYCLIBUFSIZE {
					msg_content = string(data[cur:]) + string(data[:(i + head_len + msg_size + 2) % CYCLIBUFSIZE])
				} else {
					msg_content = string(data[cur : (i + head_len + msg_size + 2) % CYCLIBUFSIZE ])
				}
			} else {
				if (i + head_len + msg_size + 2) > CYCLIBUFSIZE {
					msg_content = string(data[(i + head_len + 2):]) + string(data[:(i + head_len + msg_size + 2) % CYCLIBUFSIZE])
				} else {
					msg_content = string(data[(i + head_len + 2) : (i + head_len + msg_size + 2)])
				}
			}

			fmt.Printf("客户端 %d 发来的消息：%s", seq, string(msg_content))
			msg_r <- append(PackHead[0:2], msg_content...)
			i = (i + head_len + 2 + msg_size) % CYCLIBUFSIZE

		} else {
			*cur_curcor = i % CYCLIBUFSIZE
			break
		}

	}
}

// 消息包传到各个连接的发送队列
func HandlData(){
	for {
		select {
		case  msg := <- msg_r:
			map_lock.lock.Lock()
			for _, conn_pro := range map_lock.cli_conn_map {
				conn_pro.msg_send <- msg
			}
			map_lock.lock.Unlock()
		}
	}
}