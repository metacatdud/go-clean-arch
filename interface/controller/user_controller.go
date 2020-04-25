package controller

import (
	"math/big"
	"net/http"
)

type userController struct {
	Name  string   `json:"name"`
	Count int      `json:"count"`
	Fib   *big.Int `json:"fibonacci"`
}

type UserController interface {
	Get(c Context) error
}

func (uc *userController) Get(c Context) error {
	u := &userController{
		Name:  "Test",
		Count: 10,
		Fib:   fib(40),
	}
	return c.JSON(http.StatusOK, u)
}

func NewUserController() UserController {

	return &userController{}
}

func fib(n int) *big.Int {
	fn := make(map[int]*big.Int)

	for i := 0; i <= n; i++ {
		var f = big.NewInt(0)
		if i <= 2 {
			f.SetUint64(1)
		} else {
			f = f.Add(fn[i-1], fn[i-2])
		}
		fn[i] = f
	}
	return fn[n]
}
