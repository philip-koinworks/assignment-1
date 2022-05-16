package model

import (
	"encoding/json"
)

type student struct {
	Nama      string `json:"nama"`
	Alamat    string `json:"alamat"`
	Pekerjaan string `json:"pekerjaan"`
	Alasan    string `json:"alasan"`
}

var students []student = []student{
	{"philip", "jl. krendang selatan", "backend engineer", "kebutuhan perusahaan"},
	{"tasrifin", "jl. super duper", "backend engineer", "tertarik dengan bahasa go"},
}

func GetStudent(idx int) student {
	return students[idx-1]
}

func (s student) ToJson() (string, error) {
	j, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(j), nil
}
