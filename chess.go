package main

import "fmt"

func drawTab(tab []byte) {
	fmt.Println("+---+---+---+---+---+---+---+---+")
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Printf("| %c ", tab[i*8 + j])
		}
		fmt.Println("|\n+---+---+---+---+---+---+---+---+")
	}
}

func checkToken(token byte, tipo byte) bool{
	if tipo == 'm'{
		if token >= 97 && token <= 122{
			return true
		}
	}
	if tipo == 'M'{
		if token >= 65 && token <= 90{
			return true
		}
	}
	return false
}

func scanJugada(tab []byte, p byte) {
	var i, j, token byte
	valid := false
	for !valid {
		fmt.Printf("Ficha elegida %c [0-7] [0-7]: ", rune(p))
		fmt.Scanf("%d %d\n", &i, &j)
		idx := i + j * 8
		if i >= 0 && i < 8 && j >= 0 && j < 8 && checkToken(tab[idx], p) {
			token = tab[idx]
			tab[idx] = 0
			valid = true
		} else {
			fmt.Println(" --- No hay ficha permitida ---")
		}
	}
	valid2 := false
	for !valid2  {
		fmt.Printf("Nueva posicion %c [0-7] [0-7]: ", rune(p))
		fmt.Scanf("%d %d\n", &i, &j)
		idx := i  + j * 8
		if i >= 0 && i < 8 && j >= 0 && j < 8 && !checkToken(tab[idx], p) {
			tab[idx] = token
			valid2 = true
		} else {
			fmt.Println(" --- Jugada no permitida ---")
		}
	}
}

func findWinner(tab []byte) byte {
	var contm, contM int
	for i := 0; i < 64; i++ {
		if checkToken(tab[i],109){
			contm++
		}
		if checkToken(tab[i],77){
			contM++
		}
	}
	if contm == 0 {
		return 77
	}
	if contM == 0{
		return 109
	}
	return 0
}

func chooseOpositeToken(tab [] byte) byte {
	for i := 0; i < 16; i++ {
		if tab[i] == 0 {
			return 77
		}
	}
	for i := 48; i < 64; i++{
		if tab[i] == 0{
			return 109
			}
	}
	return 0
}

func pickToken() byte {
	p := '-'
	for p != 'm' && p != 'M' {
		fmt.Println("Seleccione ficha [m,M]: ")
		fmt.Scanf("%c\n", &p)
		if p != 'm' && p != 'M' {
			fmt.Println(" --- Ficha no permitida ---")
		}
	}
	return byte(p)
}

func getMsg(buff []byte) byte {
	return buff[0]
}
func getSessId(buff []byte) byte {
	return buff[1]
}
func getPlayerId(buff []byte) byte {
	return buff[2]
}
func getTab(buff []byte) []byte {
	return buff[3:]
}
func setMsg(buff []byte, msg byte) {
	buff[0] = msg
}
func setSessId(buff []byte, sid byte) {
	buff[1] = sid
}
func setPlayerId(buff []byte, pid byte) {
	buff[2] = pid
}
func setTab(buff []byte, tab []byte) {
	for i, e := range tab {
		buff[i + 3] = e
	}
}

const (
	NEW = byte(0) // Mensajes del cliente
	UPDATE = byte(1)
	PLAY = byte(2)

	WAIT = byte(3)  // Mensajes del server
	TURN = byte(4)
)