// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNickName holds the string denoting the nick_name field in the database.
	FieldNickName = "nick_name"
	// EdgeSubscribeStocks holds the string denoting the subscribe_stocks edge name in mutations.
	EdgeSubscribeStocks = "subscribe_stocks"
	// Table holds the table name of the user in the database.
	Table = "user"
	// SubscribeStocksTable is the table that holds the subscribe_stocks relation/edge. The primary key declared below.
	SubscribeStocksTable = "stock_subscribers"
	// SubscribeStocksInverseTable is the table name for the Stock entity.
	// It exists in this package in order to avoid circular dependency with the "stock" package.
	SubscribeStocksInverseTable = "stock"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldNickName,
}

var (
	// SubscribeStocksPrimaryKey and SubscribeStocksColumn2 are the table columns denoting the
	// primary key for the subscribe_stocks relation (M2M).
	SubscribeStocksPrimaryKey = []string{"stock_id", "user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NickNameValidator is a validator for the "nick_name" field. It is called by the builders before save.
	NickNameValidator func(string) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(uint64) error
)