package commons

import(
 "errors"

)

var (
	ErrNoItems = errors.New("Item must have at least one item")
)