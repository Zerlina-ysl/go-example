package main

go func(){
	for {
		conn,_:=listener.Accept()
		go func(){
			conn.Read(request)
			...
			conn.Write(response)
		}
	}
}

go func(){
	for{
		readableConn,_ :=Monitor(conns)
		for conn:=range radableConn{
			go func(){
				conn.Read(request)
				...
				conn.Write(response)
			}
		}
	}
}
