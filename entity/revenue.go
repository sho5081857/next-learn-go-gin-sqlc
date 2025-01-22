package entity

type Revenue struct {
	month   string
	revenue int32
}

func (r *Revenue) Month() string {
	return r.month
}

func (r *Revenue) Revenue() int32 {
	return r.revenue
}

func ReconstructRevenue(
	month string,
	revenue int32,
) (*Revenue, error) {
	return newRevenue(
		month,
		revenue,
	)
}

func newRevenue(
	month string,
	revenue int32,
) (*Revenue, error) {
	return &Revenue{
		month:   month,
		revenue: revenue,
	}, nil
}
