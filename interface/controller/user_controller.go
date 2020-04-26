package controller

import (
	"fmt"
	"math/big"
	"net/http"

	"github.com/jinzhu/gorm"
)

type userController struct {
	db *gorm.DB
}

type dummyData struct {
	Name  string   `json:"name"`
	Count int      `json:"count"`
	Fib   *big.Int `json:"fibonacci"`
}

type UserController interface {
	Get(c Context) error
}

func (uc *userController) Get(c Context) error {
	u := &dummyData{
		Name:  "Test",
		Count: 10,
		Fib:   fib(0),
	}

	// uc.db.Exec(`INSERT INTO test (str, raw_data) VALUES ("x", "b")`)

	// Scan
	type Result struct {
		Str     string
		RawData string
	}

	var result Result
	uc.db.Raw("SELECT * FROM test WHERE str = ?", "x").Scan(&result)
	// res := uc.db.Exec(`SELECT * FROM test WHERE str=?`, "x").Row()

	fmt.Printf("%+v", result)

	return c.JSON(http.StatusOK, u)
}

func NewUserController(db *gorm.DB) UserController {

	return &userController{
		db: db,
	}
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
