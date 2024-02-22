module main

go 1.21.3

require (
	github.com/blendle/zapdriver v0.0.0-00010101000000-000000000000
	github.com/go-chi/chi v1.5.5
	github.com/justinas/alice v1.2.0
	github.com/lib/pq v1.10.9
	go.uber.org/fx v1.20.1
	go.uber.org/zap v1.27.0
)

require (
	go.uber.org/dig v1.17.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
)

replace github.com/blendle/zapdriver => github.com/Shivam010/zapdriver v1.3.2-0.20201201095836-e7d0e8f9ced0
