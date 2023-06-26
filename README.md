# GHIP-8

Go implementation of [CHIP-8](https://en.wikipedia.org/wiki/CHIP-8). Currently, the emulator[^1] can only render the contents of `ibm.ch8` with no interactions.

Currently outputs the following in your terminal:

![](./ibm.png)

## How to run

If you have Go installed you may run this application with the following command at the root of this project:

```bash
go run .
```

Alternatively, you may utilize this project's devcontainer (which has Go preconfigured). [Learn more about devcontainers here](https://code.visualstudio.com/docs/devcontainers/containers).


## References

- [Cowgod's Chip-8 Technical Reference](http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#2.2)
- [How to Write a Go Emulator by Tobias V Langhoff](https://tobiasvl.github.io/blog/write-a-chip-8-emulator/)

## Credit

- Particle demo by zeroZShadow
- Test rom by corax89


[^1]: Technically speaking, CHIP-8 implementations aren't emulators. Langhoff writes: *"An emulator emulates physical hardware in software, but CHIP-8 isn’t a piece of hardware. To be pedantic, you’re writing a CHIP-8 interpreter."*
