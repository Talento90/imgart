package gomage

type Profile struct {
	Id          string
	Effects     []Effect
	FallbackUri string
	Format      string
	Compression int
}
