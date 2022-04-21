#include <stdlib.h>
#include <stdio.h>

void _safe_free(void **pointer)
{
    if (*pointer == NULL)
        return;
    free(*pointer);
    *pointer = NULL;
}

#define SAFE_FREE(p) _safe_free((void **)p)

void s_free(void **p)
{
    SAFE_FREE(p);
}