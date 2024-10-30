package valueobjects

type (
	VoterLimit struct {
		VoterLimit int
	}
)

func NewVoterLimit(voterlimit int) (*VoterLimit, error) {
	return &VoterLimit{
		VoterLimit: voterlimit,
	}, nil
}
