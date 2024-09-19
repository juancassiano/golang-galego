package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositório de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

// Criar insere um usuário no banco de dados
func (repositorio usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, err := repositorio.db.Prepare(
		"INSERT INTO usuarios (nome, nick, email, senha) VALUES (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()
	resultado, err := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if err != nil {
		return 0, err
	}
	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(ultimoIDInserido), nil
}

// Buscar traz todos os usuários que atendem o filtro de Nome ou Nick
func (repositorio usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) //%nomeOuNick%%
	// log.Printf("Filtro: %s", nomeOuNick)

	linhas, err := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome LIKE ? or nick LIKE ?",
		nomeOuNick, nomeOuNick)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario
		if err = linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

// BuscarPorID traz um usuário do banco de dados
func (repositorio usuarios) BuscarPorID(ID uint64) (modelos.Usuario, error) {
	linhas, err := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where id = ?",
		ID,
	)
	if err != nil {
		return modelos.Usuario{}, err
	}
	defer linhas.Close()
	var usuario modelos.Usuario
	if linhas.Next() {
		if err = linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); err != nil {
			return modelos.Usuario{}, err
		}
	}
	return usuario, nil
}

// Atualizar altera as informações de um usuário no banco de dados
func (repositorio usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error {
	statement, err := repositorio.db.Prepare(
		"update usuarios set nome = ?, nick = ?, email = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); err != nil {
		return err
	}
	return nil
}

// Deletar remove um usuário do banco de dados
func (repositorio usuarios) Deletar(ID uint64) error {
	statement, err := repositorio.db.Prepare("delete from usuarios where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(ID); err != nil {
		return err
	}
	return nil
}

// BuscarPorEmail traz um usuário do banco de dados
func (repositorio usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linha, err := repositorio.db.Query("select id, senha from usuarios where email = ?", email)
	if err != nil {
		return modelos.Usuario{}, err
	}
	defer linha.Close()

	var usuario modelos.Usuario
	if linha.Next() {
		if err = linha.Scan(&usuario.ID, &usuario.Senha); err != nil {
			return modelos.Usuario{}, err
		}
	}
	return usuario, nil
}

// Seguir permite que um usuário siga outro
func (repositorio usuarios) Seguir(usuarioID, seguidorID uint64) error {
	statement, err := repositorio.db.Prepare(
		"insert ignore into seguidores (usuario_id, seguidor_id) values (?,?)",
	)
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(usuarioID, seguidorID); err != nil {
		return err
	}
	return nil
}

// PararDeSeguir permite que um usuário deixe de seguir outro
func (repositorio usuarios) PararDeSeguir(usuarioID, seguidorID uint64) error {
	statement, err := repositorio.db.Prepare(
		"delete from seguidores where usuario_id = ? and seguidor_id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(usuarioID, seguidorID); err != nil {
		return err
	}
	return nil
}

// BuscarSeguidores traz todos os seguidores de um usuário
func (repositorio usuarios) BuscarSeguidores(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, err := repositorio.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.criadoEm
		from usuarios u inner join seguidores s on u.id = s.seguidor_id where s.usuario_id = ?
	`, usuarioID)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()
	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario
		if err = linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

// BuscarSeguindo traz todos os usuários que um determinado usuário está seguindo
func (repositorio usuarios) BuscarSeguindo(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, err := repositorio.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.criadoEm
		from usuarios u inner join seguidores s on u.id = s.usuario_id where s.seguidor_id = ?
	`, usuarioID)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()
	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario
		if err = linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

// BuscarSenha traz a senha de um usuário do banco de dados
func (repositorio usuarios) BuscarSenha(usuarioID uint64) (string, error) {
	linha, err := repositorio.db.Query("select senha from usuarios where id = ?", usuarioID)
	if err != nil {
		return "", err
	}
	defer linha.Close()
	var usuario modelos.Usuario
	if linha.Next() {
		if err = linha.Scan(&usuario.Senha); err != nil {
			return "", err
		}
	}
	return usuario.Senha, nil
}

// AtualizarSenha altera a senha de um usuário no banco de dados
func (repositorio usuarios) AtualizarSenha(usuarioID uint64, senha string) error {
	statement, err := repositorio.db.Prepare("update usuarios set senha = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(senha, usuarioID); err != nil {
		return err
	}
	return nil
}
