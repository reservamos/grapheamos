package sql

import "time"

// LineRating line rating aggregate
type LineRating struct {
	LineID      int     `dl:"line_id"`
	General     float64 `dl:"general"`
	Punctuality float64 `dl:"punctuality"`
	Staff       float64 `dl:"staff"`
	Bus         float64 `dl:"bus"`
	Count       int     `dl:"count"`
}

type Rating struct {
	ID                 int       `dl:"id"`
	LineID             int       `dl:"line_id"`
	General            float64   `dl:"general"`
	Punctuality        float64   `dl:"punctuality"`
	Staff              float64   `dl:"staff"`
	Bus                float64   `dl:"bus"`
	Count              int       `dl:"count"`
	Comments           string    `dl:"comments"`
	CreatedAt          time.Time `dl:"created_at"`
	PurchaserFirstName string    `dl:"purchaser_first_name"`
	PurchaserLastName  string    `dl:"purchaser_last_name`
}

// TransporterRating transporter rating aggregate
type TransporterRating struct {
	TransporterID int     `dl:"transporter_id"`
	General       float64 `dl:"general"`
	Punctuality   float64 `dl:"punctuality"`
	Staff         float64 `dl:"staff"`
	Bus           float64 `dl:"bus"`
	Count         int     `dl:"count"`
}

func (r Rating) CreatedAtRating() time.Time {
	return r.CreatedAt
}

func (r LineRating) GeneralRating() float64 {
	return r.General
}

func (r Rating) GeneralRating() float64 {
	return r.General
}

func (r Rating) Comment() string {
	return r.Comments
}

func (r LineRating) PunctualityRating() float64 {
	return r.Punctuality
}

func (r Rating) PunctualityRating() float64 {
	return r.Punctuality
}

func (r LineRating) BusRating() float64 {
	return r.Bus
}

func (r Rating) BusRating() float64 {
	return r.Bus
}

func (r LineRating) StaffRating() float64 {
	return r.Staff
}

func (r Rating) StaffRating() float64 {
	return r.Staff
}

func (r LineRating) RatingCount() int {
	return r.Count
}

func (r TransporterRating) GeneralRating() float64 {
	return r.General
}

func (r TransporterRating) PunctualityRating() float64 {
	return r.Punctuality
}

func (r Rating) IDRating() int {
	return r.ID
}

func (r Rating) LineIDRating() int {
	return r.LineID
}

func (r TransporterRating) BusRating() float64 {
	return r.Bus
}

func (r TransporterRating) StaffRating() float64 {
	return r.Staff
}

func (r TransporterRating) RatingCount() int {
	return r.Count
}
