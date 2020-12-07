package exchange

import "fmt"

type ClientError struct {
	Context string
	E error
}

func (c *ClientError) Error() string {
	return fmt.Sprintf("Error making %s request: %s", c.Context, c.E.Error())
}

type MultipleEventMarketsError struct {
	EventID string
}

func (m *MultipleEventMarketsError) Error() string {
	return fmt.Sprintf("Multiple markets returned for event: %s", m.EventID)
}

type MultipleMarketSelectionError struct {
	EventID string
	SelectionID uint64
}

func (m *MultipleMarketSelectionError) Error() string {
	return fmt.Sprintf("Multiple selections returned for market %s and selection %d", m.EventID, m.SelectionID)
}

type NoEventMarketError struct {}

func (m *NoEventMarketError) Error() string {
	return fmt.Sprintf("No markets returned for event and market")
}

type NoEventError struct {
	Event string
}

func (n *NoEventError) Error() string {
	return fmt.Sprintf("No event returned for: %s", n.Event)
}
