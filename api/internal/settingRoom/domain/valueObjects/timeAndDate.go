package valueobjects

type (
	TimeAndDate struct {
		Time int
		Date int
	}
)

// se puede importar un paquete time
func NewTimeAndDate(time, date int) (*TimeAndDate, error) {

	return &TimeAndDate{
		Time: time,
		Date: date,
	}, nil
}
