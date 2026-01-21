// supabase/client.go
package supabase

import (
	"os"

	"github.com/supabase-community/supabase-go"
)

func NewClient() (*supabase.Client, error) {
	return supabase.NewClient(
		os.Getenv("SUPABASE_URL"),
		os.Getenv("SUPABASE_ANON_KEY"),
		nil,
	)
}
