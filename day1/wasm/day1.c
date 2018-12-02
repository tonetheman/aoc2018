#include <stdio.h>
#include <stdlib.h>

int main() {

	char * line = NULL;
	size_t len = 0;
	int read;

	printf("running day1...\n");

	FILE * inf = fopen("data.txt","r");
	if (inf==NULL) {
		return -EXIT_FAILURE;
	}

	int freq=0;
	while ((read=getline(&line, &len, inf)!=-1)) {
		int v = atoi(line);
		printf("got %d\n",v);
		freq+=v;
	}
	fclose(inf);
	printf("end result %d\n",freq);

	return 0;

}
