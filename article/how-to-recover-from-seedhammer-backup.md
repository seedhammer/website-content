--
author: SeedHammer
authorlink: https://twitter.com/SeedHammer
title: How to Recover From a SeedHammer Backup
description: Learn how to recover a backup created with SeedHammer.
published: 2023-05-01
image: bitcoin-recovery.webp
--

## Recover Singlesig Wallet {#singlesig}

<p class="alert alert-warning">Keep the seed phrase and the SeedQR code away from all cameras except the signing device.</p>

1. Load the seed into a stateless signer.
    - Use the seed words or scan the CompactSeedQR.
1. Most stateless signers can export the wallet as read-only (into eg. Sparrow Wallet).

## Recover Multisig Wallet {#multisig}

<p class="alert alert-warning">Keep the SeedQR code away from any camera controlled by Sparrow Wallet.</p>

<p class="alert alert-warning">Never input the seed words into Sparrow Wallet.</p>

1. Create a new wallet in Sparrow Wallet.
1. Under "Script Policy" click the camera icon. ![Click Script Policy in Sparrow](/static/img/sparrow-import-descriptor.webp)
1. Have the quorum of plates at hand. Cover the seed side of the plates (the side that contains the seed words).
1. Scan the QR code(s) on the descriptor side of the individual plate (the side **that does not** contain the seed words).
![Scan the QR into Sparrow](/static/img/sparrow-scan-descriptor.webp) 
    - Scan every QR on all descriptor sides until Sparrow Wallet recreates the wallet.