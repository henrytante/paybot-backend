package main

import (
	"os/exec"
)

func main() {
	cmd := exec.Command("sleep", "10000")
	if err := cmd.Start(); err != nil {
		// Lida com o erro se houver algum problema ao iniciar o comando
		panic(err)
	}

	// Espera a conclusão do comando
	if err := cmd.Wait(); err != nil {
		// Lida com o erro se houver algum problema durante a execução do comando
		panic(err)
	}
}
