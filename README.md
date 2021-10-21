

## To run on Mac

1. Either run as sudo :(
2. Or download [Wireshark](https://www.wireshark.org/download.html)
3. Run the script `Install ChmodBPF` from the download
4. Run `sudo dseditgroup -o edit -a `whoami` -t user access_bpf` to add your user to the group