
.PHONY: run
run: check
	go run ./cmd/run --day $(AOC_DAY)

.PHONY: check
check:
	@if [ -z "$(AOC_DAY)" ]; then \
		echo "must set AOC_DAY to use make scripts"; \
		exit 1; \
	fi


.PHONY: check-dirty
check-dirty:
	@if ! git diff-index --quiet HEAD --; then \
		echo "git working tree is dirty"; \
		echo "must commit or discard to use this script"; \
		exit 1; \
	fi

.PHONY: init
init: check check-dirty
	go run ./cmd/init --day $(AOC_DAY)
	git add .
	git commit -m "Setup for day $(AOC_DAY)."
	go run ./cmd/get_puzzle --day $(AOC_DAY) --part 1
	go run ./cmd/get_input --day $(AOC_DAY)
	git add .
	git commit -m "Retrieved puzzle for day $(AOC_DAY), part 1."

.PHONY: init2
init2: check check-dirty
	go run ./cmd/get_puzzle --day $(AOC_DAY) --part 2
	git add .
	git commit -m "Retrieved puzzle for day $(AOC_DAY), part 2."

.PHONY: save1
save1: check
	git add .
	git commit -m "Save solution for day $(AOC_DAY), part 1."

.PHONY: save2
save2: check
	git add .
	git commit -m "Save solution for day $(AOC_DAY), part 1."

.PHONY: submit
submit1: check
	go run ./cmd/submit --day $(AOC_DAY) --part 1

.PHONY: submit
submit2: check
	go run ./cmd/submit --day $(AOC_DAY) --part 2
