package module

import (
	"database/sql"
	"monitorservice/internal/api"
)

type Module struct {
	Tequila api.ITequila
}

func NewModule(t api.ITequila) *Module {
	return &Module{
		Tequila: t,
	}
}

func (m *Module) InitModule(logger api.ILogger, db *sql.DB, tequila api.ITequila) error {

	// if err := tequila.RegisterRPC("test", Testfunc); err != nil {
	// 	return err
	// }

	/* TODO:
	1. login
	2. add channel
	3. speak to channel
	4. speak to person
	*/

	return nil
}
