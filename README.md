# bolid-eth-client

UDP client
----------

Send commands and recalculate CRC8 (with replace the last byte).

Example send hex 010600150203a7 to RS-485:
```
Enter text: 1007010b1001060015020300
```
last 00 will be replace to CRC8 and send to c2000-eth and next translate to RS-485.
Below example blink relay # on device addr #1

Example enable relay #2 on the Device c2000-sp1 address #127
```
go run main.go -addr=1 -relay=2 -on=1 -cmd=21
```

Example disable relay #2 on the Device 127
```
go run main.go -addr=1 -relay=2 -on=0 -cmd=21
```

may bee blink "-on=3"
