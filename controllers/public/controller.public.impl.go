package public

import (
	usecases "golang-test2/usecases/public"

	"github.com/gin-gonic/gin"
)

func NewControllerPublicImpl(usecase *usecases.PublicUsecase) ControllerPublic {
	return &ControllerPublicImpl{
		Usecase: *usecase,
	}
}

type ControllerPublicImpl struct {
	Usecase usecases.PublicUsecase
}

// @Summary Root
// @Description take your config when you first access the website
// @Accept  json
// @Produce  json
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "error binding json"
// @Failure 403 {string} string "error when after models return"
// @Router / [get]
// @Tags Root
func (this_ *ControllerPublicImpl) Root(ctx *gin.Context) {
	root, err := this_.Usecase.Root("")
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": root,
	})
}
