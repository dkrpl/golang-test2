package pasien

import (
	helpers "golang-test2/helper"
	schemas "golang-test2/schemas/pasien"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (this_ *ControllerPasienImpl) AddKamar(ctx *gin.Context) {
	var schemas schemas.AddKamar
	if ctx.BindJSON(&schemas) != nil {
		helpers.APIResponse(ctx, "error binding json", http.StatusForbidden, http.MethodPost, nil)
		return
	}
	id := ctx.Param("id")
	resp, err := this_.PasienUsecase.AddKamar(id, schemas)
	if err != nil {
		helpers.APIResponse(ctx, err.Error(), http.StatusForbidden, http.MethodPost, nil)
		return
	}
	helpers.APIResponse(ctx, "Update Ok", http.StatusOK, http.MethodPost, resp)
}
