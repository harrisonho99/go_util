#include <stdlib.h>
#include <stdio.h>

void *wrap_malloc(size_t count, size_t size)
{
    return calloc(count, size);
}
