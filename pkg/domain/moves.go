package domain

type Move interface {
	Target() Square
	Via() []Square
	From() Square
}
