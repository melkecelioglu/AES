# AES

This is a Simple Example to Encrypt a Text message and Decrypt a file using a Key which is 16 byte (128-bit) password.
Our example creates two files one which contains encrypted data and other with AES key.

AES stands for Advanced Encryption Standard, also known by its original name Rijndael, is a specification for the encryption of electronic data established by the U.S. National Institute of Standards and Technology in 2001.



cyptographic Hash \
Simple Example of Cryptographic Hash Functions used in Golang.
Using Secure Hash Algorithm and Message Digest Algorithm
SHA256,512 and MD5 Packages used.

Hashing: The process of taking plaintext and transforms it into a digest of the plaintext information, in such a way that it is not intended to be decrypt is called Hashing. The output of Hashing is known as a hash, hash value, or message digest.

Cryptographic hash functions are those hash functions which have a base on block ciphers. Examples of cryptographic hash functions include message digest (MD) algorithm based such as MD2, MD4 and MD5 and secure hash algorithm based (SHA) such as SHA-1, SHA-224, SHA-256, and SHA-512.

Message-Digest algorithm 5 (MD5):
This algorithm takes a variable-length message as input and generates a fixed-length message digest of 128 bits. This algorithm uses Big-endian scheme where the least significant byte of a 32-bit word will be stored in the low-address byte position.

Secure Hash Algorithm: SHA-2 family consists of six hash functions with digests (hash values) that are 224, 256, 384 or 512 bits: SHA-224, SHA-256, SHA-384, SHA-512, SHA-512/224, SHA-512/256. SHA-256 and SHA-512 are novel hash functions computed with eight 32-bit and 64-bit words, respectively. 

Due to vulnerabilities in MD5 and SHA-1, some agencies has started using SHA-2 as early as 2011. SHA-2 is more secure, and it has a 256-bit and 512-bit block sizes. A common hash used on the Internet is SHA-2, SHA-1 is now deprecated and should be replaced if it's in use.

List of real world examples where we are using Hash functions:
Cloud storage systems use hash functions to analyze identical files and to find changed files.
Git revision control system uses hash functions to detect files in a repository.
Bitcoin using a hash function in its proof-of-work systems.
Forensic analysts use hash values to verify that digital artifacts hasn't altered.
NIDS use hashes to analyze known-malicious data going through a network.