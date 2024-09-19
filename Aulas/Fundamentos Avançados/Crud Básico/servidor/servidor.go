package servidor

import (
	"crud-basico/banco"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json: "id"`
	Nome  string `json: "nome"`
	Email string `json: "email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	var user usuario
	if err = json.Unmarshal(corpoRequisicao, &user); err != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("insert into usuarios (nome, email) values(?,?)")
	if err != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}
	defer statement.Close()
	insercao, err := statement.Exec(user.Nome, user.Email)
	if err != nil {
		w.Write([]byte("Erro ao executar o statement!"))
		return
	}

	_, err = insercao.LastInsertId()
	if err != nil {
		w.Write([]byte("Erro ao obter o ID!"))
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}
	defer db.Close()

	rows, err := db.Query("select * from usuarios")
	if err != nil {
		w.Write([]byte("Erro ao buscar os usuários!"))
		return
	}
	defer rows.Close()

	var users []usuario
	for rows.Next() {
		var user usuario
		if err := rows.Scan(&user.ID, &user.Nome, &user.Email); err != nil {
			w.Write([]byte("Erro ao escanear o usuário!"))
			return
		}
		users = append(users, user)
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("Erro ao converter os usuários para JSON"))
		return
	}

}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, err := strconv.ParseUint(parametros["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}
	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}
	defer db.Close()
	var user usuario
	linha, err := db.Query("select * from usuarios where id = ?", ID)
	if err != nil {
		w.Write([]byte("Erro ao buscar o usuário!"))
		return
	}
	defer linha.Close()
	if linha.Next() {
		if err := linha.Scan(&user.ID, &user.Nome, &user.Email); err != nil {
			w.Write([]byte("Erro ao escanear o usuário!"))
			return
		}
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte("Erro ao converter o usuário para JSON"))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, err := strconv.ParseUint(parametros["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}
	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	var user usuario
	if err := json.Unmarshal(corpoRequisicao, &user); err != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}
	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}
	defer db.Close()
	statement, err := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if err != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}
	defer statement.Close()
	if _, err := statement.Exec(user.Nome, user.Email, ID); err != nil {
		w.Write([]byte("Erro ao atualizar o usuário!"))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, err := strconv.ParseUint(parametros["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}
	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}
	defer db.Close()
	statement, err := db.Prepare("delete from usuarios where id = ?")
	if err != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}
	defer statement.Close()
	if _, err := statement.Exec(ID); err != nil {
		w.Write([]byte("Erro ao deletar o usuário!"))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
