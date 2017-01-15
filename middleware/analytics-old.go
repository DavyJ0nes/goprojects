package middleware

import (
	"fmt"
	"net/http"
)

// Add is a variadic function that adds up numbers
func Add(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// Chain holds the value of the Sum
type Chain struct {
	Sum int
}

// AddNext is a chainable sum function. Similar to a recursive function
func (c *Chain) AddNext(num int) *Chain {
	c.Sum += num
	return c
}

// ReturnSum tidies up output of Chain object
func (c *Chain) ReturnSum() int {
	return c.Sum
}

// Next runs the next function in the middleware chain
func Next(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("BeforeRequst", r)
		next.ServeHTTP(w, r)
		fmt.Println("AfterRequest", r)
	})
}
