# extsort
Basic external sort implementation on golang (based on `container/heap`).

`generator` - a tool to generate random strings for sort input.

    generator -n lines -m max_line_len -s random_seed

Repeated invocations with the same parameters produce identical outputs.
Specify different seeds if this is not desired.

`sort` - a tool for sorting input lines. Suitable for sorting plain ascii strings.

    sort -n lines_per_bunch

Each bunch of input lines is sorted and written to a separate temporary file.
The temporary files are then read line by line into a heap for final sorting.
Every bunch uses one file descriptor. Specify larger bunch size for huge inputs
to avoid "too many open files" error (default is 5000 lines per bunch).

Usage:

    cd generator
    go build
    cd ../sort
    go build

    # Feed the same inputs into stadard sort and our implementation.
    # The diff should be empty.
    diff <(../generator/generator -n 149876|sort) <(../generator/generator -n 149876|./sort)

    # huge input
    time ../generator/generator -n 2000000 | ./sort -n 20000 | tail
