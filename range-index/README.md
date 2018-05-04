# range-index

Problem:
    Previously I want to check if something cannot find in the array, add it to the exising array.
    So I use golang built-in function `range`, and I wants to check the idx equals to the len(array)
    to check the element is not exists in the array. However the code run in incorrect behavior.
    After do some test, I found that if I use built-in `range` function, the lastest idx value equals
    to len(array) - 1, not len(array).

Solve:
    Please use `for idx = 0; idx < len(array); idx++` insteads of `for idx = range array`.
