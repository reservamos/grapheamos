package sql

// Agent entity on to which generate requests
type Agent struct {
	ID   int    `gorm:"primary_key" dl:"id"`
	Abbr string `dl:"abbr"`
	Name string `dl:"name"`
}
