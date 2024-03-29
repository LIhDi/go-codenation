package algoritmos

import (
	"fmt"
	"crypto/sha1"
	"encoding/hex"
	"codenation/models"
)

type Algoritmo struct{}

func NewAlgoritmo() *Algoritmo {
	return &Algoritmo{}
}

func (p Algoritmo) DescriptografaJulioCesar(answer models.Answer) int {
	ponto := 46
	espaco := 32
	//teste := "uif cftu xbz up hfu b qspkfdu epof gbtufs jt up tubsu tppofs kjn ijhitnjui"
	var decifrado string

	for _, caractere := range answer.Cifrado {
		if int(caractere) == ponto || int(caractere) == espaco {
			decifrado = decifrado + string(caractere)
		} else{
			decifrado = decifrado + string(int(caractere)-2)
		}
	}

    hash := sha1.New()
    hash.Write([]byte(decifrado))
    resumo_criptografico := hex.EncodeToString(hash.Sum(nil))

	fmt.Println(decifrado, resumo_criptografico)
	
	return 0
}