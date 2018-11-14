package handler

import (
	"log"
	"net/http"
	"strconv"
	"vue-golang-payment-app/backend-api/db"
)

// GetLists hogehoge
func GetLists(c Context) {
	res, err := db.SelectAllItems()

	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetItem idでitemげっつ
func GetItem(c Context) {
	log.SetFlags(log.Llongfile)
	identifer, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	res, err := db.SelectItem(int64(identifer))
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, res)
}
