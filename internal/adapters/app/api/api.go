package api

import "github.com/Ayomide-Daniel/go-hex/internal/ports"

type Adapter struct {
	db    ports.DBPort
	arith ports.ArithmeticPort
}

func NewAdapter(db ports.DBPort, arith ports.ArithmeticPort) *Adapter {
	return &Adapter{db: db, arith: arith}
}

func (apiA Adapter) GetAddition(a, b int32) (int32, error) {
	ans, err := apiA.arith.Addition(a, b)

	if err != nil {
		return 0, err
	}

	err = apiA.db.Create(ans, "addition")

	if err != nil {
		return 0, err
	}

	return ans, nil
}

func (apiA Adapter) GetSubtraction(a, b int32) (int32, error) {
	ans, err := apiA.arith.Subtraction(a, b)

	if err != nil {
		return 0, err
	}

	err = apiA.db.Create(ans, "subtraction")

	if err != nil {
		return 0, err
	}

	return ans, nil
}

func (apiA Adapter) GetMultiplication(a, b int32) (int32, error) {
	ans, err := apiA.arith.Multiplication(a, b)

	if err != nil {
		return 0, err
	}

	err = apiA.db.Create(ans, "multiplication")

	if err != nil {
		return 0, err
	}

	return ans, nil
}

func (apiA Adapter) GetDivision(a, b int32) (int32, error) {
	ans, err := apiA.arith.Division(a, b)

	if err != nil {
		return 0, err
	}

	err = apiA.db.Create(ans, "division")

	if err != nil {
		return 0, err
	}

	return ans, nil
}
