run: 
	@echo "Day: "; \
	read day; \
	mkdir -p day$$day/part1 day$$day/part2; \
	touch day$$day/part1/main.go day$$day/part2/main.go day$$day/input.txt day$$day/solve.go ;\
	echo -e "package part1\n\nfunc Solve(lines []string) {}" >> day$$day/part1/main.go; \
	echo -e "package part2\n\nfunc Solve(lines []string) {}" >> day$$day/part2/main.go; \
	echo -e "package day$$day\n\nfunc Solve() {\nlines := utils.ReadLines(\"./day$$day/input.txt\")\n}" >> day$$day/solve.go; 
	
