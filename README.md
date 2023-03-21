# preftime

![Build Status](https://github.com/mokiat/preftime/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/mokiat/preftime)](https://goreportcard.com/report/github.com/mokiat/preftime)

Preftime is a tool that allows you to prefix the output of another program with timestamp information.

Ideally, the program you are running should have the capability to log timestamps on its own. Alternatively, you could check whether you terminal supports such a feature. However, I have been in situations where neither the former nor the latter did a good job for me. The tool I was using did not print timestamps and though the terminal I was using did support adding timestamps, they were not selectable with the rest of the output. I ended up writing these few lines of code and decided to share it for anyone who may need it.

If you have a Go development environment set up, you can use the following command to acquire the tool.

```bash
go get github.com/mokiat/preftime
```

Otherwise, there are pre-built versions of the tool available in the [Releases](https://github.com/mokiat/preftime/releases) section of this GitHub project.

You would use the tool as follows.

```bash
command 2>&1 | preftime
```

The `2>&1` part is optional. It just assures that both stdout and stderr go through `preftime`.

As an alternative, you could do the following.

```bash
command > >(preftime) 2> >(preftime)
```

This redirects both stdout and stderr to preftime. The `>(preftime)` part makes `preftime` behave like a writable file.

Following the same logic, you could have your bash script redirect it's output to preftime.

```bash
#!/bin/bash

set -e

exec > >(preftime)
exec 2> >(preftime)

command1
command2
```

For example, here is a real-life simple usage scenario.

```bash
$ host example.org | preftime
[2017-06-25 15:42:00.664] example.org has address 93.184.216.34
[2017-06-25 15:42:00.665] example.org has IPv6 address 2606:2800:220:1:248:1893:25c8:1946
```