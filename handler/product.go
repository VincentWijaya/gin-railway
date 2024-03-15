package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) FindOneProduct(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Malformat ID",
		})
		return
	}

	res, err := h.productService.FindOneProduct(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Failed to find product",
			"detail":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
