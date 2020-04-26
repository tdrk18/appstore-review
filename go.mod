module main

go 1.14

replace (
	github.com/tdrk18/appstore-review/slack => ./slack
	github.com/tdrk18/appstore-review/storeReview => ./storeReview
)

require (
	github.com/tdrk18/appstore-review/slack v0.0.0
	github.com/tdrk18/appstore-review/storeReview v0.0.0
)
