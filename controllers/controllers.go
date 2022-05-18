package controllers

import (
	"golang-test2/controllers/kamar"
	"golang-test2/controllers/pasien"
	"golang-test2/controllers/public"
	"golang-test2/usecases"
)

type Controllers struct {
	ControllerPasien pasien.ControllerPasien
	// ControllerUser    user.UserController
	ControllerKamar  kamar.ControllerKamar
	ControllerPublic public.ControllerPublic
}

func MainController(usecase usecases.Usecases) Controllers {
	return Controllers{
		ControllerPasien: pasien.NewControllerPasienImpl(&usecase.UsecasePasien),
		// ControllerUser:    user.NewUserController(&usecase.UseCaseUser),
		ControllerKamar:  kamar.NewControllerKamarImpl(&usecase.UseCaseKamar),
		ControllerPublic: public.NewControllerPublicImpl(&usecase.UseCasePublic),
	}
}
