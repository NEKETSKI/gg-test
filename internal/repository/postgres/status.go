package postgres

import (
	"context"
	"fmt"
	"time"
)

// StoreStatusCode stores the received status code in the database. The request URL and timestamp will also be saved
// along with the code.
func (p *psql) StoreStatusCode(requestUrl string, code int) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()
	_, err := p.db.Exec(ctx, "INSERT INTO status_codes(request, code, request_time) VALUES($1,$2,$3)",
		requestUrl, code, time.Now())
	if err != nil {
		return fmt.Errorf("statement execution failed: %w", err)
	}
	return nil
}
