
.PHONY: run
run: check
	go run ./cmd/run --day $(AOC_DAY)

.PHONY: check
check:
	@if [ -z "$(AOC_DAY)" ]; then \
		echo "must set AOC_DAY to use make scripts"; \
		exit 1; \
	fi


.PHONY: init
init: check
	go run ./cmd/init --day $(AOC_DAY)
	go run ./cmd/get_puzzle --day $(AOC_DAY) --part 1
	go run ./cmd/get_input --day $(AOC_DAY)

.PHONY: init2
init2: check
	go run ./cmd/get_puzzle --day $(AOC_DAY) --part 2

.PHONY: submit
submit1: check
	go run ./cmd/submit --day $(AOC_DAY) --part 1

.PHONY: submit
submit2: check
	go run ./cmd/submit --day $(AOC_DAY) --part 2
