package engine

import (
	"fmt"

	"github.com/jeromelesaux/fizzbuzz/model"
)

func DoFizzbuzz(p *model.Parameters) []string {
	result := make([]string, 0)
	var i int64
	for i = 1; i <= p.Limit; i++ {
		if i%p.Int1 == 0 && i%p.Int2 == 0 {
			result = append(result, p.Str1+p.Str2)
		} else {
			if i%p.Int1 == 0 {
				result = append(result, p.Str1)
			} else {
				if i%p.Int2 == 0 {
					result = append(result, p.Str2)
				} else {
					v := fmt.Sprintf("%d", i)
					result = append(result, v)
				}
			}
		}
	}
	return result
}
