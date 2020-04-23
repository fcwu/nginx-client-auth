1. build *.a
2. run .a

1. build binary
2. use memfd_create to create file like tempfs http://man7.org/linux/man-pages/man2/memfd_create.2.html https://magisterquis.github.io/2018/03/31/in-memory-only-elf-execution.html
3. run binary by fork and exec
