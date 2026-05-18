package core_domain

import (
	"errors"
	"fmt"
	"regexp"
)

type UrlDomain struct {
	Id    int
	Alias string
	Url   string
}

var (
    ErrNotFound      = errors.New("not found")
    ErrAlreadyExists = errors.New("already exists")
    ErrInvalidInput  = errors.New("invalid input")
    ErrUnauthorized  = errors.New("unauthorized")
    ErrForbidden     = errors.New("forbidden")

    ErrInternal      = errors.New("internal error")
    ErrDatabase      = errors.New("database error")
    ErrTimeout       = errors.New("timeout")
)

func NewUrl(id int, alias string, url string) UrlDomain {
	return UrlDomain{
		Id:    id,
		Alias: alias,
		Url:   url,
	}
}

func (d *UrlDomain) Validate() error {
	aliasLen := len([]rune(d.Alias))
	if aliasLen < 3 || aliasLen > 100 {
		return fmt.Errorf("invalid alias len: %d", aliasLen)
	}

	urlen := len([]rune(d.Url)) 

	if urlen < 3 || urlen > 100 {
		return fmt.Errorf("invalid url len: %d", urlen)
	}
	
	re := regexp.MustCompile(`^https?:\/\/(?:www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,63}\b(?:[-a-zA-Z0-9()@:%_\+.~#?&\/=]*)$`)

	if !re.MatchString(d.Url) {
		return fmt.Errorf("invalid `url` format: %s", d.Url)
	}

	return nil
}

type UpdateUrlDomain struct {
	Id int
	Url *string
	Alias *string
}

func NewUpdateDomain(
	Id int,
	Url *string,
	Alias *string,
	) UpdateUrlDomain {
		return UpdateUrlDomain{
			Id: Id,
			Url: Url,
			Alias: Alias,
		}
}