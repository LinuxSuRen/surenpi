## Test the disk performance

```shell
truncate -s 10G gentoo_root.img
echo 3 | sudo tee /proc/sys/vm/drop_caches
time dd if=gentoo_root.img of=/dev/null bs=8k
```
