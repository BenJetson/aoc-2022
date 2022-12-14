# aoc-2022

![Go Report Card][go-report] [![Build/Tests][build-badge]][build]

This repository contains my solutions to the [2022 Advent of Code][aoc-2022],
written in [Go][golang].

![Advent of Code with Santa Hat](https://user-images.githubusercontent.com/10427974/100974572-7db6a900-350a-11eb-9c80-18635d97861f.png)

## ⚠️ SPOILER ALERT ⚠️

If you have not yet completed the **2022 Advent of Code**, but plan to do so,
you might want to reconsider viewing the contents of this repository.

This repository contains full solutions to all of the puzzles, which might ruin
the fun for you.

You have been warned!

## Structure

To see my solution code, go to the [days directory](./days).

## Commands

I have a few Go commands that make participating in AOC easier.

When running commands, the **working directory should be the repository root**.

### Initialize a Day

```sh
go run ./cmd/init --day X
```

This program initializes for a given advent calendar day by:

- creating a new directory in [days](./days);
- creating a `README.md` file in that day's directory for the puzzle text;
- creating a `solve.go` file in that day's directory for the solution code; and
- modifying [`days/days.go`](./days/days.go) to link that day's solver

| Flag    | Explanation                                       |
| ------- | ------------------------------------------------- |
| `--day` | The advent calendar day, an integer 1 through 25. |

### Get Puzzle

```sh
go run ./cmd/get_puzzle --day X --part Y
```

This program downloads the puzzle body text from the AOC server, translates the
HTML to markdown, and then saves it in that day's `README.md` file.

| Flag     | Explanation                                        |
| -------- | -------------------------------------------------- |
| `--day`  | The advent calendar day, an integer 1 through 25.  |
| `--part` | The puzzle part number, an integer, either 1 or 2. |

### Get Input

```sh
go run ./cmd/get_input --day X
```

This program downloads the puzzle input for my account from the AOC server and
saves it in that day's `input.txt` file.

| Flag    | Explanation                                       |
| ------- | ------------------------------------------------- |
| `--day` | The advent calendar day, an integer 1 through 25. |

### Run Solver

```sh
go run ./cmd/run --day X
```

This program runs the my solver code for a given advent calendar day.

By default, it uses my account's input from `input.txt` for that day. If you
would like to try your own input, you may pass a filename to the input flag.

| Flag      | Explanation                                       |
| --------- | ------------------------------------------------- |
| `--day`   | The advent calendar day, an integer 1 through 25. |
| `--input` | Optional, path to a different input file to use.  |

### Submit Answer

```sh
go run ./cmd/submit_answer --day 1 --part 1
```

This program submits the answer generated by my solver code for a given advent
calendar day and puzzle part to the AOC servers.

| Flag    | Explanation                                       |
| ------- | ------------------------------------------------- |
| `--day` | The advent calendar day, an integer 1 through 25. |

## Authentication

The commands that interact with the AOC servers require authentication.

AOC uses a session cookie to authenticate requests. After logging in, you can
find the value of this cookie using the developer tools in your browser.

The client expects to find the **value only** of the cookie in `.aoc-session` at
the repository root.

## Network Request Advisory

**⚠️ WARNING**: if you attempt to use the commands from this repository that
make network requests to the AOC servers, please **be respectful**.

The AOC maintainers ask that users do not make frequent automated requests.

[go-report]: https://goreportcard.com/badge/github.com/BenJetson/aoc-2022
[build]: https://github.com/BenJetson/aoc-2022/actions/workflows/go.yml
[build-badge]:
  https://github.com/BenJetson/aoc-2022/actions/workflows/go.yml/badge.svg
[aoc-2022]: https://adventofcode.com/2022
[golang]: https://go.dev
