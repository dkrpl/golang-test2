package kamar

import (
	helpers "golang-test2/helper"
	schemas "golang-test2/schemas/kamar"
	usecase "golang-test2/usecases/kamar"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewControllerKamarImpl(uc *usecase.KamarUsecase) ControllerKamar {
	return &ControllerKamarImpl{
		Usecase: *uc,
	}
}

type ControllerKamarImpl struct {
	Usecase usecase.KamarUsecase
}

// @Summary Create Feature
// @Description Create Feature
// @Accept  json
// @Produce  json
// @Param Body body schemas.Input true "Feature Create Request"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "error binding json"
// @Failure 403 {string} string "error when after models return"
// @Router /feature/create [post]
// @Tags Feature
func (this_ *ControllerKamarImpl) Create(ctx *gin.Context) {
	var schema schemas.Input
	if ctx.BindJSON(&schema) != nil {
		helpers.APIResponse(ctx, "invalid postform", http.StatusAccepted, http.MethodPost, nil)
		return
	}
	resp, err := this_.Usecase.Create(schema)
	if err != nil {
		helpers.APIResponse(ctx, err.Error(), http.StatusAccepted, http.MethodPost, resp)
		return
	}
	helpers.APIResponse(ctx, "Ok", http.StatusOK, http.MethodPost, resp)
}

// @Summary List Feature
// @Description List Feature
// @Accept  json
// @Produce  json
// @Param search query string true "search name"
// @Param filter query string true "filter | EX: filter=role_access|building:tenant,role_name|building manager"
// @Param per_page query string true "per_page"
// @Param page_no query string true "page_no"
// @Success 200 {array} []schemas.List "ok"
// @Failure 400 {string} string "error binding json"
// @Failure 403 {string} string "error when after models return"
// @Router /feature/list [get]
// @Tags Feature
func (this_ *ControllerKamarImpl) List(ctx *gin.Context) {
	search := ctx.Query("search")
	filter := ctx.Query("filter")
	per_page := ctx.Query("per_page")
	page_no := ctx.Query("page_no")

	resp, count, err := this_.Usecase.List(search, filter, per_page, page_no)
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

// @Summary Detail Feature
// @Description Detail Feature
// @Accept  json
// @Produce  json
// @Param featureID path string true "featureID"
// @Success 200 {object} schemas.Detail "ok"
// @Failure 400 {string} string "error binding json"
// @Failure 403 {string} string "error when after models return"
// @Router /feature/get/{featureID} [get]
// @Tags Feature
func (this_ *ControllerKamarImpl) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	resp, err := this_.Usecase.Get(id)
	if err != nil {
		helpers.APIResponse(ctx, err.Error(), http.StatusAccepted, http.MethodGet, nil)
		return
	}
	helpers.APIResponse(ctx, "Ok", http.StatusOK, http.MethodGet, resp)
}

// @Summary Remove Feature
// @Description You can delete a feature with a note that the feature is not used in all cities
// @Accept  json
// @Produce  json
// @Param featureID path string true "cityID"
// @Success 200 {object} schemas.Delete_Response "ok"
// @Failure 400 {string} string "error binding json"
// @Failure 403 {string} string "error when after models return"
// @Router /feature/delete/{featureID} [delete]
// @Tags Feature
func (this_ *ControllerKamarImpl) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := this_.Usecase.Delete(id)
	if err != nil {
		helpers.APIResponse(ctx, err.Error(), http.StatusForbidden, http.MethodPost, nil)
		return
	}
	helpers.APIResponse(ctx, "Delete Ok", http.StatusOK, http.MethodPost, response)
}

// @Summary Remove Feature
// @Description Remove Feature
// @Accept  json
// @Produce  json
// @Param featureID path string true "featureID"
// @Param staffBody body schemas.Edit true "City Update Request"
// @Success 200 {object} schemas.Edit_Response "ok"
// @Failure 400 {string} string "error binding json"
// @Failure 403 {string} string "error when after models return"
// @Router /feature/update/{featureID} [put]
// @Tags Feature
func (this_ *ControllerKamarImpl) Edit(ctx *gin.Context) {
	var schemas schemas.Edit
	if ctx.BindJSON(&schemas) != nil {
		helpers.APIResponse(ctx, "error binding json", http.StatusForbidden, http.MethodPost, nil)
		return
	}
	id := ctx.Param("id")
	resp, err := this_.Usecase.Edit(id, schemas)
	if err != nil {
		helpers.APIResponse(ctx, err.Error(), http.StatusForbidden, http.MethodPost, nil)
		return
	}
	helpers.APIResponse(ctx, "Update Ok", http.StatusOK, http.MethodPost, resp)
}
