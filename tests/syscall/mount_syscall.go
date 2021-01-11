package main

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

func do_bind_mount(s, t string) error {

	err := unix.Mount(s, t, "", unix.MS_BIND, "")
	if err != nil {
		fmt.Printf("bind-mount error received: %v\n", err)
		return err
	}

	return nil
}

func do_remount_ro(s, t string) error {

	if err := unix.Mount(s, t, "", unix.MS_BIND|unix.MS_REMOUNT|unix.MS_RDONLY, ""); err != nil {
		fmt.Printf("RO remount error received: %v\n", err)
		return err
	}

	return nil
}

func do_remount_rw(s, t string) error {

	if err := unix.Mount(s, t, "", unix.MS_BIND|unix.MS_REMOUNT, ""); err != nil {
		fmt.Printf("RW remount error received: %v\n", err)
		return err
	}

	return nil
}

func do_shiftfs_mark_ro(s, t string) error {

	if err := unix.Mount(s, s, "shiftfs", unix.MS_RDONLY, "mark"); err != nil {
		fmt.Printf("RW remount error received: %v\n", err)
		return err
	}

	return nil
}

func do_shiftfs_mark_rw(s, t string) error {

	if err := unix.Mount(s, s, "shiftfs", 0, "mark"); err != nil {
		fmt.Printf("RW remount error received: %v\n", err)
		return err
	}

	return nil
}

func do_shiftfs_mount_ro(s, t string) error {

	if err := unix.Mount(s, s, "shiftfs", unix.MS_RDONLY, ""); err != nil {
		fmt.Printf("RW remount error received: %v\n", err)
		return err
	}

	return nil
}

func do_shiftfs_mount_rw(s, t string) error {

	if err := unix.Mount(s, t, "shiftfs", 0, ""); err != nil {
		fmt.Printf("RW remount error received: %v\n", err)
		return err
	}

	return nil
}

func usage() {
	fmt.Printf("\nUsage: mount_syscall <bind | ro-remount | rw-remount> <source> <target>\n\n")
}

func main() {
	args := os.Args[1:]

	if len(args) != 3 {
		fmt.Printf("\nNumber of arguments received: %d, expected: %d\n", len(args), 3)
		usage()
		os.Exit(1)
	}

	var err error

	switch args[0] {
	case "bind":
		err = do_bind_mount(args[1], args[2])
	case "ro-remount":
		err = do_remount_ro(args[1], args[2])
	case "rw-remount":
		err = do_remount_rw(args[1], args[2])
	case "ro-shiftfs-mark":
		err = do_shiftfs_mark_ro(args[1], args[1])
	case "rw-shiftfs-mark":
		err = do_shiftfs_mark_rw(args[1], args[1])
	case "ro-shiftfs-mount":
		err = do_shiftfs_mount_ro(args[1], args[1])
	case "rw-shiftfs-mount":
		err = do_shiftfs_mount_rw(args[1], args[1])
	default:
		fmt.Printf("Unsupported command option: %s\n", args[0])
		err = fmt.Errorf("Unsupported command option: %s", args[0])
	}

	if err != nil {
		os.Exit(1)
	}
}
