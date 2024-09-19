package main

func main() {
	// i := 0

	// for i < 10 {
	// 	println(i)
	// 	time.Sleep(time.Second)
	// 	i++
	// }

	// for j := 0; j < 10; j++ {
	// 	println(j)
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"João", "José", "Maria"}

	for k, nome := range nomes {
		println(k, nome)
	}
	for _, nome := range nomes {
		println(nome)
	}

	usuario := map[string]string{
		"nome":      "João",
		"sobrenome": "Silva",
	}
	for chave, valor := range usuario {
		println(chave, valor)
	}

}
