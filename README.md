# zmachine
Z-machine interpreter in Go calls ChatGPT

- wine infodump.exe -f ../zmachine/905.z5 | less
-  ./frotz -d -p ../905.z5
- go build
- ./zmachine -file 905.z5 | grep "ip\|opValues\|Store\|Tok\|Input" > out.txt
- diff -u out.txt frotz/trace.txt  | colordiff | less -R
