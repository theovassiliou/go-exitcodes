# exitcodes

The package exitcodes defines some typical exit codes. As there is no standard we stick to the best practice, which is

    exitcode = 0 ; SUCCESS
    exitcoe != 0 ; FAILURE

As this is not very helpful and you always make up your mind, what to pass to `os.exit()` this packages captures some typical situations.

## References

- [Bash Exit Status](https://www.gnu.org/software/bash/manual/html_node/Exit-Status.html)
- [Discussion on Stackexchange](https://unix.stackexchange.com/questions/110348/how-do-i-get-the-list-of-exit-codes-and-or-return-codes-and-meaning-for-a-comm)
