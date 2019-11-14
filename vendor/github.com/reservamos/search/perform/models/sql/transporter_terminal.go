package sql

import "time"

// TransporterTerminal holds the information of a terminal for a specific Transporter
// dl Tag used for dataloader population
type TransporterTerminal struct {
	ID            int       `gorm:"primary_key" dl:"id"`
	TerminalID    int       `dl:"terminal_id"`
	TransporterID int       `dl:"transporter_id"`
	Code          string    `dl:"code"`
	CreatedAt     time.Time `dl:"created_at"`
	UpdatedAt     time.Time `dl:"updated_at"`

	Terminal    Terminal
	Transporter Transporter
}
