package users

import (
	"database/sql/driver"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Password []byte

func (p Password) IsSamePassword(other string) bool {
	return nil == bcrypt.CompareHashAndPassword([]byte(p), []byte(other))
}

func (p *Password) Scan(src any) error {
	v, ok := src.(string)
	if !ok {
		return fmt.Errorf("unexpected type: %T, expect %T", src, Password{})
	}
	*p = Password(v)
	return nil
}

func (p *Password) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}

	return []byte(*p), nil
}
