package pasien

import (
	helpers "golang-test2/helper"
	schemas "golang-test2/schemas/pasien"
	pasien_usecase "golang-test2/usecases/pasien"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ControllerPasienImpl struct {
	PasienUsecase pasien_usecase.PasienUsecase
}

type ControllerKamarImpl struct {
	PasienUsecase pasien_usecase.PasienUsecase
}

func NewControllerPasienImpl(usecase *pasien_usecase.PasienUsecase) ControllerPasien {
	return &ControllerPasienImpl{
		PasienUsecase: *usecase,
	}
}

// @Summary Create City
// @Description Create City
// @Accept  json
// @Produce  json
// @Param staffID path string true "staffID"
// @Param staffBody body schemas.Input true "City Create Request"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "error binding json"
// @Failure 403 {string} string "error when after models return"
// @Router /city/create [post]
// @Tags City
func (this_ *ControllerPasienImpl) Create(ctx *gin.Context) {
	var schema schemas.Input

	file, _ := ctx.FormFile("kecamatan")
	schema.Nama_pasien = ctx.PostForm("nama_pasien")
	schema.Alamat = ctx.PostForm("alamat")
	schema.No_rumah, _ = strconv.Atoi(ctx.PostForm("no_rumah"))
	schema.Kabupaten = ctx.PostForm("kabupaten")
	if schema == (schemas.Input{}) {
		helpers.APIResponse(ctx, "invalid postform", http.StatusAccepted, http.MethodPost, nil)
		return
	}
	resp, err := this_.PasienUsecase.Create(schema, file)
	if err != nil {
		helpers.APIResponse(ctx, err.Error(), http.StatusAccepted, http.MethodPost, resp)
		return
	}
	helpers.APIResponse(ctx, "Ok", http.StatusOK, http.MethodPost, resp)
}

// @Summary List City
// @Description List City
// @Accept  json
// @Produce  json
// @Param search query string true "search name"
// @Param filter query string true "filter | EX: filter=role_access|building:tenant,role_name|building manager"
// @Param per_page query string true "per_page"
// @Param page_no query string true "page_no"
// @Success 200 {array} []schemas.List "ok"
// @Failure 400 {string} string "error binding json"
// @Failure 403 {string} string "error when after models return"
// @Router /city/list [get]
// @Tags City
func (this_ *ControllerPasienImpl) List(ctx *gin.Context) {
	search := ctx.Query("search")
	filter := ctx.Query("filter")
	per_page := ctx.Query("per_page")
	page_no := ctx.Query("page_no")

	resp, count, err := this_.PasienUsecase.List(search, filter, per_page, page_no)
	per_page_int, _ := strconv.Atoi(per_page)
	page_no_int, _ := strconv.Atoi(page_no)
	if err != nil {
		helpers.APIResponseList(
			ctx,
			err.Error(),
			http.StatusAccepted,
			http.MethodGet,
			[]interface{}{},
			count,
			per_page_int,
			page_no_int,
		)
		return
	}
	helpers.APIResponseList(
		ctx,
		"Ok",
		http.StatusOK,
		http.MethodGet,
		resp,
		count,
		per_page_int,
		page_no_int,
	)
}

// @Summary Detail City
// @Description Detail City
// @Accept  json
// @Produce  json
// @Param cityID path string true "cityID"
// @Success 200 {object} schemas.Get_Response "ok"
// @Failure 400 {string} string "error binding json"
// @Failure 403 {string} string "error when after models return"
// @Router /city/detail/{cityID} [get]
// @Tags City
func (this_ *ControllerPasienImpl) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	resp, err := this_.PasienUsecase.Get(id)
	if err != nil {
		helpers.APIResponse(ctx, err.Error(), http.StatusAccepted, http.MethodGet, nil)
		return
	}
	helpers.APIResponse(ctx, "Ok", http.StatusOK, http.MethodGet, resp)
}

// @Summary Remove City
// @Description Remove City
// @Accept  json
// @Produce  json
// @Param cityID path string true "cityID"
// @Success 200 {object} schemas.Delete_Response "ok"
// @Failure 400 {string} string "error binding json"
// @Failure 403 {string} string "error when after models return"
// @Router /city/delete/{cityID} [delete]
// @Tags City
func (this_ *ControllerPasienImpl) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := this_.PasienUsecase.Delete(id)
	if err != nil {
		helpers.APIResponse(ctx, err.Error(), http.StatusForbidden, http.MethodPost, nil)
		return
	}
	helpers.APIResponse(ctx, "Delete Ok", http.StatusOK, http.MethodPost, response)
}

// @Summary Remove City
// @Description Remove City
// @Accept  json
// @Produce  json
// @Param cityID path string true "cityID"
// @Param staffBody body schemas.Edit true "City Update Request"
// @Success 200 {object} schemas.Edit_Response "ok"
// @Failure 400 {string} string "error binding json"
// @Failure 403 {string} string "error when after models return"
// @Router /city/update/{cityID} [put]
// @Tags City
func (this_ *ControllerPasienImpl) Edit(ctx *gin.Context) {
	var schemas schemas.Edit
	if ctx.BindJSON(&schemas) != nil {
		helpers.APIResponse(ctx, "error binding json", http.StatusForbidden, http.MethodPost, nil)
		return
	}
	id := ctx.Param("id")
	resp, err := this_.PasienUsecase.Edit(id, schemas)
	if err != nil {
		helpers.APIResponse(ctx, err.Error(), http.StatusForbidden, http.MethodPost, nil)
		return
	}
	helpers.APIResponse(ctx, "Update Ok", http.StatusOK, http.MethodPost, resp)
}
