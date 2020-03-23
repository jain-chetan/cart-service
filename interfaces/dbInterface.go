package interfaces

import (
	"github.com/jain-chetan/cart-service/model"
)

//DBClient object
var DBClient DBInteractions

//-----------------------------------------------------------------------------------------------------------------------//

//DBInteractions interface contains database methods
type DBInteractions interface {
	DBConnect(config model.DBConfig) error
}
