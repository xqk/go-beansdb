package memcache

import (
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {

	t.Run("ping", func(t *testing.T) {
		sch := NewModScheduler([]string{"192.168.0.17:7901"}, "md5")
		client := NewRClient(sch, 1, 1, 1)
		item, targets, err := client.Get("LB_44906648")
		if err != nil {
			t.Error(err)
			return
		}

		if item != nil {
			fmt.Println(string(item.Body))
		}

		if len(targets) > 0 {
			fmt.Println(targets)
		}
	})
}
