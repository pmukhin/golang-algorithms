# QuickSort

QuickSort is an algorithm with floating complexity which depends on what input array consists of. Wikipedia says that qsort is an enhanced version of BubbleSort.

Here's algorithm structure:

1. If array length is less or equals to 1, we return input array.
2. If not, we take pivot value - an element of array which could be got by two ways:
 - By random int (which we get from array like `input[rand.Intn(len(input) - 1)]`
 - By middle element (like `input[len(input)/2]`
3. Now we have to create three new arrays - left (for elements less than pivot), middle (for those equal to pivot) and right (for values that greater than pivot)
4. Let's iterate through input array. As described above we put elements less than pivot to left, equal to pivot to middle and so on
5. We recursively sort left and right arrays by using MergeSort (no need to sort middle, because all values in that are equal)
6. We attach all elements from middle array to left
7. We attach all elements from right array to left
8. Return left

