package main

func main() {
	conn := bind()
	defer conn.Close()

	go listen(conn)

	calculator()
}
