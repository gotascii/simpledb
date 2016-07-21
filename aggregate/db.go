package aggregate

type DB interface {
	Copy() DB
	Append(idx string, val int)
	Remove(idx string, val int)
	Compute(val string) int
}
