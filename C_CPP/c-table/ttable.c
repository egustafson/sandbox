// I am really -*-c++-*- code
//
// The goal of this example is to create an indexable 
// table that can be used in a statemachine that needs 
// ~4 pieces of data at each branch.
//


#define CCPU_MAGIC    0x55504343 // uncompressed kernel magic
#define ZIMAGE_MAGIC  0x016F2818 // compressed kernel magic
#define GZIP_MAGIC    0x00008b1f // gzip magic
#define ZCRAMFS_MAGIC GZIP_MAGIC
#define KPARAM_MAGIC  0x00ff00ff // kernel param magic num - little endian 'r', 'o', 'o', 't'

// UTIL - Bank 0
#define U0KERN_MAGIC      (unsigned long*)0x00100000 // kernel magic (uncompressed)
#define U0ZKERN_MAGIC     (unsigned long*)0x00100024 // kernel magic (compressed)
#define U0KERN_ADDR       (unsigned long*)0x00100010 // kernel binary
#define U0KERN_SIZE_OFF   (unsigned long*)0x00100004 // kernel size word offset
#define U0ZKERN_SIZE_OFF  (unsigned long*)0x0010002c // kernel size word (zImage)
#define U0PARAM_ADDR      (unsigned char*)0x00040000 // kernel parameters
#define U0ROOT_ADDR       (unsigned long*)0x00200000 // root cramfs image

// APP  - Bank 0
#define A0KERN_MAGIC      (unsigned long*)0x00400000 // kernel magic (uncompressed)
#define A0ZKERN_MAGIC     (unsigned long*)0x00400024 // kernel magic (compressed)
#define A0KERN_ADDR       (unsigned long*)0x00400010 // kernel binary
#define A0KERN_SIZE_OFF   (unsigned long*)0x00400004 // kernel size word offset
#define A0ZKERN_SIZE_OFF  (unsigned long*)0x0010002c // kernel size word (zImage)
#define A0PARAM_ADDR      (unsigned char*)0x00040000 // kernel parameters
#define A0ROOT_ADDR       (unsigned long*)0x00200000 // root cramfs image

// UTIL - Bank 1
#define U1KERN_MAGIC      (unsigned long*)0x02100000 // kernel magic (uncompressed)
#define U1ZKERN_MAGIC     (unsigned long*)0x02100024 // kernel magic (compressed)
#define U1KERN_ADDR       (unsigned long*)0x02100010 // kernel binary
#define U1KERN_SIZE_OFF   (unsigned long*)0x02100004 // kernel size word offset
#define U1ZKERN_SIZE_OFF  (unsigned long*)0x0210002c // kernel size word (zImage)
#define U1PARAM_ADDR      (unsigned char*)0x02040000 // kernel parameters
#define U1ROOT_ADDR       (unsigned long*)0x02200000 // root cramfs image

// APP  - Bank 1
#define A1KERN_MAGIC      (unsigned long*)0x02400000 // kernel magic (uncompressed)
#define A1ZKERN_MAGIC     (unsigned long*)0x02400024 // kernel magic (compressed)
#define A1KERN_ADDR       (unsigned long*)0x02400010 // kernel binary
#define A1KERN_SIZE_OFF   (unsigned long*)0x02400004 // kernel size word offset
#define A1ZKERN_SIZE_OFF  (unsigned long*)0x0210002c // kernel size word (zImage)
#define A1PARAM_ADDR      (unsigned char*)0x02040000 // kernel parameters
#define A1ROOT_ADDR       (unsigned long*)0x02200000 // root cramfs image


struct boot_rec {
    unsigned long  magic;       // The magic number
    unsigned long* kmagic;      // Magic num location
    unsigned long* kaddr;       // Kernel binary location
    unsigned long* ksize;       // Kernel size
    unsigned char* kparam;      // Kernel parameters location
    unsigned long* rootfs;      // Root file system location
};


static struct boot_rec brec[] = {
    {   // 0
        magic:  CCPU_MAGIC,
        kmagic: U0KERN_MAGIC,
        kaddr:  U0KERN_ADDR,
        ksize:  U0KERN_SIZE_OFF,
        kparam: U0PARAM_ADDR,
        rootfs: U0ROOT_ADDR,
    },
    {   // 1
        magic:  ZIMAGE_MAGIC,
        kmagic: U0ZKERN_MAGIC,
        kaddr:  U0KERN_ADDR,
        ksize:  U0ZKERN_SIZE_OFF,
        kparam: U0PARAM_ADDR,
        rootfs: U0ROOT_ADDR,
    },
    {   // 2
        magic:  CCPU_MAGIC,
        kmagic: A0KERN_MAGIC,
        kaddr:  A0KERN_ADDR,
        ksize:  A0KERN_SIZE_OFF,
        kparam: A0PARAM_ADDR,
        rootfs: A0ROOT_ADDR,
    },
    {   // 3
        magic:  ZIMAGE_MAGIC,
        kmagic: A0ZKERN_MAGIC,
        kaddr:  A0KERN_ADDR,
        ksize:  A0ZKERN_SIZE_OFF,
        kparam: A0PARAM_ADDR,
        rootfs: A0ROOT_ADDR,
    },
    {   // 4
        magic:  CCPU_MAGIC,
        kmagic: U1KERN_MAGIC,
        kaddr:  U1KERN_ADDR,
        ksize:  U1KERN_SIZE_OFF,
        kparam: U1PARAM_ADDR,
        rootfs: U1ROOT_ADDR,
    },
    {   // 5
        magic:  ZIMAGE_MAGIC,
        kmagic: U1ZKERN_MAGIC,
        kaddr:  U1KERN_ADDR,
        ksize:  U1ZKERN_SIZE_OFF,
        kparam: U1PARAM_ADDR,
        rootfs: U1ROOT_ADDR,
    },
    {   // 6
        magic:  CCPU_MAGIC,
        kmagic: A1KERN_MAGIC,
        kaddr:  A1KERN_ADDR,
        ksize:  A1KERN_SIZE_OFF,
        kparam: A1PARAM_ADDR,
        rootfs: A1ROOT_ADDR,
    },
    {   // 7
        magic:  ZIMAGE_MAGIC,
        kmagic: A1ZKERN_MAGIC,
        kaddr:  A1KERN_ADDR,
        ksize:  A1ZKERN_SIZE_OFF,
        kparam: A1PARAM_ADDR,
        rootfs: A1ROOT_ADDR,
    },
};

static int boot_order[4][8] = 
    {
        { 0, 1, 2, 3, 4, 5, 6, 7 }, // U/0 First
        { 2, 3, 0, 1, 6, 7, 4, 5 }, // A/0 First
        { 4, 5, 6, 7, 0, 1, 2, 3 }, // U/1 First
        { 6, 7, 4, 5, 2, 3, 0, 1 }, // A/1 First
    };

// static int u0_first[] = { 0, 1, 2, 3, 4, 5, 6, 7 };
// static int a0_first[] = { 2, 3, 0, 1, 6, 7, 4, 5 };
// static int u1_first[] = { 4, 5, 6, 7, 0, 1, 2, 3 };
// static int a1_first[] = { 6, 7, 4, 5, 2, 3, 0, 1 };



// Algorithm:
//   
//   loop through each canidate record and
//     Fetch potential kernel's magic num
//     if [ magic = desired magic ] 
//       load:
//         kimage_addr
//         kimage_size
//         kparam_addr
//         rimage_addr
//         -break loop-
//     end if
//   end loop
// 
//   If no good kernel magic - bail to forth
// 
//   If filesystem magic bad - bail to forth
// 
//   If kern param magic bad - use hard coded params
// 
//   Load kernel parameters into param memory space, pad w/ 0's
// 
//   Load kernel image into image memory space
// 
//   Branch to kernel.

  

  


