# Tricks On Go
Share some tricks on Go during my journey on Go. 🚀

## Convert strings 2 slice of Byte (or reverse) without memalloc

![bench_s2b_b2s_nomemallocs](./pic/bench_s2b_b2s_nomemallocs.png)

    see string_to_byteslice_without_memalloc

## Bench test 3 different ways to connect string

![bench_connect_of_string](./pic/bench_connect_of_string.png)

Buffer2 use`buffer.Write()` while Buffer use `buffer.WriteString()`, and in Buffer2, i use []byte(s) to convert string



## Bench test 3 different ways to convert integer to string

![bench_i2s](./pic/bench_i2s.png)

