package repositorys

import (
	"golang-test2/repositorys/kamar"
	"golang-test2/repositorys/pasien"

	"gopkg.in/mgo.v2"
)

type Repositorys struct {
	RepositoryPasien *pasien.PasienRepository
	// RepositoryUser    *users.UserRepository
	RepositoryKamar *kamar.KamarRepository
}

func NewRepository(db *mgo.Database) Repositorys {
	repositoryPasien := pasien.NewPasienRepository(db)
	// repositoryUser := users.NewUserRepository(db)
	repositoryKamar := kamar.NewKamarRepository(db)
	return Repositorys{
		RepositoryPasien: &repositoryPasien,
		// RepositoryUser:    &repositoryUser,
		RepositoryKamar: &repositoryKamar,
	}
}
