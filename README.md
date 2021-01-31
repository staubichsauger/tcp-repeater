# tcp repeater
Simple program, which relays an incoming tcp connection on port 9990 to an outgoing tcp connection on port 9991.
It first waits for the receiver to connect on port 9991 befor it lets the tranceiver connect to port 9990.

## Use case real time audio relay with ffmpeg
I use it to stream real time audio (1-2 sec) from one PC to another PC via a jump server using ffmpeg.

First, on the receiver side:
`ffplay tcp://<server-ip>:9991`

Afterwards, on the transmitter side:
`ffmpeg -f alsa -ac 2 -i default -c:a aac -b:a 128k -f matroska tcp://<server-ip>:9990`

The reason the program is set up this way, is that there is an endless loop script running on the transmittig PC trying to send audio.
Once a receiver connects to the tcp-repeater, audio will start streaming.
When the receiver closes the connection (e.g. ffplay is closed), the program exits.
I run it as a systemd service, to ensure a receiver can connect any time to receive audio.

If I let the transmitter connect first, there is a bigger latency.
Not 100% sure why, but there probably is some buffer on the receiving connection, which leads to the buffer being sent to the receiver from the start, which will contain old audio.
