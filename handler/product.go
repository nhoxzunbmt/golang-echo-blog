package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/xesina/golang-echo-realworld-example-app/model"
)

func (h *Handler) Products(c echo.Context) error {
	var (
		articles []model.Article
		count    int
	)

	tag := c.QueryParam("tag")
	author := c.QueryParam("author")
	favoritedBy := c.QueryParam("favorited")

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 20
	}

	if tag != "" {
		articles, count, err = h.articleStore.ListByTag(tag, offset, limit)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}
	} else if author != "" {
		articles, count, err = h.articleStore.ListByAuthor(author, offset, limit)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}
	} else if favoritedBy != "" {
		articles, count, err = h.articleStore.ListByWhoFavorited(favoritedBy, offset, limit)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}
	} else {
		articles, count, err = h.articleStore.List(offset, limit)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}
	}

	return c.JSON(http.StatusOK, newArticleListResponse(h.userStore, userIDFromToken(c), articles, count))
}
