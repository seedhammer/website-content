--
author: SeedHammer
authorlink: https://twitter.com/SeedHammer
title: The SeedHammer Descriptor Partitioning Scheme
description: A technical explanation of the distribution model of descriptor parts used in SeedHammer backups to achieve wallet recovery from any given quorum combination.
published: 2023-08-17
image: partitioning-scheme.webp
--

## The Problem: Multisig Backup Requires All Public Keys

It is widely known that to sign a transaction with an M-of-N multisig wallet, the quorum of private keys (M) is required. What is less known is that all (N) public keys are also needed:

<div class="grid center-layout">
    <div></div>
    <div>
        <div class="img-center">
            <img src="/static/img/multisig-wallet-sign-transaction.svg" alt="In order to restore a multisig wallet all public keys must be present" style="width: 300px">
        </div> 
    </div>
    <div></div>
</div>

You don't usually notice this, because the public keys are saved in the wallet software in the form of a list of XPUBs - or more recently as a wallet output descriptor.

Creating secure and durable backups for multisig wallets necessitates the storage of both private and public keys.

*Note that singlesig backups technically also require both kinds of keys, but since the public key can be derived from the private, losing the public key for a singlesig is usually not a problem.*

## Size Matters

The straightforward backup approach for an M-of-N multisig wallet is to produce N backup shares, each containing a BIP39 seed phrase and the complete set of public keys in the form of the wallet descriptor. This way, no matter which M shares you have access to, the quorum of private keys and the complete set of public keys are available.

For SeedHammer backup plates, however, the complete descriptor is too large to engrave within the physical constraints of our metal plates. Even if we supported plates large enough to engrave the descriptor, the corresponding QR code would be exceedingly difficult to scan.

Splitting the descriptor into parts is a better approach to minimize the data size. But how? Consider the descriptor for a 2-of-3 multisig wallet. The naive split of the descriptor into 3 parts, one for each share, will not work because any quorum (2) of shares contain only two thirds of the descriptor. Splitting the descriptor into 2 parts, A and B, is better, but then how do you assign the parts to shares? If you assign A to share 1 and B to share 2, any one part assigned to share 3 will leave you with two A parts or two B parts in the unfortunate cases.

## The UR Format

The [Uniform Resources (UR) format](https://github.com/BlockchainCommons/Research/blob/master/papers/bcr-2020-005-ur.md) is an existing standard for a related problem: splitting potentially large data items into smaller shares that are easier to scan as QR codes. The standard includes an [efficient representation](https://github.com/BlockchainCommons/Research/blob/master/papers/bcr-2020-010-output-desc.md) of output descriptors, and is supported by most multisig wallet software.

Most importantly, the standard supports a fountain encoding scheme which allows combining data shares with the XOR operator, ⊕. The operator creates a mix of two or more parts the same size as the constituent parts, so A ⊕ B does not take up more engraving space than just part A or B by themselves. In addition, XOR'ed data cancel out; for example, A ⊕ B ⊕ B equals A.

Consider the 2-of-3 descriptor from before split into 2 halves, A and B. The fountain encoding assignment to shares is the following:

- Share 1: A
- Share 2: B
- Share 3: A ⊕ B

Recovery works no matter which combination of 2 shares are available:

<div class="grid">
    <img src="/static/img/part-a-b-ab.svg" alt="Illustration of how the different descriptor parts distributes over the 3 shares in a 2-of-3 multisig."> 
</div>

- Share 1 and 2 trivially combines into the full descriptor,
- Share 1 and 3 recovers part A and A ⊕ B, which recovers B by computing A ⊕ A ⊕ B = B,
- Share 2 and 3 recovers B and A ⊕ B, which recovers A by computing B ⊕ A ⊕ B = A.

The split is also minimal in size, in that each share contains data exactly half the size of the complete descriptor.

## Other wallet configurations

Up to this point, we have successfully applied the fountain encoding scheme to the following multisig configurations:

- 1-of-1
- 1-of-2
- 2-of-3
- 3-of-5
- Any M-of-N, where M = N-1

For posterity, the 3-of-5 scheme splits the descriptor into 6 parts, A B C D E F, where each share contains 2 parts, like so:

- Share 1: Contains A and F ⊕ B ⊕ E
- Share 2: Contains B and F ⊕ A ⊕ C
- Share 3: Contains C and F ⊕ B ⊕ D
- Share 4: Contains D and F ⊕ C ⊕ E
- Share 5: Contains E and F ⊕ D ⊕ A

Every combination of 3 shares recovers all 6 parts. For example, shares 1, 2, 3 contains the parts A, B, C, F ⊕ B ⊕ E, F ⊕ A ⊕ C, F ⊕ B ⊕ D, and:

- part F is computed by A ⊕ C ⊕ F ⊕ A ⊕ C,
- part E is computed by F ⊕ B ⊕ F ⊕ B ⊕ E,
- part D is computed by F ⊕ B ⊕ F ⊕ B ⊕ D.

The scheme is also minimal because every share contains 2 parts each 1/6 of the descriptor, or 1/3 of the descriptor.

<div class="text-center">
    <img src="/static/img/descriptor-sides.webp" alt="Example of the public sides of all 3 different sizes of SeedHammer plates"> 
    <p><small>Example of a descriptor side of singlesig, 2-of-3, and 3-of-5 backups respectively.</small></p>
</div>

Read more about the SeedHammer plates in general in our article [The SeedHammer Metal Plates
](/article/the-seedhammer-metal-plates).