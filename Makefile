NAME=n-puzzle
SRC=$(shell find . -name "*.go")

all: $(NAME)

$(NAME): $(SRC)
	go build

re: fclean all

fclean:
	rm -f $(NAME)

.PHONY: all re fclean