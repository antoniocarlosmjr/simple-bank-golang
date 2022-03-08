package entities

type Account struct {
	User            User
	Agency, Account int
	Type            string
	Balance         float64
}
