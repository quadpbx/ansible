# QuadPBX Ansible (Trixie)

This is just a random collection of tools and roles to help configure
QuadPBX machines.

# Prep

Run `make setup` to install all the basic prereqs and set up a
generic .gitconfig - It'll ask you for your name and email. Don't
accept the defaults unless you ACTUALLY ARE xrobau

## Volumes

You probably want to add these as seperate volumes:

* /var/lib/docker
* /usr/local/data
* /usr/local/build

## SSH Keys

Update `group_vars/all/ssh_keys` unless you're actually xrobau. At least add yours!

# Base setup

Run `make base` to set the local machine up as a development
machine.

# Example

After installing the base vm, I added a 500gb disk as `/dev/vdb`

## Partitioning

```
root@template:~# fdisk /dev/vdb

Welcome to fdisk (util-linux 2.41).
Changes will remain in memory only, until you decide to write them.
Be careful before using the write command.

Device does not contain a recognized partition table.
Created a new DOS (MBR) disklabel with disk identifier 0x3974446b.

Command (m for help): n
Partition type
   p   primary (0 primary, 0 extended, 4 free)
   e   extended (container for logical partitions)
Select (default p):

Using default response p.
Partition number (1-4, default 1):
First sector (2048-1048575999, default 2048):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (2048-1048575999, default 1048575999):

Created a new partition 1 of type 'Linux' and of size 500 GiB.

Command (m for help): t
Selected partition 1
Hex code or alias (type L to list all): L

00 Empty            27 Hidden NTFS Win  82 Linux swap / So  c1 DRDOS/sec (FAT-
01 FAT12            39 Plan 9           83 Linux            c4 DRDOS/sec (FAT-
02 XENIX root       3c PartitionMagic   84 OS/2 hidden or   c6 DRDOS/sec (FAT-
03 XENIX usr        40 Venix 80286      85 Linux extended   c7 Syrinx
04 FAT16 <32M       41 PPC PReP Boot    86 NTFS volume set  da Non-FS data
05 Extended         42 SFS              87 NTFS volume set  db CP/M / CTOS / .
06 FAT16            4d QNX4.x           88 Linux plaintext  de Dell Utility
07 HPFS/NTFS/exFAT  4e QNX4.x 2nd part  8e Linux LVM        df BootIt
08 AIX              4f QNX4.x 3rd part  93 Amoeba           e1 DOS access
09 AIX bootable     50 OnTrack DM       94 Amoeba BBT       e3 DOS R/O
0a OS/2 Boot Manag  51 OnTrack DM6 Aux  9f BSD/OS           e4 SpeedStor
0b W95 FAT32        52 CP/M             a0 IBM Thinkpad hi  ea Linux extended
0c W95 FAT32 (LBA)  53 OnTrack DM6 Aux  a5 FreeBSD          eb BeOS fs
0e W95 FAT16 (LBA)  54 OnTrackDM6       a6 OpenBSD          ee GPT
0f W95 Ext'd (LBA)  55 EZ-Drive         a7 NeXTSTEP         ef EFI (FAT-12/16/
10 OPUS             56 Golden Bow       a8 Darwin UFS       f0 Linux/PA-RISC b
11 Hidden FAT12     5c Priam Edisk      a9 NetBSD           f1 SpeedStor
12 Compaq diagnost  61 SpeedStor        ab Darwin boot      f4 SpeedStor
14 Hidden FAT16 <3  63 GNU HURD or Sys  af HFS / HFS+       f2 DOS secondary
16 Hidden FAT16     64 Novell Netware   b7 BSDI fs          f8 EBBR protective
17 Hidden HPFS/NTF  65 Novell Netware   b8 BSDI swap        fb VMware VMFS
18 AST SmartSleep   70 DiskSecure Mult  bb Boot Wizard hid  fc VMware VMKCORE
1b Hidden W95 FAT3  75 PC/IX            bc Acronis FAT32 L  fd Linux raid auto
1c Hidden W95 FAT3  80 Old Minix        be Solaris boot     fe LANstep
1e Hidden W95 FAT1  81 Minix / old Lin  bf Solaris          ff BBT
24 NEC DOS

Aliases:
   linux          - 83
   swap           - 82
   extended       - 05
   uefi           - EF
   raid           - FD
   lvm            - 8E
   linuxex        - 85
Hex code or alias (type L to list all): 8e
Changed type of partition 'Linux' to 'Linux LVM'.

Command (m for help): w
The partition table has been altered.
Calling ioctl() to re-read partition table.
Syncing disks.

```

## LVM pvcreate


```
root@template:~# pvcreate /dev/vdb1
```

## Create volumes

I just created them a 100g each, so they can be expanded later.


```
root@template:~/ansible# vgcreate build /dev/vdb1
  Volume group "build" successfully created
root@template:~/ansible# lvcreate -n data -l100G build
  Invalid argument for --extents: 100G
  Error during parsing of command line.
root@template:~/ansible# lvcreate -n data -L100G build
  Logical volume "data" created.
root@template:~/ansible# lvcreate -n build -L100G build
  Logical volume "build" created.
root@template:~/ansible#
```

## Format

Note: I put docker on a different pv, which is why it's not here.


```
root@quadbuilder:~/ansible# mkfs.ext4 /dev/build/data
mke2fs 1.47.2 (1-Jan-2025)
Discarding device blocks: done
Creating filesystem with 26214400 4k blocks and 6553600 inodes
Filesystem UUID: a261065c-d0e6-48eb-add4-9ec0e7a6bb9e
Superblock backups stored on blocks:
        32768, 98304, 163840, 229376, 294912, 819200, 884736, 1605632, 2654208,
        4096000, 7962624, 11239424, 20480000, 23887872

Allocating group tables: done
Writing inode tables: done
Creating journal (131072 blocks): done
Writing superblocks and filesystem accounting information: done

root@quadbuilder:~/ansible# mkfs.ext4 /dev/build/build
mke2fs 1.47.2 (1-Jan-2025)
Discarding device blocks: done
Creating filesystem with 26214400 4k blocks and 6553600 inodes
Filesystem UUID: ac52da4c-e91a-4656-9138-8fc472e36192
Superblock backups stored on blocks:
        32768, 98304, 163840, 229376, 294912, 819200, 884736, 1605632, 2654208,
        4096000, 7962624, 11239424, 20480000, 23887872

Allocating group tables: done
Writing inode tables: done
Creating journal (131072 blocks): done
Writing superblocks and filesystem accounting information: done

root@quadbuilder:~/ansible#
```

## Mount

Mounting is up to you - mkdir and edit /etc/fstab - if you can't do this,
ask for help somewhere.


