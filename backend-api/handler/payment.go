package handler

import (
	"context"
	"net/http"
	"strconv"
	"vue-golang-payment-app/backend-api/db"
	"vue-golang-payment-app/backend-api/domain"
	gpay "vue-golang-payment-app/payment-service/proto"

	"google.golang.org/grpc"
)

var addr = "localhost:50051"

// Charge ほげほげ
func Charge(c Context) {
	// パラメータやbodyを受け取る
	t := domain.Payment{}
	c.Bind(&t)
	identifer, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	// idからitem情報取得
	res, err := db.SelectItem(int64(identifer))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	greq := &gpay.PayRequest{
		Id:          int64(identifer),
		Token:       t.Token,
		Amount:      res.Amount,
		Name:        res.Name,
		Description: res.Description,
	}

	// IPアドレス(ここではlocalhost)とポート番号(ここでは5000)を指定して、サーバーと接続する
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		c.JSON(http.StatusForbidden, err)
		return
	}
	defer conn.Close()
	client := gpay.NewPayManagerClient(conn)

	// gRPCマイクロサービスの支払処理関数を叩く
	gres, err := client.Charge(context.Background(), greq)
	if err != nil {
		c.JSON(http.StatusForbidden, err)
		return
	}
	c.JSON(http.StatusOK, gres)
}
