package core_domain

type StatDomain struct {
	Id        int
	UrlId     int
	Ip        string
	UserAgent string
	Device    string
}

func NewStatDomain(id int, UrlId int, Ip string, device string, userAgent string) StatDomain {
	return StatDomain{
		Id:        id,
		UrlId:     UrlId,
		Ip:        Ip,
		Device:    device,
		UserAgent: userAgent,
	}
}