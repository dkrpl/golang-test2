package usecases

import (
	"golang-test2/middleware"
	"golang-test2/repositorys"
	"golang-test2/usecases/kamar"
	"golang-test2/usecases/pasien"
	"golang-test2/usecases/public"
)

type Usecases struct {
	UsecasePasien pasien.PasienUsecase
	// UseCaseUser    user.UserUsecase
	UseCaseKamar  kamar.KamarUsecase
	UseCasePublic public.PublicUsecase
}

func MainUsecase(repository repositorys.Repositorys, dynamic *middleware.Dynamic) Usecases {
	var all_Usecase Usecases

	all_Usecase.UsecasePasien = pasien.NewUsecasePasienImpl(repository.RepositoryPasien)
	// all_Usecase.UseCaseUser = user.NewUsecaseUserImpl(repository.RepositoryUser)
	all_Usecase.UseCaseKamar = kamar.NewUsecaseKamar(&repository)
	all_Usecase.UseCasePublic = public.NewPublicUsecase(&repository, dynamic)
	return all_Usecase
}
