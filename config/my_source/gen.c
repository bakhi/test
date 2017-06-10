#include <stdio.h>
#include <stdint.h>
#include <time.h>

int main(int argc, char* argv[]){
	FILE* wfp;
	int32_t i;
	int64_t size;

	wfp = fopen("input.json", "w");

	if (argc <2) {
		return -1;
	} else {
		srand(time(NULL));
		for(i = 0; i < atoi(argv[1]); i++){
			fprintf(wfp,"{\"key\" : %d}\n", rand()%1000);
		}
	}
	fclose(wfp);
}
