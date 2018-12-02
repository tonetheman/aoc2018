#include <stdio.h>
#include <stdlib.h>

int main() {

	char * line = NULL;
	size_t len = 0;
	int read;

	printf("running day1b...\n");

	FILE * inf = fopen("data.txt","r");
	if (inf==NULL) {
		return -EXIT_FAILURE;
	}

	int already_found[2000];
	for(int i=0;i<2000;i++) {
		already_found[i]=0;
	}
	already_found[0] = 0;
	int max_already_found = 1;
	int found = -1;

	int freq=0;
	while ((read=getline(&line, &len, inf)!=-1)) {
		if (found==1) {
			printf("DONE\n");
			break;
		}
		int v = atoi(line);
		freq+=v;
		printf("got %d and freq is now %d\n",v,freq);

		// see if freq already found
		for(int i=0;i<max_already_found;i++) {
			printf("\tcmd %d %d\n",already_found[i],freq);
			if (already_found[i]==freq) {
				printf("MATCH!!!!\n");
				found=1;
				break;
			}
		}
		if (found==-1) {
			already_found[max_already_found] = freq;
			max_already_found+=1;
		}
	}
	fclose(inf);
	printf("end result %d\n",freq);

	return 0;

}
