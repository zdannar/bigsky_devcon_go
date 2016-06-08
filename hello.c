// R1 OMIT
#include<stdio.h>
#include<stdbool.h>
#include<string.h>

bool statement(int n, char* buffPtr)
{
    switch(n) {
       case 1 :
          strcpy(buffPtr, "loneliest");
          break;
       case 2 :
          strcpy(buffPtr, "cozyiest");
          break;
       case 3 :
          strcpy(buffPtr, "kinkiest");
          break;
       default :
          printf("What was the number? %i", n);
          return false;
    }
    return true;
}
// R1E OMIT
// R2 OMIT
int main()
{
    int friends = 1;
    char* stringPtr;

    if (!statement(friends, stringPtr))
    {
        printf("Wrong number");
        return 1;
    }

    printf("\n%i is the %s number.\n", friends, stringPtr);
    return 0;
}
// R2E OMIT
