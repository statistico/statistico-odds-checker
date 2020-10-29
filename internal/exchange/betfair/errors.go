package betfair

import "fmt"

type clientError struct {
	context string
	e error
}

func (c *clientError) Error() string {
	return fmt.Sprintf("Error making %s request: %s", c.context, c.e.Error())
}

type multipleEventMarketsError struct {
	eventID string
}

func (m *multipleEventMarketsError) Error() string {
	return fmt.Sprintf("Multiple markets returned for event: %s", m.eventID)
}

type multipleMarketSelectionError struct {
	eventID string
	selectionID uint64
}

func (m *multipleMarketSelectionError) Error() string {
	return fmt.Sprintf("Multiple selections returned for market %s and selection %d", m.eventID, m.selectionID)
}

type noEventError struct {
	event string
}

func (n *noEventError) Error() string {
	return fmt.Sprintf("No event returned for: %s", n.event)
}

type orderFailure struct {
	marketID string
	selectionID uint64
	errorCode string
	reason string
}

func (o *orderFailure) Error() string {
	return fmt.Sprintf(
		"Error placing order for market %s and selection %d. Code: %s, Reason %s",
		o.marketID,
		o.selectionID,
		o.errorCode,
		o.reason,
	)
}
