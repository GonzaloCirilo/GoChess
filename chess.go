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

func scanJugada(tab []byte, p byte) {
	var i, j byte
	valid := false
	for !valid {
		fmt.Printf("Ficha elegida %c [0-7] [0-7]: ", rune(p))
		fmt.Scanf("%d %d\n", &i, &j)
		idx := i + j * 8
		if i >= 0 && i < 8 && j >= 0 && j < 8 && tab[idx] == p {
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
		if i >= 0 && i < 8 && j >= 0 && j < 8 && tab[idx] != p {
			tab[idx] = p
			valid2 = true
		} else {
			fmt.Println(" --- Jugada no permitida ---")
		}
	}
}

func findWinner(tab []byte) byte {
	/*for i := 0; i < 8; i++ {
		if (tab[i*8] == tab[i*8+1] && tab[i*8+1] == tab[i*3+2]) ||
		   (tab[i] == tab[i+3] && tab[i+3] == tab[i+6]) {
			return tab[i*4]
		}
	}
	if (tab[0] == tab[4] && tab[4] == tab[8]) ||
	   (tab[2] == tab[4] && tab[4] == tab[6]) {
	   	return tab[4]
	}
	for i := 0; i < 9; i++ {
		if tab[i] != 0 {
			return 0;
		}
	}
	return byte('-')*/ // empate
	return 0
}

func chooseOpositeToken(tab [] byte) byte {
	var token byte
	for i := 16; i < 48; i++ {
		if tab[i] != 0 {
			token = tab[i]
			break
		}
	}
	if token == 111 {
		return 120
	} else if token == 120 {
		return 111
	}
	return 0
}

func pickToken() byte {
	p := '-'
	for p != 'o' && p != 'x' {
		fmt.Println("Seleccione ficha [x,o]: ")
		fmt.Scanf("%c\n", &p)
		if p != 'o' && p != 'x' {
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