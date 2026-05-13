--
author: SeedHammer
authorlink: https://twitter.com/SeedHammer
title: The SeedHammer Metal Plates
description: Guidelines, limitations and information about the SeedHammer steel plates.
published: 2023-02-27
image: seedhammer-metal-plates.webp
--

It is possible to backup both single and multisig setups. However there are some guidelines for which type of plates to use for each setup.

## Overview

To get a full overview look up this table:

|                                        | SH02 | SH03 |
|----------------------------------------|------|------|
| Singlesig, 12 and 24 words             | Yes  |      |
| 2-of-2, 12 and 24 words                | Yes  |      |
| 2-of-3, 12 and 24 words                | Yes  |      |
| 3-of-3, 12 and 24 words                | Yes  |      |
| 3-of-5, 12 and 24 words                |      | Yes  |
| M-of-N, where M = N-1, 12 and 24 words | Yes  |      |
| 1-of-2, 12 and 24 words                |      | Yes  |

## Plate SH01

The SH01 plate was deprecated as part of release [v1.3.1](https://github.com/seedhammer/seedhammer/releases/tag/v1.3.1). Every engraving that would have used SH01 will now use the SH02 plate size instead.

<del>SH01 is credit card sized and works for single-signature wallets with 12 word seed phrases.</del>

<del>The plate are split into two - a seed side and a descriptor side.</del>

- <del>Seed Side:</del>
    - <del>12 words fit on one side incl. a CompactSeedQR, the master fingerprint and plate number.</del>
- <del>Descriptor Side:</del>
    - <del>On the opposite side the full descriptor is engraved together with a QR code representing the same data.</del>

## Plate SH02

SH02 is square and works for single-signature wallets with 12 and 24 words, 2-of-2, 2-of-3, 3-of-3 multi-signature wallets, as well as M-of-N, where M = N-1, 12 and 24 words.

The plate are split into two - a seed side and a descriptor side.

It is possible to engrave both on the same plate as well as to split them into two plates. Two plates make it possible to screw the plates together, with the shiny sides facing out.

- Seed Side:
    - 12 or 24 words fit on one side incl. a CompactSeedQR, the master fingerprint and plate number.
- Descriptor Side:
    - On the opposite side different parts of the descriptor are engraved together with one or two QR codes representing the same data. In the case of a singlesig wallet the full descriptor is engraved.

<div class="grid">
    <div>
        <div class="img-center"><img src="/static/img/multisig-1-of-3-seed-side.webp" alt="The SeedHammer Metal Plate SH02 with 24 words" style="width: 300px"><small>SH02: Seed side in a 2-of-3 multisig.</small></div>
    </div>
    <div>
        <div class="img-center"><img src="/static/img/multisig-1-of-3-descriptor-side.webp" alt="The SeedHammer Metal Plate SH02 with 24 words" style="width: 300px"><small>SH02: Descriptor side in a 2-of-3 multisig.</small></div>       
    </div>
</div>

## Plate SH03

SH03 is rectangular and works for 1-of-2 and 3-of-5 multisig wallets.

The plate are split into two - a seed side and a descriptor side.

It is possible to engrave both on the same plate as well as to split them into two plates. Two plates make it possible to screw the plates together, with the shiny sides facing out.

- Seed Side:
    - 12 or 24 words fit on one side incl. a CompactSeedQR, the master fingerprint and plate number.
- Descriptor Side:
    - On the opposite side different parts of the descriptor are engraved together with one or two QR codes representing the same data.

<div class="grid">
    <div>
        <div class="img-center"><img src="/static/img/multisig-1-of-5-seed-side.webp" alt="The SeedHammer Metal Plate SH03 with 24 words" style="width: 300px"><small>SH03: Seed side in a 3-of-5 multisig.</small></div>
    </div>
    <div>
        <div class="img-center"><img src="/static/img/multisig-1-of-5-descriptor-side.webp" alt="The SeedHammer Metal Plate SH03 with 24 words" style="width: 300px"><small>SH03: Descriptor side in a 3-of-5 multisig.</small></div>
    </div>
</div>

## Advantages

The descriptor-parts are distributed among the plates in such a way that any quorum can recover the wallet. No need to store a digital or paper copy of the descriptor. Read our article [The SeedHammer Descriptor Partitioning Scheme](/article/the-seedhammer-descriptor-partitioning-scheme) for a more technical explaination.

The QR codes make it easy and fast to recover the multisig wallet by scanning them into a coordinator software like Sparrow Wallet. 

Complies with multisig standards, no special tools needed for recovery.

<div class="grid center-layout">
    <div></div>
    <div>
        <div class="img-center"><img src="/static/img/3-of-5-multisig.webp" alt="The SeedHammer Metal Plate SH03 with 24 words" style="width: 300px"><small>Any combination of the minimum number of shares will be enough to perform a full recovery.</small></div>      
    </div>
    <div></div>
</div>

## Backup Plate Material

The metal plates are made of 2 or 3 mm thick stainless steel 316L - also known as marine grade steel.

Steel 316's superior resistance to corrosion and heat as well as its ability to withstand harsh environments makes it a more suitable choice for cold storage wallets compared to ordinary stainless steel 304.

Read more about the advantages of steel 316 [here](/article/steel-316).
