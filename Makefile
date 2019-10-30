NAME=n-puzzle

all: $(NAME)

$(NAME):
	go build

re: fclean all

fclean:
	rm -f $(NAME)

.PHONY: all re fclean