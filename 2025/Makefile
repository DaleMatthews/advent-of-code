init:
	for folder in day*; do \
		touch $$folder/input1.txt; \
		touch $$folder/example1.txt; \
	done

day%:
	go test ./day$* -v

bench%:
	go test -bench=. ./day$* -benchmem
