# infinite-loop-select-timeout

`Select` grammar:
1.Every case `must be` a channel.
2.If `multiple channel` could execute, it will radomly choose a case to execute; otherwise, execute default case.
3.If there is no default case, select will be blocked until any case(s) could be execute.

Problem:
    A infinite loop and select with time.After to stop infinite loop, the stop time is not correctly.
    Since select is `randomly choose` a case to execute if there are multiple cases could execute.

Solve:
    Use multiple select to fairly choose a case to execute.
