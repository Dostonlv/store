package handler

import (
	"app/api/models"
	"context"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePromo(ctx *gin.Context) {
	var createPromo models.CreatePromoCode

	err := ctx.ShouldBindJSON(&createPromo)
	if err != nil {
		h.handlerResponse(ctx, "create promo", 400, err.Error())
		return
	}

	id, err := h.storages.PromoCode().Create(context.Background(), &createPromo)
	if err != nil {
		h.handlerResponse(ctx, "storage.promo.create", 500, err.Error())
		return
	}

	resp, err := h.storages.PromoCode().GetByID(context.Background(), &models.PromoCodePrimaryKey{ID: id})
	if err != nil {
		h.handlerResponse(ctx, "storage.promo.getByID", 500, err.Error())
		return
	}

	h.handlerResponse(ctx, "create promo", 201, resp)

}
