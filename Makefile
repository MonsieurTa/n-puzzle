NAME=n-puzzle

$(NAME): all

all:
	go build

re: fclean all

fclean:
	rm -f $(NAME)

.PHONY: all re fclean