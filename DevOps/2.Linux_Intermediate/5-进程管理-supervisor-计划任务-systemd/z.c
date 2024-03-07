#include	<sys/types.h>
#include	<sys/wait.h>
#include	<errno.h>
#include	<unistd.h>
#include	<stdlib.h>
#include	<stdio.h>
#include	<string.h>
int	main(int argc, char *argv[])
{
pid_t pid;
pid = fork();
if (pid==0) {
	int iPid =(int)getpid();
	fprintf(stderr,"子进程,%d\n",iPid);
	sleep(1);
	fprintf(stderr,	"Child	exits\n");
	return	EXIT_SUCCESS;
}
int iPid = (int)getpid();
	fprintf(stderr,"父进程,%d\n",iPid);
	fprintf(stderr,	"sleep....\n");
	sleep(360);
	fprintf(stderr,	"parent	exits\n");
	return	EXIT_SUCCESS;
}
