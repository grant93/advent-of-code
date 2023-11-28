#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>

typedef struct Node {
	int64_t value;
	int idx;

	struct Node *prev;
	struct Node *next;
} Node_t;

Node_t *head = NULL;

void print_list() {
	Node_t *curr = head;
	printf("--------------------------\n");
	while(1) {
		printf("Original Index: %d Value: %ld \n", curr->idx, curr->value);
		if (curr->next == head) {
			break;
		}
		curr = curr->next;
	}
	printf("--------------------------\n");
}


/* could maintain a hashmap with idx as key and pointer to entry as value to
 * speed it up */
Node_t *find_element(int index) {
	Node_t *curr = head;
	while(1) {
		if (curr->idx == index) {
			return curr;
		}
		curr = curr->next;
	}
}

Node_t *find_value(int value) {
	Node_t *curr = head;
	while(1) {
		if (curr->value == value) {
			return curr;
		}
		curr = curr->next;
	}
}

void move_element(Node_t *element) {
	Node_t *curr = element;
	/* detatch from list*/
	if (element->value == 0) {
		return;
	}
	element->next->prev = element->prev;
	element->prev->next = element->next;

	if (element->value > 0) {
		for(int i = 0; i < element->value; i++) {
			curr = curr->next;
		}
		if (curr->next == head) {
			head = element;
		} else if (element == head) {
			head = head->next;
		}
		element->next = curr->next;
		element->prev = curr;
		element->prev->next = element;
		element->next->prev = element;
	} else if (element->value < 0) {
		for(int i = 0; i > element->value; i--) {
			curr = curr->prev;
		}
		if (curr->next == head) {
			head = element;
		} else if (element == head) {
			head = head->next;
		}
		element->next = curr;
		element->prev = curr->prev;
		element->prev->next = element;
		element->next->prev = element;
	}
}

size_t parse_input() {
	Node_t *_head = NULL;
	Node_t *prev = NULL;
	Node_t *curr = NULL;

	char *line = NULL;
	size_t len = 0;
	int idx = 0;

	while(getline(&line, &len, stdin) > 0) {
		prev = curr;
		curr = malloc(sizeof(Node_t));
		if (!_head) {
			_head = curr;
		}

		curr->value = strtol(line, NULL, 10);
		curr->idx = idx;
		free(line);
		line = NULL;
		len = 0;

		if (prev) {
			curr->prev = prev;
			prev->next = curr;
		}
		idx++;
	}

	/* make it circular */
	if (_head && curr) {
		_head->prev = curr;
		curr->next = _head;
	}

	head = _head;
	return idx;
}

/* sort the list back to the original index */
void reinitialise(int len) {
	Node_t *_head = NULL;
	Node_t *curr = NULL;
	Node_t *prev = NULL;

	for(int i=0; i<len; i++) {
		printf("init: %d len: %d\n", i, len);
		if (!_head) {
			_head = find_element(i); 
			curr = _head;
		} else {
			prev = curr;
			curr = find_element(i);
		}

		curr->value *= 811589153;
		if (prev) {
			curr->prev = prev;
			prev->next = curr;
		}
	}

	_head->prev = curr;
	curr->next = _head;
	head = _head;
	
}

Node_t *traverse(Node_t *start, int distance) {
	Node_t *curr = start;
	for(int i=0; i< distance; i++) {
		curr = curr->next;
	}
	return curr;
}

void decrypt(size_t len) {
	for(int i=0; i < len; i++) {
		Node_t *item = find_element(i);
		move_element(item);
	}
}

void partOne() {
	Node_t *zero = find_value(0);
	int total = 0;
	for(int i=1; i<=3; i++) {
		zero = traverse(zero, 1000);
		total += zero->value;
	}
	printf("Part One: %d\n", total);
}

void partTwo(int len) {
	for(int j=0; j<10; j++) {
		printf("....%d\n", j);
		decrypt(len);
	}
	Node_t *zero = find_value(0);
	int total = 0;
	for(int i=1; i<=3; i++) {
		zero = traverse(zero, 1000);
		total += zero->value;
	}
	printf("Part Two: %d\n", total);

}

int main() {
	size_t len = parse_input();	
	print_list(head);
	decrypt(len);
	partOne();
	reinitialise(len);
	printf("done");
	partTwo(len);
}
