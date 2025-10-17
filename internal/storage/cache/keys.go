package cache

import "fmt"

func (s *Store) KeyProduct(ID int32) string {
	return s.prefix + ":product:" + fmt.Sprint(ID)
}
func (s *Store) KeyAllProducts() string {
	return s.prefix + ":products:all"
}
