fritzer audio processing pipeline
=================================

This package implements an audio processing pipeline known as Fritzer.

Samples
-------

Fritzer represents samples as fritzer.Sample, which contains both a left
and right channel and stores values as float64.

Buffers
-------

Fritzer uses a struct called fritzer.Buffer to implement fixed-length buffers.
You write a sample with fritzer.Buffer.Write(), and can access the current
buffer contents with fritzer.Buffer.Get(). fritzer.Buffer.Head() returns
the oldest sample in the buffer.
