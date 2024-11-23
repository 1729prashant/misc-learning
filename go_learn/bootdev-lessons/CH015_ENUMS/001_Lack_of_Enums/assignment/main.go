package main
import "fmt"

func (a *analytics) handleEmailBounce(em email) error {
	updateReturn := em.recipient.updateStatus(em.status)
	if updateReturn != nil {
		return fmt.Errorf("error updating user status: %w", updateReturn)
	}
	
	trackReturn := a.track(em.status)
	if trackReturn != nil {
				return fmt.Errorf("error tracking user bounce: %w", updateReturn)
	}
	
	return nil
}

